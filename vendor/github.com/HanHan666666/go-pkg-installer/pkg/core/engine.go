// Package core provides the workflow engine and step management.
package core

import (
	"errors"
	"fmt"
	"sync"
)

// StepStatus represents the state of a step.
type StepStatus int

const (
	StepNotStarted StepStatus = iota
	StepCurrent
	StepCompleted
	StepBlocked
	StepDisabled
)

// Step represents a single step in the workflow.
type Step struct {
	ID        string
	Title     string
	ScreenCfg map[string]any
	TasksCfg  []map[string]any
	GuardsCfg []map[string]any
	Next      string // Next step ID (empty = sequential)
	Prev      string // Previous step ID (empty = sequential)
	Branch    *BranchConfig
	AllowJump bool
	Route     string
	// Config holds the original step configuration
	Config *StepConfig
}

// Flow represents a complete workflow (install, uninstall, etc.).
type Flow struct {
	ID    string
	Entry string // Entry step ID
	Steps []*Step
}

// Workflow manages the state machine for step navigation.
type Workflow struct {
	mu sync.RWMutex

	// Configuration
	flows   map[string]*Flow
	current *Flow

	// State
	stepIndex  map[string]int // step ID -> index in steps
	stepStatus map[string]StepStatus
	visited    map[string]bool
	currentIdx int

	// Dependencies
	ctx *InstallContext
	bus *EventBus
}

// NewWorkflow creates a new workflow engine.
func NewWorkflow(ctx *InstallContext, bus *EventBus) *Workflow {
	return &Workflow{
		flows:      make(map[string]*Flow),
		stepIndex:  make(map[string]int),
		stepStatus: make(map[string]StepStatus),
		visited:    make(map[string]bool),
		currentIdx: -1,
		ctx:        ctx,
		bus:        bus,
	}
}

// AddFlow adds a flow to the workflow.
func (w *Workflow) AddFlow(flow *Flow) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if flow.ID == "" {
		return errors.New("flow ID cannot be empty")
	}
	if len(flow.Steps) == 0 {
		return errors.New("flow must have at least one step")
	}
	if _, exists := w.flows[flow.ID]; exists {
		return fmt.Errorf("flow %q already exists", flow.ID)
	}

	w.flows[flow.ID] = flow
	return nil
}

// SelectFlow activates a flow by ID and initializes step state.
func (w *Workflow) SelectFlow(flowID string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	flow, ok := w.flows[flowID]
	if !ok {
		return fmt.Errorf("flow %q not found", flowID)
	}

	w.current = flow
	w.stepIndex = make(map[string]int)
	w.stepStatus = make(map[string]StepStatus)
	w.visited = make(map[string]bool)

	// Build step index
	for i, step := range flow.Steps {
		w.stepIndex[step.ID] = i
		w.stepStatus[step.ID] = StepNotStarted
	}

	// Find entry step
	entryIdx := 0
	if flow.Entry != "" {
		if idx, ok := w.stepIndex[flow.Entry]; ok {
			entryIdx = idx
		}
	}

	w.currentIdx = entryIdx
	w.stepStatus[flow.Steps[entryIdx].ID] = StepCurrent
	w.visited[flow.Steps[entryIdx].ID] = true

	// Update runtime context
	w.ctx.Runtime.FlowID = flowID
	w.ctx.Runtime.CurrentStep = flow.Steps[entryIdx].ID

	return nil
}

// CurrentStep returns the current step.
func (w *Workflow) CurrentStep() *Step {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if w.current == nil || w.currentIdx < 0 || w.currentIdx >= len(w.current.Steps) {
		return nil
	}
	return w.current.Steps[w.currentIdx]
}

// CurrentStepID returns the current step ID.
func (w *Workflow) CurrentStepID() string {
	step := w.CurrentStep()
	if step == nil {
		return ""
	}
	return step.ID
}

// Steps returns all steps in the current flow.
func (w *Workflow) Steps() []*Step {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if w.current == nil {
		return nil
	}
	return w.current.Steps
}

// StepStatus returns the status of a step.
func (w *Workflow) StepStatus(stepID string) StepStatus {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.stepStatus[stepID]
}

// IsVisited returns true if the step has been visited.
func (w *Workflow) IsVisited(stepID string) bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.visited[stepID]
}

// CanGoNext checks if we can proceed to the next step.
// Returns nil if can proceed, error explaining why otherwise.
func (w *Workflow) CanGoNext() error {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.canGoNextUnlocked()
}

func (w *Workflow) canGoNextUnlocked() error {
	if w.current == nil {
		return errors.New("no flow selected")
	}

	step := w.current.Steps[w.currentIdx]

	// Check guards
	for _, guardCfg := range step.GuardsCfg {
		typeName, ok := guardCfg["type"].(string)
		if !ok {
			continue
		}

		factory, ok := Guards.Get(typeName)
		if !ok {
			// Try go: prefix for extension
			if IsGoExtension(typeName) {
				factory, ok = Guards.Get(StripGoPrefix(typeName))
			}
		}
		if !ok {
			return fmt.Errorf("unknown guard type: %s", typeName)
		}

		guard, err := factory(guardCfg)
		if err != nil {
			return fmt.Errorf("failed to create guard: %w", err)
		}

		if err := guard.Check(w.ctx); err != nil {
			return err
		}
	}

	return nil
}

// Next moves to the next step.
// Returns the new step ID, or error if cannot proceed.
func (w *Workflow) Next() (string, error) {
	w.mu.Lock()

	if err := w.canGoNextUnlocked(); err != nil {
		w.mu.Unlock()
		return "", err
	}

	step := w.current.Steps[w.currentIdx]
	oldStepID := step.ID

	// Determine next step
	nextID := ""

	// Check for branch
	if step.Branch != nil {
		nextID = w.evaluateBranch(step.Branch)
	}

	// Check for explicit next
	if nextID == "" && step.Next != "" {
		nextID = step.Next
	}

	// Default to sequential
	if nextID == "" {
		if w.currentIdx+1 < len(w.current.Steps) {
			nextID = w.current.Steps[w.currentIdx+1].ID
		} else {
			w.mu.Unlock()
			return "", errors.New("no more steps")
		}
	}

	// Find and activate next step
	nextIdx, ok := w.stepIndex[nextID]
	if !ok {
		w.mu.Unlock()
		return "", fmt.Errorf("step %q not found", nextID)
	}

	// Skip disabled steps
	for w.stepStatus[w.current.Steps[nextIdx].ID] == StepDisabled {
		nextIdx++
		if nextIdx >= len(w.current.Steps) {
			w.mu.Unlock()
			return "", errors.New("no more steps")
		}
		nextID = w.current.Steps[nextIdx].ID
	}

	// Update status
	w.stepStatus[oldStepID] = StepCompleted
	w.stepStatus[nextID] = StepCurrent
	w.currentIdx = nextIdx
	w.visited[nextID] = true

	// Update context
	w.ctx.Runtime.CurrentStep = nextID

	bus := w.bus
	w.mu.Unlock()

	// Emit event after unlocking to avoid deadlocks in handlers.
	if bus != nil {
		bus.PublishStepChange(oldStepID, nextID)
	}

	return nextID, nil
}

// Prev moves to the previous step.
// Returns the new step ID, or error if cannot go back.
func (w *Workflow) Prev() (string, error) {
	w.mu.Lock()

	if w.current == nil {
		w.mu.Unlock()
		return "", errors.New("no flow selected")
	}

	step := w.current.Steps[w.currentIdx]
	oldStepID := step.ID

	// Determine previous step
	prevID := ""

	// Check for explicit prev
	if step.Prev != "" {
		prevID = step.Prev
	}

	// Default to sequential
	if prevID == "" {
		if w.currentIdx-1 >= 0 {
			prevID = w.current.Steps[w.currentIdx-1].ID
		} else {
			w.mu.Unlock()
			return "", errors.New("already at first step")
		}
	}

	// Find and activate previous step
	prevIdx, ok := w.stepIndex[prevID]
	if !ok {
		w.mu.Unlock()
		return "", fmt.Errorf("step %q not found", prevID)
	}

	// Skip disabled steps going backward
	for w.stepStatus[w.current.Steps[prevIdx].ID] == StepDisabled {
		prevIdx--
		if prevIdx < 0 {
			w.mu.Unlock()
			return "", errors.New("already at first step")
		}
		prevID = w.current.Steps[prevIdx].ID
	}

	// Update status
	w.stepStatus[oldStepID] = StepNotStarted
	w.stepStatus[prevID] = StepCurrent
	w.currentIdx = prevIdx

	// Update context
	w.ctx.Runtime.CurrentStep = prevID

	bus := w.bus
	w.mu.Unlock()

	// Emit event after unlocking to avoid deadlocks in handlers.
	if bus != nil {
		bus.PublishStepChange(oldStepID, prevID)
	}

	return prevID, nil
}

// JumpTo moves directly to a step (if allowed).
func (w *Workflow) JumpTo(stepID string) error {
	w.mu.Lock()

	if w.current == nil {
		w.mu.Unlock()
		return errors.New("no flow selected")
	}

	targetIdx, ok := w.stepIndex[stepID]
	if !ok {
		w.mu.Unlock()
		return fmt.Errorf("step %q not found", stepID)
	}

	targetStep := w.current.Steps[targetIdx]

	// Check if jump is allowed
	if !targetStep.AllowJump && !w.visited[stepID] {
		w.mu.Unlock()
		return fmt.Errorf("cannot jump to unvisited step %q", stepID)
	}

	// Check if step is disabled
	if w.stepStatus[stepID] == StepDisabled {
		w.mu.Unlock()
		return fmt.Errorf("step %q is disabled", stepID)
	}

	oldStepID := w.current.Steps[w.currentIdx].ID

	// Update status
	if w.stepStatus[oldStepID] != StepCompleted {
		w.stepStatus[oldStepID] = StepNotStarted
	}
	w.stepStatus[stepID] = StepCurrent
	w.currentIdx = targetIdx
	w.visited[stepID] = true

	// Update context
	w.ctx.Runtime.CurrentStep = stepID

	bus := w.bus
	w.mu.Unlock()

	// Emit event after unlocking to avoid deadlocks in handlers.
	if bus != nil {
		bus.PublishStepChange(oldStepID, stepID)
	}

	return nil
}

// DisableStep disables a step (it will be skipped).
func (w *Workflow) DisableStep(stepID string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if _, ok := w.stepIndex[stepID]; !ok {
		return fmt.Errorf("step %q not found", stepID)
	}

	// Cannot disable current step
	if w.current != nil && w.currentIdx >= 0 {
		if w.current.Steps[w.currentIdx].ID == stepID {
			return errors.New("cannot disable current step")
		}
	}

	w.stepStatus[stepID] = StepDisabled
	return nil
}

// EnableStep enables a previously disabled step.
func (w *Workflow) EnableStep(stepID string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if _, ok := w.stepIndex[stepID]; !ok {
		return fmt.Errorf("step %q not found", stepID)
	}

	if w.stepStatus[stepID] == StepDisabled {
		w.stepStatus[stepID] = StepNotStarted
	}
	return nil
}

// IsFirstStep returns true if current step is the first (non-disabled) step.
func (w *Workflow) IsFirstStep() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if w.current == nil || w.currentIdx < 0 {
		return true
	}

	// Check if there's any non-disabled step before current
	for i := 0; i < w.currentIdx; i++ {
		if w.stepStatus[w.current.Steps[i].ID] != StepDisabled {
			return false
		}
	}
	return true
}

// CanGoBack returns true if navigation to a previous step is possible.
func (w *Workflow) CanGoBack() bool {
	return !w.IsFirstStep()
}

// IsLastStep returns true if current step is the last (non-disabled) step.
func (w *Workflow) IsLastStep() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if w.current == nil || w.currentIdx < 0 {
		return true
	}

	// Check if there's any non-disabled step after current
	for i := w.currentIdx + 1; i < len(w.current.Steps); i++ {
		if w.stepStatus[w.current.Steps[i].ID] != StepDisabled {
			return false
		}
	}
	return true
}

// IsComplete returns true if the workflow has been completed.
func (w *Workflow) IsComplete() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.ctx.Runtime.Completed
}

// evaluateBranch evaluates a branch condition and returns the target step ID.
func (w *Workflow) evaluateBranch(branch *BranchConfig) string {
	// Simple implementation: check condition value in context
	val, ok := w.ctx.Get(branch.Condition)
	if !ok {
		return branch.Default
	}

	// Convert to string and look up in branches
	valStr := fmt.Sprintf("%v", val)
	if target, ok := branch.Branches[valStr]; ok {
		return target
	}

	return branch.Default
}

// Complete marks the workflow as completed.
func (w *Workflow) Complete() {
	w.mu.Lock()

	if w.current != nil && w.currentIdx >= 0 {
		w.stepStatus[w.current.Steps[w.currentIdx].ID] = StepCompleted
	}
	w.ctx.Runtime.Completed = true

	bus := w.bus
	w.mu.Unlock()

	// Emit event after unlocking to avoid deadlocks in handlers.
	if bus != nil {
		bus.Publish(Event{Type: EventFlowComplete})
	}
}
