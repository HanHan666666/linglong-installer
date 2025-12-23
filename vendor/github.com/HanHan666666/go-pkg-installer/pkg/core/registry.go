// Package core provides the registry mechanism for extensible components.
package core

import (
	"fmt"
	"strings"
	"sync"
)

// Registry provides a type-safe registry for components.
type Registry[T any] struct {
	mu    sync.RWMutex
	items map[string]T
	name  string // For error messages
}

// NewRegistry creates a new registry with the given name.
func NewRegistry[T any](name string) *Registry[T] {
	return &Registry[T]{
		items: make(map[string]T),
		name:  name,
	}
}

// Register adds an item to the registry.
// Returns an error if the key already exists.
func (r *Registry[T]) Register(key string, item T) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.items[key]; exists {
		return fmt.Errorf("%s: key %q already registered", r.name, key)
	}
	r.items[key] = item
	return nil
}

// MustRegister adds an item to the registry, panicking if the key already exists.
func (r *Registry[T]) MustRegister(key string, item T) {
	if err := r.Register(key, item); err != nil {
		panic(err)
	}
}

// Get retrieves an item from the registry.
// Returns the item and true if found, zero value and false otherwise.
func (r *Registry[T]) Get(key string) (T, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, ok := r.items[key]
	return item, ok
}

// MustGet retrieves an item from the registry, panicking if not found.
func (r *Registry[T]) MustGet(key string) T {
	item, ok := r.Get(key)
	if !ok {
		panic(fmt.Sprintf("%s: key %q not found", r.name, key))
	}
	return item
}

// Has checks if a key exists in the registry.
func (r *Registry[T]) Has(key string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.items[key]
	return ok
}

// Keys returns all registered keys.
func (r *Registry[T]) Keys() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	keys := make([]string, 0, len(r.items))
	for k := range r.items {
		keys = append(keys, k)
	}
	return keys
}

// Clear removes all items from the registry.
func (r *Registry[T]) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.items = make(map[string]T)
}

// IsGoExtension checks if a type name uses the go: namespace prefix.
func IsGoExtension(typeName string) bool {
	return strings.HasPrefix(typeName, "go:")
}

// StripGoPrefix removes the go: prefix from a type name.
func StripGoPrefix(typeName string) string {
	return strings.TrimPrefix(typeName, "go:")
}

// TaskFactory creates a Task from configuration.
type TaskFactory func(config map[string]any, ctx *InstallContext) (Task, error)

// ScreenFactory creates a Screen from configuration.
type ScreenFactory func(config map[string]any) (Screen, error)

// GuardFactory creates a Guard from configuration.
type GuardFactory func(config map[string]any) (Guard, error)

// Task represents an executable installation task.
type Task interface {
	// ID returns the unique identifier for this task.
	ID() string

	// Type returns the task type name.
	Type() string

	// Validate checks if the task configuration is valid.
	Validate() error

	// Execute runs the task.
	Execute(ctx *InstallContext, bus *EventBus) error

	// Rollback undoes the task (if possible).
	Rollback(ctx *InstallContext, bus *EventBus) error

	// CanRollback returns true if the task supports rollback.
	CanRollback() bool
}

// Screen represents a wizard step screen.
type Screen interface {
	// ID returns the screen identifier.
	ID() string

	// Type returns the screen type name.
	Type() string

	// Validate checks if the screen configuration is valid.
	Validate() error

	// Bind returns the context path this screen binds to (if any).
	Bind() string
}

// Guard represents a condition that must be met to proceed.
type Guard interface {
	// Type returns the guard type name.
	Type() string

	// Check evaluates the guard condition.
	// Returns nil if the guard passes, an error describing why it failed otherwise.
	Check(ctx *InstallContext) error

	// Message returns a user-friendly message explaining the guard requirement.
	Message() string
}

// Global registries (initialized in init.go)
var (
	Tasks   *Registry[TaskFactory]
	Screens *Registry[ScreenFactory]
	Guards  *Registry[GuardFactory]
)

func init() {
	Tasks = NewRegistry[TaskFactory]("TaskRegistry")
	Screens = NewRegistry[ScreenFactory]("ScreenRegistry")
	Guards = NewRegistry[GuardFactory]("GuardRegistry")
}
