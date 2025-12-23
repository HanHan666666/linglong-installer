package ui

import (
	"errors"
	"os"
	"path/filepath"

	. "modernc.org/tk9.0"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// DirectoryScreen renders a directory selection screen.
type DirectoryScreen struct {
	step       *core.StepConfig
	pathEntry  *TEntryWidget
	varName    string
	defaultDir string
}

// NewDirectoryScreen creates a directory screen renderer.
func NewDirectoryScreen(step *core.StepConfig) ScreenRenderer {
	return &DirectoryScreen{step: step}
}

// Render creates the directory selection screen UI.
func (s *DirectoryScreen) Render(parent *TFrameWidget, ctx *core.InstallContext, bus *core.EventBus) error {
	// Title
	titleText := s.step.Screen.Title
	if titleText == "" {
		titleText = tr(ctx, "title.directory", "Select Installation Directory")
	}
	titleText = ctx.Render(titleText)

	title := parent.TLabel(Txt(titleText), Font("TkHeadingFont"))
	Pack(title, Pady("10"), Side("top"))

	// Description
	desc := s.step.Screen.Description
	if desc == "" {
		desc = tr(ctx, "desc.directory", "Choose the folder where you want to install the application.")
	}
	desc = ctx.Render(desc)

	descLabel := parent.TLabel(Txt(desc), Wraplength("600"))
	Pack(descLabel, Pady("10"), Side("top"))

	// Spacer
	spacer1 := parent.TFrame()
	Pack(spacer1, Fill("both"), Expand(true))

	// Directory selection frame
	dirFrame := parent.TFrame()
	Pack(dirFrame, Fill("x"), Pady("20"))

	label := dirFrame.TLabel(Txt(tr(ctx, "label.install.dir", "Installation Directory:")))
	Pack(label, Side("left"), Padx("5"))

	// Get variable name and default value
	s.varName = "install.dir"
	if s.step.Screen.Bind != "" {
		s.varName = s.step.Screen.Bind
	}
	if len(s.step.Screen.Fields) > 0 {
		for _, field := range s.step.Screen.Fields {
			if field.Type == "directory" || field.Type == "path" {
				s.varName = field.Variable
				if field.Default != "" {
					s.defaultDir = ctx.Render(field.Default)
				}
				break
			}
		}
	}

	// Default install directory
	if s.defaultDir == "" {
		homeDir, _ := os.UserHomeDir()
		productName := ctx.RenderOrDefault("product.name", "app")
		s.defaultDir = filepath.Join(homeDir, productName)
	}

	// Check if we have a previously set value
	if existingDir, ok := ctx.Get(s.varName); ok {
		s.defaultDir = existingDir.(string)
	}

	// Path entry with text variable binding
	s.pathEntry = dirFrame.TEntry(
		Textvariable(s.defaultDir),
		Width(50),
	)
	s.pathEntry.Configure(Textvariable(s.defaultDir))
	Pack(s.pathEntry, Side("left"), Padx("5"), Fill("x"), Expand(true))

	// Browse button
	browseBtn := dirFrame.TButton(
		Txt(tr(ctx, "button.browse", "Browse...")),
		Style("Secondary.TButton"),
		Command(func() {
			dir := ChooseDirectory(
				Title(tr(ctx, "title.directory", "Select Installation Directory")),
				Initialdir(s.pathEntry.Textvariable()),
			)
			if dir != "" {
				s.pathEntry.Configure(Textvariable(dir))
			}
		}),
	)
	Pack(browseBtn, Side("right"), Padx("5"))

	// Space requirements info
	infoFrame := parent.TFrame()
	Pack(infoFrame, Fill("x"), Pady("10"))

	// Show disk space info if available
	requiredSpace := ctx.RenderOrDefault("install.required_space", "")
	if requiredSpace != "" {
		spaceLabel := infoFrame.TLabel(Txt(tr(ctx, "label.required.space", "Required space: ") + requiredSpace))
		Pack(spaceLabel, Side("left"))
	}

	// Spacer
	spacer2 := parent.TFrame()
	Pack(spacer2, Fill("both"), Expand(true))

	return nil
}

// Validate validates the selected directory.
func (s *DirectoryScreen) Validate() error {
	dir := s.pathEntry.Textvariable()
	if dir == "" {
		return errors.New("please select an installation directory")
	}

	// Check if parent directory exists or can be created
	parentDir := filepath.Dir(dir)
	if _, err := os.Stat(parentDir); os.IsNotExist(err) {
		// Try to check if we can create it
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			return errors.New("cannot create installation directory: " + err.Error())
		}
		// Clean up test directory
		os.Remove(parentDir)
	}

	// Check write permissions on parent
	if info, err := os.Stat(parentDir); err == nil {
		if !info.IsDir() {
			return errors.New("parent path is not a directory")
		}
	}

	return nil
}

// Collect collects the selected directory.
func (s *DirectoryScreen) Collect(ctx *core.InstallContext) error {
	dir := s.pathEntry.Textvariable()
	ctx.Set(s.varName, dir)
	// Also set as install_dir for convenience
	ctx.Set("install_dir", dir)
	return nil
}

// Cleanup cleans up the directory screen resources.
func (s *DirectoryScreen) Cleanup() {}

// Type returns the screen type identifier.
func (s *DirectoryScreen) Type() string {
	return "directory"
}
