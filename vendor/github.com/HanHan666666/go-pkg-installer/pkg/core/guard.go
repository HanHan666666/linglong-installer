// Package core provides guard implementations for workflow navigation.
package core

import (
	"errors"
	"fmt"
)

// MustAcceptGuard requires a field to be true (e.g., license accepted).
type MustAcceptGuard struct {
	Field string
	Msg   string
}

// NewMustAcceptGuard creates a MustAcceptGuard from config.
func NewMustAcceptGuard(config map[string]any) (Guard, error) {
	field, ok := config["field"].(string)
	if !ok || field == "" {
		return nil, errors.New("mustAccept guard requires 'field' property")
	}

	msg := "Please accept to continue"
	if m, ok := config["message"].(string); ok && m != "" {
		msg = m
	}

	return &MustAcceptGuard{Field: field, Msg: msg}, nil
}

func (g *MustAcceptGuard) Type() string    { return "mustAccept" }
func (g *MustAcceptGuard) Message() string { return g.Msg }

func (g *MustAcceptGuard) Check(ctx *InstallContext) error {
	val := ctx.GetBool(g.Field)
	if !val {
		return errors.New(g.Msg)
	}
	return nil
}

// DiskSpaceGuard requires minimum disk space.
type DiskSpaceGuard struct {
	MinMB int64
	Msg   string
}

// NewDiskSpaceGuard creates a DiskSpaceGuard from config.
func NewDiskSpaceGuard(config map[string]any) (Guard, error) {
	minMB := int64(0)

	switch v := config["minMB"].(type) {
	case int:
		minMB = int64(v)
	case int64:
		minMB = v
	case float64:
		minMB = int64(v)
	default:
		return nil, errors.New("diskSpace guard requires 'minMB' property as number")
	}

	if minMB <= 0 {
		return nil, errors.New("diskSpace guard requires positive 'minMB'")
	}

	msg := fmt.Sprintf("At least %d MB of free disk space is required", minMB)
	if m, ok := config["message"].(string); ok && m != "" {
		msg = m
	}

	return &DiskSpaceGuard{MinMB: minMB, Msg: msg}, nil
}

func (g *DiskSpaceGuard) Type() string    { return "diskSpace" }
func (g *DiskSpaceGuard) Message() string { return g.Msg }

func (g *DiskSpaceGuard) Check(ctx *InstallContext) error {
	available := ctx.Env.DiskFreeMB
	if available < g.MinMB {
		return fmt.Errorf("%s (available: %d MB)", g.Msg, available)
	}
	return nil
}

// FieldNotEmptyGuard requires a field to have a non-empty value.
type FieldNotEmptyGuard struct {
	Field string
	Msg   string
}

// NewFieldNotEmptyGuard creates a FieldNotEmptyGuard from config.
func NewFieldNotEmptyGuard(config map[string]any) (Guard, error) {
	field, ok := config["field"].(string)
	if !ok || field == "" {
		return nil, errors.New("fieldNotEmpty guard requires 'field' property")
	}

	msg := "This field is required"
	if m, ok := config["message"].(string); ok && m != "" {
		msg = m
	}

	return &FieldNotEmptyGuard{Field: field, Msg: msg}, nil
}

func (g *FieldNotEmptyGuard) Type() string    { return "fieldNotEmpty" }
func (g *FieldNotEmptyGuard) Message() string { return g.Msg }

func (g *FieldNotEmptyGuard) Check(ctx *InstallContext) error {
	val := ctx.GetString(g.Field)
	if val == "" {
		return errors.New(g.Msg)
	}
	return nil
}

// ExpressionGuard evaluates a simple expression.
type ExpressionGuard struct {
	Expression string
	Expected   any
	Msg        string
}

// NewExpressionGuard creates an ExpressionGuard from config.
func NewExpressionGuard(config map[string]any) (Guard, error) {
	expr, ok := config["expression"].(string)
	if !ok || expr == "" {
		return nil, errors.New("expression guard requires 'expression' property")
	}

	expected, ok := config["expected"]
	if !ok {
		expected = true // Default: expression field should be truthy
	}

	msg := "Condition not met"
	if m, ok := config["message"].(string); ok && m != "" {
		msg = m
	}

	return &ExpressionGuard{Expression: expr, Expected: expected, Msg: msg}, nil
}

func (g *ExpressionGuard) Type() string    { return "expression" }
func (g *ExpressionGuard) Message() string { return g.Msg }

func (g *ExpressionGuard) Check(ctx *InstallContext) error {
	val, ok := ctx.Get(g.Expression)
	if !ok {
		return fmt.Errorf("%s (field not found: %s)", g.Msg, g.Expression)
	}

	if fmt.Sprintf("%v", val) != fmt.Sprintf("%v", g.Expected) {
		return errors.New(g.Msg)
	}
	return nil
}

// RegisterBuiltinGuards registers all built-in guards with the global registry.
func RegisterBuiltinGuards() {
	Guards.Register("mustAccept", NewMustAcceptGuard)
	Guards.Register("diskSpace", NewDiskSpaceGuard)
	Guards.Register("fieldNotEmpty", NewFieldNotEmptyGuard)
	Guards.Register("expression", NewExpressionGuard)
}
