package ui

import (
	"strings"

	. "modernc.org/tk9.0"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// OptionsScreen renders a list of selectable options.
type OptionsScreen struct {
	step *core.StepConfig
	bind string
	vars map[string]*VariableOpt
}

// NewOptionsScreen creates an options screen renderer.
func NewOptionsScreen(step *core.StepConfig) ScreenRenderer {
	return &OptionsScreen{step: step}
}

// Render creates the options screen UI.
func (s *OptionsScreen) Render(parent *TFrameWidget, ctx *core.InstallContext, bus *core.EventBus) error {
	// Title
	titleText := s.step.Screen.Title
	if titleText == "" {
		titleText = tr(ctx, "title.options", "Installation Options")
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

	s.bind = s.step.Screen.Bind
	if s.bind == "" {
		s.bind = "options"
	}

	optionsFrame := parent.TFrame()
	Pack(optionsFrame, Fill("both"), Expand(true), Pady("10"))

	s.vars = make(map[string]*VariableOpt, len(s.step.Screen.Options))
	for i, opt := range s.step.Screen.Options {
		checked := opt.Default
		if existing, ok := ctx.Get(s.bind + "." + opt.Value); ok {
			if b, ok := existing.(bool); ok {
				checked = b
			}
		}

		varOpt := Variable("")
		check := optionsFrame.TCheckbutton(
			Txt(opt.Label),
			varOpt,
		)
		if checked {
			varOpt.Set("1")
		}
		s.vars[opt.Value] = varOpt
		Grid(check, Row(i), Column(0), Sticky("w"), Padx("5"), Pady("3"))
	}

	return nil
}

// Validate validates the options screen (always valid).
func (s *OptionsScreen) Validate() error {
	return nil
}

// Collect collects selected options into the context.
func (s *OptionsScreen) Collect(ctx *core.InstallContext) error {
	selected := make([]string, 0, len(s.vars))
	for value, varOpt := range s.vars {
		isOn := strings.EqualFold(strings.TrimSpace(varOpt.Get()), "1")
		ctx.Set(s.bind+"."+value, isOn)
		if isOn {
			selected = append(selected, value)
		}
	}
	ctx.Set(s.bind, selected)
	return nil
}

// Cleanup cleans up the options screen resources.
func (s *OptionsScreen) Cleanup() {}

// Type returns the screen type identifier.
func (s *OptionsScreen) Type() string {
	return "options"
}
