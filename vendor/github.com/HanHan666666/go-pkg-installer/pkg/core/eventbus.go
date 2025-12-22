// Package core provides the event bus for decoupled communication.
package core

import (
	"sync"
)

// EventType represents the type of event.
type EventType string

const (
	// EventProgress is emitted when task progress changes.
	EventProgress EventType = "progress"
	// EventLog is emitted when a log message is added.
	EventLog EventType = "log"
	// EventStepChange is emitted when the current step changes.
	EventStepChange EventType = "step_change"
	// EventTaskStart is emitted when a task starts.
	EventTaskStart EventType = "task_start"
	// EventTaskComplete is emitted when a task completes.
	EventTaskComplete EventType = "task_complete"
	// EventTaskError is emitted when a task fails.
	EventTaskError EventType = "task_error"
	// EventStepFailure is emitted when a step fails.
	EventStepFailure EventType = "step_failure"
	// EventFlowComplete is emitted when the entire flow completes.
	EventFlowComplete EventType = "flow_complete"
)

// Event represents an event in the system.
type Event struct {
	Type    EventType
	Payload any
}

// ProgressPayload returns the payload as ProgressPayload, or nil.
func (e Event) ProgressPayload() *ProgressPayload {
	if p, ok := e.Payload.(*ProgressPayload); ok {
		return p
	}
	if p, ok := e.Payload.(ProgressPayload); ok {
		return &p
	}
	return nil
}

// TaskPayload returns the payload as TaskPayload, or nil.
func (e Event) TaskPayload() *TaskPayload {
	if p, ok := e.Payload.(*TaskPayload); ok {
		return p
	}
	if p, ok := e.Payload.(TaskPayload); ok {
		return &p
	}
	return nil
}

// LogPayload returns the payload as LogPayload, or nil.
func (e Event) LogPayload() *LogPayload {
	if p, ok := e.Payload.(*LogPayload); ok {
		return p
	}
	if p, ok := e.Payload.(LogPayload); ok {
		return &p
	}
	return nil
}

// StepChangePayload returns the payload as StepChangePayload, or nil.
func (e Event) StepChangePayload() *StepChangePayload {
	if p, ok := e.Payload.(*StepChangePayload); ok {
		return p
	}
	if p, ok := e.Payload.(StepChangePayload); ok {
		return &p
	}
	return nil
}

// StepFailurePayload returns the payload as StepFailurePayload, or nil.
func (e Event) StepFailurePayload() *StepFailurePayload {
	if p, ok := e.Payload.(*StepFailurePayload); ok {
		return p
	}
	if p, ok := e.Payload.(StepFailurePayload); ok {
		return &p
	}
	return nil
}

// ProgressPayload contains progress update data.
type ProgressPayload struct {
	TaskID   string
	Progress float64 // 0.0 to 1.0
	Message  string
}

// LogPayload contains log event data.
type LogPayload struct {
	Level   LogLevel
	Message string
}

// StepChangePayload contains step change data.
type StepChangePayload struct {
	FromStep string
	ToStep   string
}

// StepFailurePayload contains step failure data.
type StepFailurePayload struct {
	StepID string
}

// TaskPayload contains task event data.
type TaskPayload struct {
	TaskID   string
	TaskType string
	Error    error
}

// EventHandler is a function that handles events.
type EventHandler func(event Event)

// EventBus provides publish-subscribe functionality for events.
type EventBus struct {
	mu          sync.RWMutex
	handlers    map[EventType][]EventHandler
	allHandlers []EventHandler
}

// NewEventBus creates a new EventBus.
func NewEventBus() *EventBus {
	return &EventBus{
		handlers:    make(map[EventType][]EventHandler),
		allHandlers: make([]EventHandler, 0),
	}
}

// Subscribe registers a handler for a specific event type.
func (eb *EventBus) Subscribe(eventType EventType, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

// SubscribeAll registers a handler for all event types.
func (eb *EventBus) SubscribeAll(handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.allHandlers = append(eb.allHandlers, handler)
}

// Publish sends an event to all registered handlers.
func (eb *EventBus) Publish(event Event) {
	eb.mu.RLock()
	handlers := make([]EventHandler, 0, len(eb.handlers[event.Type])+len(eb.allHandlers))
	handlers = append(handlers, eb.handlers[event.Type]...)
	handlers = append(handlers, eb.allHandlers...)
	eb.mu.RUnlock()

	for _, handler := range handlers {
		handler(event)
	}
}

// PublishProgress is a convenience method for publishing progress events.
func (eb *EventBus) PublishProgress(taskID string, progress float64, message string) {
	eb.Publish(Event{
		Type: EventProgress,
		Payload: ProgressPayload{
			TaskID:   taskID,
			Progress: progress,
			Message:  message,
		},
	})
}

// PublishLog is a convenience method for publishing log events.
func (eb *EventBus) PublishLog(level LogLevel, message string) {
	eb.Publish(Event{
		Type: EventLog,
		Payload: LogPayload{
			Level:   level,
			Message: message,
		},
	})
}

// PublishStepChange is a convenience method for publishing step change events.
func (eb *EventBus) PublishStepChange(from, to string) {
	eb.Publish(Event{
		Type: EventStepChange,
		Payload: StepChangePayload{
			FromStep: from,
			ToStep:   to,
		},
	})
}

// PublishStepFailure is a convenience method for publishing step failure events.
func (eb *EventBus) PublishStepFailure(stepID string) {
	eb.Publish(Event{
		Type: EventStepFailure,
		Payload: StepFailurePayload{
			StepID: stepID,
		},
	})
}

// PublishTaskStart is a convenience method for publishing task start events.
func (eb *EventBus) PublishTaskStart(taskID, taskType string) {
	eb.Publish(Event{
		Type: EventTaskStart,
		Payload: TaskPayload{
			TaskID:   taskID,
			TaskType: taskType,
		},
	})
}

// PublishTaskComplete is a convenience method for publishing task complete events.
func (eb *EventBus) PublishTaskComplete(taskID, taskType string) {
	eb.Publish(Event{
		Type: EventTaskComplete,
		Payload: TaskPayload{
			TaskID:   taskID,
			TaskType: taskType,
		},
	})
}

// PublishTaskError is a convenience method for publishing task error events.
func (eb *EventBus) PublishTaskError(taskID, taskType string, err error) {
	eb.Publish(Event{
		Type: EventTaskError,
		Payload: TaskPayload{
			TaskID:   taskID,
			TaskType: taskType,
			Error:    err,
		},
	})
}

// Clear removes all handlers.
func (eb *EventBus) Clear() {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.handlers = make(map[EventType][]EventHandler)
	eb.allHandlers = make([]EventHandler, 0)
}
