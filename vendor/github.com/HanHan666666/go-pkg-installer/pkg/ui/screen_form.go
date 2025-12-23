package ui

import (
	"errors"
	"fmt"
	"strings"

	. "modernc.org/tk9.0"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// FormScreen renders a form input screen with multiple field types.
type FormScreen struct {
	step   *core.StepConfig
	fields []formField
}

type formField struct {
	config   core.FieldConfig
	widget   any
	variable *VariableOpt
}

// NewFormScreen creates a form screen renderer.
func NewFormScreen(step *core.StepConfig) ScreenRenderer {
	return &FormScreen{step: step}
}

// Render creates the form screen UI.
func (s *FormScreen) Render(parent *TFrameWidget, ctx *core.InstallContext, bus *core.EventBus) error {
	// Title
	titleText := s.step.Screen.Title
	if titleText == "" {
		titleText = tr(ctx, "title.form", "Configuration")
	}
	titleText = ctx.Render(titleText)

	title := parent.TLabel(Txt(titleText), Font("TkHeadingFont"))
	Pack(title, Pady("10"), Side("top"))

	// Description
	if desc := s.step.Screen.Description; desc != "" {
		desc = ctx.Render(desc)
		descLabel := parent.TLabel(Txt(desc), Wraplength("600"))
		Pack(descLabel, Pady("10"), Side("top"))
	}

	// Form frame
	formFrame := parent.TFrame()
	Pack(formFrame, Fill("both"), Expand(true), Pady("10"))

	// Create fields
	s.fields = make([]formField, 0, len(s.step.Screen.Fields))

	for i, fieldConfig := range s.step.Screen.Fields {
		field := formField{config: fieldConfig}

		// Get existing value or default
		defaultVal := ctx.Render(fieldConfig.Default)
		if existing, ok := ctx.Get(fieldConfig.Variable); ok {
			defaultVal = fmt.Sprintf("%v", existing)
		}

		// Field row frame
		rowFrame := formFrame.TFrame()
		Pack(rowFrame, Fill("x"), Pady("5"))

		// Label
		labelText := fieldConfig.Label
		if fieldConfig.Required {
			labelText += " *"
		}
		label := rowFrame.TLabel(Txt(labelText), Width(20), Anchor("w"))
		Grid(label, Row(i), Column(0), Sticky("w"), Padx("5"))

		// Create widget based on type
		switch fieldConfig.Type {
		case "text", "string", "":
			entry := rowFrame.TEntry(
				Textvariable(defaultVal),
				Width(40),
			)
			Grid(entry, Row(i), Column(1), Sticky("ew"), Padx("5"))
			field.widget = entry

		case "password":
			entry := rowFrame.TEntry(
				Textvariable(defaultVal),
				Width(40),
				Show("*"),
			)
			Grid(entry, Row(i), Column(1), Sticky("ew"), Padx("5"))
			field.widget = entry

		case "directory", "path":
			pathFrame := rowFrame.TFrame()
			Grid(pathFrame, Row(i), Column(1), Sticky("ew"), Padx("5"))

			entry := pathFrame.TEntry(
				Textvariable(defaultVal),
				Width(35),
			)
			Pack(entry, Side("left"), Fill("x"), Expand(true))

			browseBtn := pathFrame.TButton(
				Txt("..."),
				Width(3),
				Style("Secondary.TButton"),
				Command(func() {
					dir := ChooseDirectory(
						Title("Select Directory"),
						Initialdir(entry.Textvariable()),
					)
					if dir != "" {
						entry.Configure(Textvariable(dir))
					}
				}),
			)
			Pack(browseBtn, Side("right"), Padx("2"))
			field.widget = entry

		case "file":
			fileFrame := rowFrame.TFrame()
			Grid(fileFrame, Row(i), Column(1), Sticky("ew"), Padx("5"))

			entry := fileFrame.TEntry(
				Textvariable(defaultVal),
				Width(35),
			)
			Pack(entry, Side("left"), Fill("x"), Expand(true))

			browseBtn := fileFrame.TButton(
				Txt("..."),
				Width(3),
				Style("Secondary.TButton"),
				Command(func() {
					files := GetOpenFile(
						Title("Select File"),
					)
					if len(files) > 0 {
						entry.Configure(Textvariable(files[0]))
					}
				}),
			)
			Pack(browseBtn, Side("right"), Padx("2"))
			field.widget = entry

		case "checkbox", "bool":
			field.variable = Variable("")
			if defaultVal == "true" || defaultVal == "1" {
				field.variable.Set("1")
			} else {
				field.variable.Set("")
			}
			check := rowFrame.TCheckbutton(
				Txt(""),
				Variable(field.variable),
			)
			Grid(check, Row(i), Column(1), Sticky("w"), Padx("5"))
			field.widget = check

		case "select", "dropdown", "combo":
			// Get options as individual arguments
			combo := rowFrame.TCombobox(
				Textvariable(defaultVal),
				Width(38),
				State("readonly"),
			)
			// Set values using Configure
			optionStrs := make([]string, 0)
			for _, opt := range fieldConfig.Options {
				optionStrs = append(optionStrs, opt.Value)
			}
			if len(optionStrs) > 0 {
				combo.Configure(Values(optionStrs))
			}
			Grid(combo, Row(i), Column(1), Sticky("ew"), Padx("5"))
			field.widget = combo

		case "radio":
			radioFrame := rowFrame.TFrame()
			Grid(radioFrame, Row(i), Column(1), Sticky("w"), Padx("5"))

			field.variable = Variable(defaultVal)
			for j, opt := range fieldConfig.Options {
				radio := radioFrame.TRadiobutton(
					Txt(opt.Label),
					Variable(field.variable),
					Value(opt.Value),
				)
				Grid(radio, Row(0), Column(j), Padx("5"))
			}
			field.widget = radioFrame

		default:
			// Default to text entry
			entry := rowFrame.TEntry(
				Textvariable(defaultVal),
				Width(40),
			)
			Grid(entry, Row(i), Column(1), Sticky("ew"), Padx("5"))
			field.widget = entry
		}

		// Hint text
		if fieldConfig.Hint != "" {
			hint := formFrame.TLabel(Txt(ctx.Render(fieldConfig.Hint)), Foreground("gray"))
			Pack(hint, Side("top"), Anchor("w"), Padx("130"))
		}

		s.fields = append(s.fields, field)
	}

	return nil
}

// Validate validates all form fields.
func (s *FormScreen) Validate() error {
	var errs []string

	for _, field := range s.fields {
		value := s.getFieldValue(field)

		// Check required
		if field.config.Required && value == "" {
			errs = append(errs, fmt.Sprintf("%s is required", field.config.Label))
			continue
		}

		// Check validation pattern (simplified)
		// In production, use regex for field.config.Validation
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}
	return nil
}

func (s *FormScreen) getFieldValue(field formField) string {
	if field.variable == nil {
		if getter, ok := field.widget.(interface{ Textvariable() string }); ok {
			return getter.Textvariable()
		}
		return ""
	}

	switch field.config.Type {
	case "checkbox", "bool":
		if field.variable.Get() == "1" {
			return "true"
		}
		return "false"
	default:
		return field.variable.Get()
	}
}

// Collect collects all form data into the context.
func (s *FormScreen) Collect(ctx *core.InstallContext) error {
	for _, field := range s.fields {
		value := s.getFieldValue(field)

		// Convert types as needed
		switch field.config.Type {
		case "checkbox", "bool":
			ctx.Set(field.config.Variable, value == "true")
		default:
			ctx.Set(field.config.Variable, value)
		}
	}
	return nil
}

// Cleanup cleans up the form screen resources.
func (s *FormScreen) Cleanup() {}

// Type returns the screen type identifier.
func (s *FormScreen) Type() string {
	return "form"
}
