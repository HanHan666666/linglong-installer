package ui

import (
	"fmt"
	"strings"

	. "modernc.org/tk9.0"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// SummaryScreen renders an installation summary screen.
type SummaryScreen struct {
	step      *core.StepConfig
	launchVar *VariableOpt
}

// NewSummaryScreen creates a summary screen renderer.
func NewSummaryScreen(step *core.StepConfig) ScreenRenderer {
	return &SummaryScreen{step: step}
}

// Render creates the summary screen UI.
func (s *SummaryScreen) Render(parent *TFrameWidget, ctx *core.InstallContext, bus *core.EventBus) error {
	// Check installation status
	isComplete, ok := ctx.Get("install.complete")
	success := false
	hasResult := false
	if ok {
		if value, ok := isComplete.(bool); ok {
			success = value
			hasResult = true
		}
	}
	body := parent.TFrame()
	Pack(body, Fill("both"), Expand(true))
	row := 0

	// Title
	titleText := s.step.Screen.Title
	if titleText == "" {
		if !hasResult && s.step.Title != "" {
			titleText = s.step.Title
		} else if success {
			titleText = tr(ctx, "title.complete", "Installation Complete")
		} else {
			titleText = tr(ctx, "title.summary", "Installation Summary")
		}
	}
	titleText = ctx.Render(titleText)

	title := body.TLabel(Txt(titleText), Font("TkHeadingFont"))
	Grid(title, Row(row), Column(0), Sticky("w"), Pady("20"))
	row++

	// Status icon/message
	var statusText string
	if !hasResult {
		statusText = tr(ctx, "msg.ready", "Review the details before installing.")
	} else if success {
		statusText = tr(ctx, "msg.success", "✓ The installation completed successfully!")
	} else {
		statusText = tr(ctx, "msg.failure", "✗ The installation encountered errors.")
	}

	statusLabel := body.TLabel(Txt(statusText))
	Grid(statusLabel, Row(row), Column(0), Sticky("w"), Pady("10"))
	row++

	// Description
	desc := s.step.Screen.Description
	if desc == "" && success {
		productName := ctx.RenderOrDefault("product.name", "The application")
		desc = fmt.Sprintf("%s has been installed on your computer.", productName)
	}
	if desc != "" {
		desc = ctx.Render(desc)
		descLabel := body.TLabel(Txt(desc), Wraplength("600"))
		Grid(descLabel, Row(row), Column(0), Sticky("w"), Pady("10"))
		row++
	}

	// Summary text with scrollbar
	summaryFrame := body.TFrame()
	Grid(summaryFrame, Row(row), Column(0), Sticky("nsew"), Pady("10"))

	scrollbar := summaryFrame.TScrollbar()
	Pack(scrollbar, Side("right"), Fill("y"))

	summaryText := summaryFrame.Text(
		Width(80),
		Height(12),
		Wrap("word"),
		Yscrollcommand(func(e *Event) { e.ScrollSet(scrollbar) }),
	)
	applyTextStyle(summaryText)
	scrollbar.Configure(Command(func(e *Event) { e.Yview(summaryText) }))
	Pack(summaryText, Side("left"), Fill("both"), Expand(true))

	var lines []string

	// Show installation details
	if installDir, ok := ctx.Get("install_dir"); ok {
		lines = append(lines, fmt.Sprintf("%s %v", tr(ctx, "label.installed.to", "Installed to:"), installDir))
	}

	// Show task plan if available
	if ctx.Plan != nil && len(ctx.Plan.Tasks) > 0 {
		if len(lines) > 0 {
			lines = append(lines, "")
		}
		lines = append(lines, tr(ctx, "label.plan", "Planned actions:"))
		for _, item := range ctx.Plan.Tasks {
			line := fmt.Sprintf("- %s", item.Description)
			if item.RequiresRoot {
				line = fmt.Sprintf("%s (admin)", line)
			}
			lines = append(lines, line)
		}
	}

	// Show errors if any
	if len(ctx.Runtime.Errors) > 0 {
		if len(lines) > 0 {
			lines = append(lines, "")
		}
		lines = append(lines, tr(ctx, "label.errors", "Errors:"))
		for _, err := range ctx.Runtime.Errors {
			lines = append(lines, fmt.Sprintf("- %s", err.Error()))
		}
	}

	// Log file path
	if logPath := ctx.LogPath(); logPath != "" {
		if len(lines) > 0 {
			lines = append(lines, "")
		}
		lines = append(lines, fmt.Sprintf(tr(ctx, "label.logfile", "Log file: %s"), logPath))
	}

	summaryText.Configure(State("normal"))
	summaryText.Insert("1.0", strings.Join(lines, "\n"))
	summaryText.Configure(State("disabled"))

	GridRowConfigure(body, row, Weight(1))
	row++

	// Optional: Launch application checkbox
	if success {
		launchFrame := body.TFrame()
		Grid(launchFrame, Row(row), Column(0), Sticky("w"), Pady("8"))

		// Check if there's a launch command configured
		if launchCmd, ok := ctx.Get("launch.command"); ok && launchCmd != "" {
			s.launchVar = Variable("1")
			launchCheck := launchFrame.TCheckbutton(
				Txt(tr(ctx, "label.launch", "Launch application after closing")),
				Variable(s.launchVar),
			)
			Pack(launchCheck, Side("left"))

			// Store the preference
			ctx.Set("launch.on_close", true)
		}
		row++
	}

	// Footer
	footer := body.TLabel(Txt(tr(ctx, "footer.close", "Click 'Close' to exit the installer.")))
	Grid(footer, Row(row), Column(0), Sticky("w"), Pady("10"))
	GridColumnConfigure(body, 0, Weight(1))

	return nil
}

// Validate validates the summary screen (always valid).
func (s *SummaryScreen) Validate() error {
	return nil
}

// Collect collects data from the summary screen.
func (s *SummaryScreen) Collect(ctx *core.InstallContext) error {
	if s.launchVar != nil {
		ctx.Set("launch.on_close", s.launchVar.Get() == "1")
	}
	return nil
}

// Cleanup cleans up the summary screen resources.
func (s *SummaryScreen) Cleanup() {}

// Type returns the screen type identifier.
func (s *SummaryScreen) Type() string {
	return "summary"
}
