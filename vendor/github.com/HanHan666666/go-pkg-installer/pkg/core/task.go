// Package core provides the task interface and runner for installation tasks.
package core

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// TaskState represents the state of a task.
type TaskState int

const (
	TaskPending TaskState = iota
	TaskRunning
	TaskCompleted
	TaskFailed
	TaskCancelled
	TaskRolledBack
)

func (s TaskState) String() string {
	switch s {
	case TaskPending:
		return "pending"
	case TaskRunning:
		return "running"
	case TaskCompleted:
		return "completed"
	case TaskFailed:
		return "failed"
	case TaskCancelled:
		return "cancelled"
	case TaskRolledBack:
		return "rolled_back"
	default:
		return "unknown"
	}
}

// FailurePolicy defines how to handle task failures.
type FailurePolicy int

const (
	// FailureAbort aborts the entire task run on failure.
	FailureAbort FailurePolicy = iota
	// FailureSkip skips the failed task and continues.
	FailureSkip
	// FailureRollback triggers rollback of completed tasks.
	FailureRollback
	// FailureRetry retries the failed task.
	FailureRetry
)

// TaskResult contains the result of a task execution.
type TaskResult struct {
	TaskID    string
	TaskType  string
	State     TaskState
	Error     error
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
}

// BaseTask provides common functionality for task implementations.
type BaseTask struct {
	TaskID   string
	TaskType string
	Config   map[string]any
}

// ID returns the task identifier.
func (t *BaseTask) ID() string {
	return t.TaskID
}

// Type returns the task type name.
func (t *BaseTask) Type() string {
	return t.TaskType
}

// CanRollback returns false by default.
func (t *BaseTask) CanRollback() bool {
	return false
}

// Rollback does nothing by default.
func (t *BaseTask) Rollback(ctx *InstallContext, bus *EventBus) error {
	return nil
}

// GetConfigString safely retrieves a string from config.
func (t *BaseTask) GetConfigString(key string) string {
	if v, ok := t.Config[key].(string); ok {
		return v
	}
	return ""
}

// GetConfigBool safely retrieves a bool from config.
func (t *BaseTask) GetConfigBool(key string) bool {
	if v, ok := t.Config[key].(bool); ok {
		return v
	}
	return false
}

// GetConfigInt safely retrieves an int from config.
func (t *BaseTask) GetConfigInt(key string) int {
	switch v := t.Config[key].(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	default:
		return 0
	}
}

// TaskRunner executes a sequence of tasks with progress tracking and rollback support.
type TaskRunner struct {
	mu sync.RWMutex

	tasks          []Task
	results        []TaskResult
	completedTasks []Task // For rollback

	ctx           *InstallContext
	bus           *EventBus
	failurePolicy FailurePolicy
	maxRetries    int

	// Cancellation
	cancelCtx  context.Context
	cancelFunc context.CancelFunc
	cancelled  bool

	// State
	currentIndex int
	running      bool
}

// NewTaskRunner creates a new TaskRunner.
func NewTaskRunner(ctx *InstallContext, bus *EventBus) *TaskRunner {
	cancelCtx, cancelFunc := context.WithCancel(context.Background())
	return &TaskRunner{
		tasks:          make([]Task, 0),
		results:        make([]TaskResult, 0),
		completedTasks: make([]Task, 0),
		ctx:            ctx,
		bus:            bus,
		failurePolicy:  FailureAbort,
		maxRetries:     3,
		cancelCtx:      cancelCtx,
		cancelFunc:     cancelFunc,
		currentIndex:   -1,
	}
}

// SetFailurePolicy sets the failure handling policy.
func (r *TaskRunner) SetFailurePolicy(policy FailurePolicy) {
	r.failurePolicy = policy
}

// SetMaxRetries sets the maximum retry count for FailureRetry policy.
func (r *TaskRunner) SetMaxRetries(count int) {
	r.maxRetries = count
}

// AddTask adds a task to the run queue.
func (r *TaskRunner) AddTask(task Task) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks = append(r.tasks, task)
}

// AddTasks adds multiple tasks to the run queue.
func (r *TaskRunner) AddTasks(tasks []Task) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks = append(r.tasks, tasks...)
}

// QueueConfig adds a task from a TaskConfig to the run queue.
// This method uses the task registry to create the task.
func (r *TaskRunner) QueueConfig(config TaskConfig) error {
	taskType := config.Type
	factory, ok := Tasks.Get(taskType)
	if !ok && IsGoExtension(taskType) {
		factory, ok = Tasks.Get(StripGoPrefix(taskType))
	}
	if !ok {
		return fmt.Errorf("unknown task type: %s", taskType)
	}

	// Build params map including inline params
	params := make(map[string]any)
	for k, v := range config.Params {
		params[k] = v
	}
	params["type"] = taskType
	if config.ID != "" {
		params["id"] = config.ID
	}

	task, err := factory(params, r.ctx)
	if err != nil {
		return fmt.Errorf("failed to create task %s: %w", config.Type, err)
	}

	r.AddTask(task)
	return nil
}

// Run executes all tasks in sequence.
func (r *TaskRunner) Run() error {
	r.mu.Lock()
	if r.running {
		r.mu.Unlock()
		return errors.New("runner already running")
	}
	r.running = true
	r.mu.Unlock()

	defer func() {
		r.mu.Lock()
		r.running = false
		r.mu.Unlock()
	}()

	totalTasks := len(r.tasks)
	if totalTasks == 0 {
		return nil
	}

	for i, task := range r.tasks {
		r.mu.Lock()
		r.currentIndex = i
		r.mu.Unlock()

		// Check cancellation
		select {
		case <-r.cancelCtx.Done():
			return r.handleCancellation()
		default:
		}

		result := r.runTask(task, i, totalTasks)
		r.mu.Lock()
		r.results = append(r.results, result)
		r.mu.Unlock()

		if result.State == TaskFailed {
			r.ctx.AddError(result.Error)
			if result.Error != nil {
				r.ctx.AddLog(LogError, fmt.Sprintf("Task %s failed: %v", task.ID(), result.Error))
			}
			switch r.failurePolicy {
			case FailureAbort:
				return result.Error
			case FailureSkip:
				r.ctx.AddLog(LogWarn, fmt.Sprintf("Skipping failed task: %s", task.ID()))
				continue
			case FailureRollback:
				if err := r.rollback(); err != nil {
					return fmt.Errorf("rollback failed: %w (original error: %v)", err, result.Error)
				}
				return result.Error
			case FailureRetry:
				// Retry is handled in runTask
				return result.Error
			}
		}

		if result.State == TaskCompleted && task.CanRollback() {
			r.mu.Lock()
			r.completedTasks = append(r.completedTasks, task)
			r.mu.Unlock()
		}
	}

	return nil
}

func (r *TaskRunner) runTask(task Task, index, total int) TaskResult {
	result := TaskResult{
		TaskID:    task.ID(),
		TaskType:  task.Type(),
		StartTime: time.Now(),
	}

	// Validate task
	if err := task.Validate(); err != nil {
		result.State = TaskFailed
		result.Error = fmt.Errorf("validation failed: %w", err)
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		return result
	}

	// Publish start event
	r.bus.PublishTaskStart(task.ID(), task.Type())

	// Calculate progress
	baseProgress := float64(index) / float64(total)
	r.bus.PublishProgress(task.ID(), baseProgress, fmt.Sprintf("Running: %s", task.ID()))

	// Execute with retry support
	var lastErr error
	retries := 0
	maxRetries := 1
	if r.failurePolicy == FailureRetry {
		maxRetries = r.maxRetries
	}

	for retries < maxRetries {
		select {
		case <-r.cancelCtx.Done():
			result.State = TaskCancelled
			result.Error = errors.New("task cancelled")
			result.EndTime = time.Now()
			result.Duration = result.EndTime.Sub(result.StartTime)
			return result
		default:
		}

		err := task.Execute(r.ctx, r.bus)
		if err == nil {
			result.State = TaskCompleted
			break
		}

		lastErr = err
		retries++

		if retries < maxRetries {
			r.ctx.AddLog(LogWarn, fmt.Sprintf("Task %s failed, retrying (%d/%d): %v", task.ID(), retries, maxRetries, err))
			time.Sleep(time.Second * time.Duration(retries)) // Exponential backoff
		}
	}

	if result.State != TaskCompleted {
		result.State = TaskFailed
		result.Error = lastErr
		r.bus.PublishTaskError(task.ID(), task.Type(), lastErr)
	} else {
		r.bus.PublishTaskComplete(task.ID(), task.Type())
	}

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	// Update overall progress
	progress := float64(index+1) / float64(total)
	r.ctx.SetProgress(progress)
	r.bus.PublishProgress(task.ID(), progress, fmt.Sprintf("Completed: %s", task.ID()))

	return result
}

// Cancel cancels the running tasks.
func (r *TaskRunner) Cancel() {
	r.mu.Lock()
	r.cancelled = true
	r.mu.Unlock()
	r.cancelFunc()
}

// IsCancelled returns true if the runner was cancelled.
func (r *TaskRunner) IsCancelled() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.cancelled
}

func (r *TaskRunner) handleCancellation() error {
	r.ctx.AddLog(LogWarn, "Task execution cancelled")
	if r.failurePolicy == FailureRollback {
		return r.rollback()
	}
	return errors.New("execution cancelled")
}

// rollback executes rollback for all completed tasks in reverse order.
func (r *TaskRunner) rollback() error {
	r.mu.RLock()
	tasks := make([]Task, len(r.completedTasks))
	copy(tasks, r.completedTasks)
	r.mu.RUnlock()

	r.ctx.AddLog(LogInfo, fmt.Sprintf("Rolling back %d tasks", len(tasks)))

	// Rollback in reverse order
	for i := len(tasks) - 1; i >= 0; i-- {
		task := tasks[i]
		if !task.CanRollback() {
			continue
		}

		r.ctx.AddLog(LogInfo, fmt.Sprintf("Rolling back: %s", task.ID()))
		if err := task.Rollback(r.ctx, r.bus); err != nil {
			r.ctx.AddLog(LogError, fmt.Sprintf("Rollback failed for %s: %v", task.ID(), err))
			return err
		}

		r.mu.Lock()
		r.results = append(r.results, TaskResult{
			TaskID:   task.ID(),
			TaskType: task.Type(),
			State:    TaskRolledBack,
		})
		r.mu.Unlock()
	}

	return nil
}

// Rollback manually triggers rollback of completed tasks.
func (r *TaskRunner) Rollback() error {
	return r.rollback()
}

// Results returns all task results.
func (r *TaskRunner) Results() []TaskResult {
	r.mu.RLock()
	defer r.mu.RUnlock()

	results := make([]TaskResult, len(r.results))
	copy(results, r.results)
	return results
}

// Progress returns the current progress (0.0 to 1.0).
func (r *TaskRunner) Progress() float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.tasks) == 0 {
		return 1.0
	}
	return float64(r.currentIndex+1) / float64(len(r.tasks))
}

// IsRunning returns true if the runner is currently executing tasks.
func (r *TaskRunner) IsRunning() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.running
}
