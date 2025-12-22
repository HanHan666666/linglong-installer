// Package ui provides the tk9-based UI adapter for the installer.
package ui

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	. "modernc.org/tk9.0"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// InstallerWindow represents the main installer window.
type InstallerWindow struct {
	mu sync.Mutex

	ctx      *core.InstallContext
	bus      *core.EventBus
	workflow *core.Workflow

	// UI state
	mainFrame    *TFrameWidget
	sidebarFrame *TFrameWidget
	contentFrame *TFrameWidget
	navFrame     *TFrameWidget

	// Navigation buttons
	backBtn   *TButtonWidget
	nextBtn   *TButtonWidget
	cancelBtn *TButtonWidget

	// Current screen
	currentScreen ScreenRenderer

	// Screen registry
	screenRenderers map[string]ScreenRendererFactory

	// Sidebar logo
	logoImage *Img
	logoPath  string

	// Callbacks
	onComplete func()
	onCancel   func()
}

const (
	sidebarWidth    = 180
	sidebarLogoSize = sidebarWidth
)

// ScreenRenderer interface for rendering different screen types.
type ScreenRenderer interface {
	// Render creates the screen UI in the given frame
	Render(parent *TFrameWidget, ctx *core.InstallContext, bus *core.EventBus) error
	// Validate checks if the screen input is valid
	Validate() error
	// Collect gathers input data into the context
	Collect(ctx *core.InstallContext) error
	// Cleanup removes any screen-specific resources
	Cleanup()
	// Type returns the screen type identifier
	Type() string
}

// ScreenRendererFactory creates a ScreenRenderer for a step configuration.
type ScreenRendererFactory func(step *core.StepConfig) ScreenRenderer

// NewInstallerWindow creates a new installer window.
func NewInstallerWindow(ctx *core.InstallContext, workflow *core.Workflow, bus *core.EventBus) *InstallerWindow {
	w := &InstallerWindow{
		ctx:             ctx,
		bus:             bus,
		workflow:        workflow,
		screenRenderers: make(map[string]ScreenRendererFactory),
	}

	// Register built-in screen renderers
	// Each screen type is registered automatically by creating a sample instance
	// and reading its Type() method
	w.autoRegister(NewWelcomeScreen)
	w.autoRegister(NewLicenseScreen)
	w.autoRegister(NewDirectoryScreen)
	w.autoRegister(NewProgressScreen)
	w.autoRegister(NewDetectScreen)
	w.autoRegister(NewRichtextScreen)
	w.autoRegister(NewSummaryScreen)
	w.autoRegister(NewFormScreen)
	w.autoRegister(NewOptionsScreen)

	// Register aliases for convenience
	w.RegisterScreenRenderer("pathPicker", NewDirectoryScreen)
	w.RegisterScreenRenderer("installType", NewOptionsScreen)
	w.RegisterScreenRenderer("finish", NewSummaryScreen)

	return w
}

// autoRegister automatically registers a screen factory by creating a sample instance
// and reading its Type() method.
func (w *InstallerWindow) autoRegister(factory ScreenRendererFactory) {
	// Create a sample instance with nil step to get the type
	sample := factory(nil)
	screenType := sample.Type()
	w.screenRenderers[screenType] = factory
}

// RegisterScreenRenderer registers a screen renderer factory.
func (w *InstallerWindow) RegisterScreenRenderer(screenType string, factory ScreenRendererFactory) {
	w.screenRenderers[screenType] = factory
}

// OnComplete sets the callback for installation completion.
func (w *InstallerWindow) OnComplete(fn func()) {
	w.onComplete = fn
}

// OnCancel sets the callback for installation cancellation.
func (w *InstallerWindow) OnCancel(fn func()) {
	w.onCancel = fn
}

// Run initializes and runs the installer UI.
func (w *InstallerWindow) Run() error {
	// Lock to OS thread for tk9
	runtime.LockOSThread()

	// Set up the main window
	w.setupWindow()

	// Subscribe to events
	w.subscribeEvents()

	// Render the first screen
	w.renderCurrentStep()

	// Start the main loop
	App.Wait()

	return nil
}

func (w *InstallerWindow) setupWindow() {
	// Get product info from context
	productName, _ := w.ctx.Get("product.name")
	if productName == nil {
		productName = "Installer"
	}

	applyTheme(w.ctx)

	// Configure main window title and size
	App.WmTitle(fmt.Sprintf("%v", productName))
	WmGeometry(App, "700x500")
	WmMinSize(App, 700, 500)

	// Create main container frame
	w.mainFrame = TFrame(Padding("12"), Style("Main.TFrame"))
	Pack(w.mainFrame, Fill("both"), Expand(true))

	// Create sidebar frame (for steps/branding)
	w.sidebarFrame = w.mainFrame.TFrame(Width(sidebarWidth), Padding("10"), Style("Sidebar.TFrame"))
	Pack(w.sidebarFrame, Side("left"), Fill("y"), Padx("6"))

	// Create content wrapper frame (content + nav)
	contentWrapper := w.mainFrame.TFrame(Style("Content.TFrame"))
	Pack(contentWrapper, Side("right"), Fill("both"), Expand(true))

	// Create content frame (for screen content)
	w.contentFrame = contentWrapper.TFrame(Padding("12"), Style("Content.TFrame"))

	// Create separator
	separator := contentWrapper.TSeparator()

	// Create navigation frame
	w.navFrame = contentWrapper.TFrame(Style("Nav.TFrame"))
	Grid(w.contentFrame, Row(0), Column(0), Sticky("nsew"))
	Grid(separator, Row(1), Column(0), Sticky("ew"), Pady("8"))
	Grid(w.navFrame, Row(2), Column(0), Sticky("ew"))
	GridRowConfigure(contentWrapper, 0, Weight(1))
	GridRowConfigure(contentWrapper, 1, Weight(0))
	GridRowConfigure(contentWrapper, 2, Weight(0))
	GridColumnConfigure(contentWrapper, 0, Weight(1))

	// Create navigation buttons
	w.cancelBtn = w.navFrame.TButton(Txt(tr(w.ctx, "button.cancel", "Cancel")), Command(w.handleCancel))
	Pack(w.cancelBtn, Side("left"), Padx("5"))

	w.nextBtn = w.navFrame.TButton(Txt(tr(w.ctx, "button.continue", "Continue")), Command(w.handleNext))
	Pack(w.nextBtn, Side("right"), Padx("5"))

	w.backBtn = w.navFrame.TButton(Txt(tr(w.ctx, "button.back", "Go Back")), Command(w.handleBack))
	Pack(w.backBtn, Side("right"), Padx("5"))

	primaryStyle, secondaryStyle, tertiaryStyle := navButtonStyles()
	w.cancelBtn.Configure(Style(tertiaryStyle))
	w.backBtn.Configure(Style(secondaryStyle))
	w.nextBtn.Configure(Style(primaryStyle))
}

func (w *InstallerWindow) subscribeEvents() {
	// Subscribe to progress events
	w.bus.Subscribe(core.EventProgress, func(e core.Event) {
		// Update progress bar if on progress screen
		if ps, ok := w.currentScreen.(*ProgressScreen); ok {
			if p := e.ProgressPayload(); p != nil {
				ps.UpdateProgress(p.Progress*100, p.Message)
			}
		}
	})

	// Subscribe to log events
	w.bus.Subscribe(core.EventLog, func(e core.Event) {
		if ps, ok := w.currentScreen.(*ProgressScreen); ok {
			if p := e.LogPayload(); p != nil {
				ps.AddLogMessage(p.Message)
			}
		}
	})

	// Subscribe to step change events
	w.bus.Subscribe(core.EventStepChange, func(e core.Event) {
		// Render the new step
		w.renderCurrentStep()
	})

	// Subscribe to step failure events
	w.bus.Subscribe(core.EventStepFailure, func(e core.Event) {
		PostEvent(func() {
			w.updateNavButtons()
		}, true)
	})
}

func (w *InstallerWindow) renderCurrentStep() {
	w.mu.Lock()
	defer w.mu.Unlock()

	// Clean up current screen
	if w.currentScreen != nil {
		w.currentScreen.Cleanup()
	}

	// Clear content frame by destroying children
	children := WinfoChildren(w.contentFrame.Window)
	for _, child := range children {
		Destroy(child)
	}

	// Get current step
	step := w.workflow.CurrentStep()
	if step == nil || step.Config == nil {
		return
	}

	stepConfig := step.Config

	// Get screen type
	screenType := ""
	if stepConfig.Screen != nil {
		screenType = stepConfig.Screen.Type
	}
	if core.IsGoExtension(screenType) {
		screenType = core.StripGoPrefix(screenType)
	}

	factory, ok := w.screenRenderers[screenType]
	if !ok {
		// Fallback to richtext for unknown types
		factory = NewRichtextScreen
	}

	// Create and render screen
	w.currentScreen = factory(stepConfig)
	if err := w.currentScreen.Render(w.contentFrame, w.ctx, w.bus); err != nil {
		w.ctx.AddLog(core.LogError, fmt.Sprintf("Failed to render screen: %v", err))
	}

	// Update sidebar and navigation buttons
	w.renderSidebar()
	w.updateNavButtons()
}

func (w *InstallerWindow) renderSidebar() {
	if w.sidebarFrame == nil {
		return
	}

	children := WinfoChildren(w.sidebarFrame.Window)
	for _, child := range children {
		Destroy(child)
	}

	productName := w.ctx.RenderOrDefault("product.name", "Installer")
	title := w.sidebarFrame.TLabel(
		Txt(productName),
		Font("TkHeadingFont"),
		Anchor("w"),
		Style("SidebarTitle.TLabel"),
	)
	Pack(title, Side("top"), Fill("x"), Pady("5"))

	steps := w.workflow.Steps()
	for _, step := range steps {
		status := w.workflow.StepStatus(step.ID)
		prefix := "○"
		style := "SidebarStep.TLabel"
		switch status {
		case core.StepCurrent:
			prefix = "▶"
			style = "SidebarActive.TLabel"
		case core.StepCompleted:
			prefix = "✅"
			style = "SidebarDone.TLabel"
		case core.StepDisabled:
			prefix = "𐄂"
			style = "SidebarDisabled.TLabel"
		}
		text := fmt.Sprintf("%s %s", prefix, step.Title)
		label := w.sidebarFrame.TLabel(
			Txt(text),
			Anchor("w"),
			Wraplength("160"),
			Style(style),
		)
		Pack(label, Side("top"), Fill("x"), Padx("5"), Pady("2"))
	}

	spacer := w.sidebarFrame.TFrame(Style("Sidebar.TFrame"))
	Pack(spacer, Fill("both"), Expand(true))

	logoKey := logoKeyFromContext(w.ctx)
	if logoKey == "" {
		w.logoImage = nil
		w.logoPath = ""
	} else if w.logoImage == nil || w.logoPath != logoKey {
		w.logoImage = loadLogoForKey(w.ctx, logoKey)
		if w.logoImage != nil {
			w.logoPath = logoKey
		} else {
			w.logoPath = ""
		}
	}

	if w.logoImage != nil {
		logoCanvas := w.sidebarFrame.Canvas(
			Width(sidebarLogoSize),
			Height(sidebarLogoSize),
			Background(currentPalette.sidebar),
			Borderwidth(0),
			Highlightthickness(0),
		)
		Pack(logoCanvas, Side("bottom"), Pady("6"))
		logoCanvas.CreateImage(sidebarLogoSize/2, sidebarLogoSize/2, Image(w.logoImage), Anchor("center"))
	}
}

func (w *InstallerWindow) updateNavButtons() {
	step := w.workflow.CurrentStep()
	if step == nil || step.Config == nil {
		return
	}

	stepConfig := step.Config
	nextStep := w.nextEnabledStep(step.ID)

	// Back button - disable on first step
	if w.workflow.CanGoBack() {
		w.backBtn.Configure(State("normal"))
	} else {
		w.backBtn.Configure(State("disabled"))
	}

	// Next button - change text based on screen type
	screenType := ""
	if stepConfig.Screen != nil {
		screenType = stepConfig.Screen.Type
	}
	if core.IsGoExtension(screenType) {
		screenType = core.StripGoPrefix(screenType)
	}

	nextScreenType := ""
	if nextStep != nil && nextStep.Config != nil && nextStep.Config.Screen != nil {
		nextScreenType = nextStep.Config.Screen.Type
	}
	if core.IsGoExtension(nextScreenType) {
		nextScreenType = core.StripGoPrefix(nextScreenType)
	}

	if isAnyStepFailed(w.ctx) {
		w.backBtn.Configure(State("disabled"))
		w.nextBtn.Configure(Txt(tr(w.ctx, "button.exit", "Exit")))
		return
	}

	if screenType == "progress" {
		w.nextBtn.Configure(Txt(tr(w.ctx, "button.finish", "Finish")))
	} else if screenType == "summary" || screenType == "finish" {
		if nextScreenType == "progress" {
			w.nextBtn.Configure(Txt(tr(w.ctx, "button.install", "Install")))
		} else if w.workflow.IsLastStep() || nextStep == nil {
			w.nextBtn.Configure(Txt(tr(w.ctx, "button.close", "Close")))
		} else {
			w.nextBtn.Configure(Txt(tr(w.ctx, "button.continue", "Continue")))
		}
	} else if w.workflow.IsLastStep() || nextStep == nil {
		w.nextBtn.Configure(Txt(tr(w.ctx, "button.close", "Close")))
	} else {
		w.nextBtn.Configure(Txt(tr(w.ctx, "button.continue", "Continue")))
	}
}

func (w *InstallerWindow) handleNext() {
	w.mu.Lock()
	screen := w.currentScreen
	w.mu.Unlock()

	step := w.workflow.CurrentStep()

	// Validate current screen
	if step != nil && isAnyStepFailed(w.ctx) {
		Destroy(App)
		os.Exit(1)
		return
	}

	if screen != nil {
		if err := screen.Validate(); err != nil {
			MessageBox(Icon("error"), Msg(err.Error()), Title(tr(w.ctx, "dialog.validation.title", "Validation Error")))
			return
		}

		// Collect data from screen
		if err := screen.Collect(w.ctx); err != nil {
			MessageBox(Icon("error"), Msg(err.Error()), Title(tr(w.ctx, "dialog.error.title", "Error")))
			return
		}
	}

	// Check if this is the last step or summary
	if w.workflow.IsLastStep() || w.nextEnabledStep(step.ID) == nil {
		if w.onComplete != nil {
			w.onComplete()
		}
		Destroy(App)
		return
	}

	_, err := w.workflow.Next()
	if err != nil {
		MessageBox(Icon("error"), Msg(err.Error()), Title(tr(w.ctx, "dialog.error.title", "Navigation Error")))
		return
	}

	// Render will be triggered by the step-change event.
	if w.bus == nil {
		w.renderCurrentStep()
	}
}

func (w *InstallerWindow) handleBack() {
	step := w.workflow.CurrentStep()
	if step != nil && isAnyStepFailed(w.ctx) {
		return
	}
	if !w.workflow.CanGoBack() {
		return
	}

	_, err := w.workflow.Prev()
	if err != nil {
		MessageBox(Icon("error"), Msg(err.Error()), Title(tr(w.ctx, "dialog.error.title", "Navigation Error")))
		return
	}

	// Render will be triggered by the step-change event.
	if w.bus == nil {
		w.renderCurrentStep()
	}
}

func (w *InstallerWindow) handleCancel() {
	result := MessageBox(
		Icon("question"),
		Msg(tr(w.ctx, "dialog.cancel.msg", "Are you sure you want to cancel the installation?")),
		Title(tr(w.ctx, "dialog.cancel.title", "Cancel Installation")),
		Type("yesno"),
	)

	if result == "yes" {
		if w.onCancel != nil {
			w.onCancel()
		}
		Destroy(App)
		os.Exit(0)
	}
}

func (w *InstallerWindow) nextEnabledStep(currentID string) *core.Step {
	steps := w.workflow.Steps()
	if len(steps) == 0 {
		return nil
	}
	index := -1
	for i, step := range steps {
		if step.ID == currentID {
			index = i
			break
		}
	}
	if index == -1 {
		return nil
	}
	for i := index + 1; i < len(steps); i++ {
		if w.workflow.StepStatus(steps[i].ID) != core.StepDisabled {
			return steps[i]
		}
	}
	return nil
}

func isAnyStepFailed(ctx *core.InstallContext) bool {
	if ctx == nil {
		return false
	}
	if failed, ok := ctx.Get("step.failed"); ok {
		if failed == true {
			return true
		}
	}
	if failed, ok := ctx.Get("install.failed"); ok {
		if failed == true {
			return true
		}
	}
	return false
}

// after schedules a function to run after a delay.
func after(d time.Duration, fn func()) {
	go func() {
		time.Sleep(d)
		fn()
	}()
}
