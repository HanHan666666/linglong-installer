package ui

import (
	"fmt"
	"os"

	. "modernc.org/tk9.0"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// WelcomeScreen renders a welcome screen with product info.
type WelcomeScreen struct {
	step *core.StepConfig
}

// NewWelcomeScreen creates a welcome screen renderer.
func NewWelcomeScreen(step *core.StepConfig) ScreenRenderer {
	return &WelcomeScreen{step: step}
}

// Render creates the welcome screen UI.
func (s *WelcomeScreen) Render(parent *TFrameWidget, ctx *core.InstallContext, bus *core.EventBus) error {
	// Get product info
	productName := ctx.RenderOrDefault("product.name", "Application")
	// version := ctx.RenderOrDefault("product.version", "")

	// Title
	titleText := s.step.Screen.Title
	if titleText == "" {
		titleText = fmt.Sprintf(tr(ctx, "title.welcome", "Welcome to %s"), productName)
	}
	titleText = ctx.Render(titleText)

	title := parent.TLabel(Txt(titleText), Font("TkHeadingFont"))
	Pack(title, Pady("20"), Side("top"))

	// // Description
	// description := s.step.Screen.Description
	// if description == "" {
	// 	description = fmt.Sprintf(tr(ctx, "desc.welcome", "This will install %s on your computer."), displayName)
	// 	if version != "" {
	// 		description = fmt.Sprintf(tr(ctx, "desc.welcome.version", "This will install %s version %s on your computer."), displayName, version)
	// 	}
	// }
	// description = trText(ctx, description)

	// descLabel := parent.TLabel(Txt(description), Wraplength("600"))
	// Pack(descLabel, Pady("10"), Side("top"))

	// Optional content block
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
	if content != "" {
		contentLabel := parent.TLabel(Txt(content), Wraplength("600"))
		Pack(contentLabel, Pady("5"), Side("top"))
	}

	// Optional banner or logo (if configured)
	if bannerPath := s.step.Screen.BannerPath; bannerPath != "" {
		bannerPath = ctx.Render(bannerPath)
		// Note: Image loading would be done here if tk9 supports it
	}

	// Spacer
	spacer := parent.TFrame()
	Pack(spacer, Fill("both"), Expand(true))

	// Footer message
	footer := parent.TLabel(Txt(tr(ctx, "footer.welcome", "Click 'Continue' to proceed with the installation.")))
	Pack(footer, Pady("20"), Side("bottom"))

	return nil
}

// Validate validates the welcome screen (always valid).
func (s *WelcomeScreen) Validate() error {
	return nil
}

// Collect collects data from the welcome screen (nothing to collect).
func (s *WelcomeScreen) Collect(ctx *core.InstallContext) error {
	return nil
}

// Cleanup cleans up the welcome screen resources.
func (s *WelcomeScreen) Cleanup() {}

// Type returns the screen type identifier.
func (s *WelcomeScreen) Type() string {
	return "welcome"
}
