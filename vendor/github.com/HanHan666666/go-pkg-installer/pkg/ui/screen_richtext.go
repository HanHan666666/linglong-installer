package ui

import (
	"os"

	. "modernc.org/tk9.0"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// RichtextScreen renders a rich text/informational screen.
type RichtextScreen struct {
	step *core.StepConfig
}

// NewRichtextScreen creates a richtext screen renderer.
func NewRichtextScreen(step *core.StepConfig) ScreenRenderer {
	return &RichtextScreen{step: step}
}

// Render creates the richtext screen UI.
func (s *RichtextScreen) Render(parent *TFrameWidget, ctx *core.InstallContext, bus *core.EventBus) error {
	// Title
	titleText := s.step.Screen.Title
	if titleText == "" {
		titleText = tr(ctx, "title.info", "Information")
	}
	titleText = ctx.Render(titleText)

	title := parent.TLabel(Txt(titleText), Font("TkHeadingFont"))
	Pack(title, Pady("10"), Side("top"))

	// Description (if present)
	if desc := s.step.Screen.Description; desc != "" {
		desc = ctx.Render(desc)
		descLabel := parent.TLabel(Txt(desc), Wraplength("600"))
		Pack(descLabel, Pady("5"), Side("top"))
	}

	// Content frame with scrollbar
	contentFrame := parent.TFrame()
	Pack(contentFrame, Fill("both"), Expand(true), Pady("10"))

	scrollbar := contentFrame.TScrollbar()
	Pack(scrollbar, Side("right"), Fill("y"))

	text := contentFrame.Text(
		Width(80),
		Height(20),
		Wrap("word"),
		Yscrollcommand(func(e *Event) { e.ScrollSet(scrollbar) }),
	)
	applyTextStyle(text)
	scrollbar.Configure(Command(func(e *Event) { e.Yview(text) }))
	Pack(text, Side("left"), Fill("both"), Expand(true))

	// Load content
	content := ""
	if s.step.Screen.Content != "" {
		content = ctx.Render(s.step.Screen.Content)
	} else if s.step.Screen.ContentFile != "" {
		filePath := ctx.Render(s.step.Screen.ContentFile)
		if data, err := os.ReadFile(filePath); err == nil {
			content = string(data)
		} else {
			content = "Content would be loaded from: " + filePath
		}
	}

	// Insert content (supporting basic formatting)
	text.Configure(State("normal"))
	s.insertFormattedText(text, content)
	text.Configure(State("disabled"))

	return nil
}

// insertFormattedText inserts text with basic markdown-like formatting.
func (s *RichtextScreen) insertFormattedText(text *TextWidget, content string) {
	// For now, just insert plain text
	// In a full implementation, we could parse basic markdown
	text.Insert("1.0", content)
}

// Validate validates the richtext screen (always valid).
func (s *RichtextScreen) Validate() error {
	return nil
}

// Collect collects data from the richtext screen (nothing to collect).
func (s *RichtextScreen) Collect(ctx *core.InstallContext) error {
	return nil
}

// Cleanup cleans up the richtext screen resources.
func (s *RichtextScreen) Cleanup() {}

// Type returns the screen type identifier.
func (s *RichtextScreen) Type() string {
	return "richtext"
}
