package ui

import (
	"errors"
	"fmt"
	"sync"
	"time"

	. "modernc.org/tk9.0"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// ProgressScreen renders an installation progress screen.
type ProgressScreen struct {
	mu sync.Mutex

	step        *core.StepConfig
	progressBar *TProgressbarWidget
	progressVar *VariableOpt
	statusLabel *TLabelWidget
	logText     *TextWidget
	isComplete  bool
	taskRunner  *core.TaskRunner
	ctx         *core.InstallContext
	bus         *core.EventBus
	active      bool
}

// NewProgressScreen creates a progress screen renderer.
func NewProgressScreen(step *core.StepConfig) ScreenRenderer {
	return &ProgressScreen{step: step}
}

// Render creates the progress screen UI.
func (s *ProgressScreen) Render(parent *TFrameWidget, ctx *core.InstallContext, bus *core.EventBus) error {
	s.ctx = ctx
	s.bus = bus
	s.active = true

	// Title
	titleText := s.step.Screen.Title
	if titleText == "" {
		titleText = tr(ctx, "title.progress", "Installing...")
	}
	titleText = ctx.Render(titleText)

	title := parent.TLabel(Txt(titleText), Font("TkHeadingFont"))
	Pack(title, Pady("10"), Side("top"))

	// Description
	desc := s.step.Screen.Description
	if desc == "" {
		desc = "Please wait while the installation completes."
	}
	desc = ctx.Render(desc)

	descLabel := parent.TLabel(Txt(desc), Wraplength("600"))
	Pack(descLabel, Pady("10"), Side("top"))

	// Status label
	s.statusLabel = parent.TLabel(Txt(tr(ctx, "status.prepare", "Preparing...")))
	Pack(s.statusLabel, Pady("5"), Side("top"))

	// Progress bar
	s.progressVar = Variable(0.0)
	s.progressBar = parent.TProgressbar(
		Length(500),
		Mode("determinate"),
		s.progressVar,
		Maximum(100),
	)
	Pack(s.progressBar, Pady("10"), Side("top"))

	// Log frame with scrollbar
	logFrame := parent.TFrame()
	Pack(logFrame, Fill("both"), Expand(true), Pady("10"))

	scrollbar := logFrame.TScrollbar()
	Pack(scrollbar, Side("right"), Fill("y"))

	s.logText = logFrame.Text(
		Width(80),
		Height(10),
		Wrap("word"),
		Yscrollcommand(func(e *Event) { e.ScrollSet(scrollbar) }),
		State("disabled"),
	)
	applyTextStyle(s.logText)
	scrollbar.Configure(Command(func(e *Event) { e.Yview(s.logText) }))
	Pack(s.logText, Side("left"), Fill("both"), Expand(true))

	// Start installation after a short delay
	go func() {
		time.Sleep(100 * time.Millisecond)
		s.startInstallation()
	}()

	return nil
}

// UpdateProgress updates the progress bar and status.
func (s *ProgressScreen) UpdateProgress(percent float64, status string) {
	PostEvent(func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		if !s.active {
			return
		}
		if s.progressVar != nil {
			s.progressVar.Set(percent)
		}
		if s.statusLabel != nil && status != "" {
			s.statusLabel.Configure(Txt(status))
		}
	}, true)
}

// AddLogMessage adds a message to the log.
func (s *ProgressScreen) AddLogMessage(msg string) {
	PostEvent(func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		if !s.active {
			return
		}
		if s.logText == nil {
			return
		}
		s.logText.Configure(State("normal"))
		s.logText.Insert("end", msg+"\n")
		s.logText.See("end")
		s.logText.Configure(State("disabled"))
	}, false)
}

func (s *ProgressScreen) startInstallation() {
	// Get the task runner from context if available
	if runner, ok := s.ctx.Get("task_runner"); ok {
		s.taskRunner = runner.(*core.TaskRunner)
	} else {
		// Create a new task runner
		s.taskRunner = core.NewTaskRunner(s.ctx, s.bus)
		s.ctx.Set("task_runner", s.taskRunner)
	}

	// Get tasks from the step configuration
	tasks := s.step.Tasks
	if len(tasks) == 0 {
		// Check for tasks in the install block
		if installTasks, ok := s.ctx.Get("install.tasks"); ok {
			if t, ok := installTasks.([]core.TaskConfig); ok {
				tasks = t
			}
		}
	}

	// If still no tasks, show warning
	if len(tasks) == 0 {
		s.AddLogMessage(tr(s.ctx, "msg.no_tasks", "No tasks configured for installation."))
		s.UpdateProgress(100, tr(s.ctx, "status.no_tasks", "No tasks to run"))
		s.isComplete = true
		s.ctx.Set("install.complete", true)
		s.ctx.Set("install.failed", false)
		s.ctx.Set("step.failed", false)
		s.ctx.Set("step.failed_id", "")
		return
	}

	// Queue tasks
	for _, task := range tasks {
		if err := s.taskRunner.QueueConfig(task); err != nil {
			s.AddLogMessage(fmt.Sprintf("Failed to queue task: %v", err))
		}
	}

	// Subscribe to task events for logging
	totalTasks := len(tasks)
	completedTasks := 0

	s.bus.Subscribe(core.EventTaskStart, func(e core.Event) {
		if p := e.TaskPayload(); p != nil {
			s.AddLogMessage(fmt.Sprintf("Starting: %s", p.TaskID))
		}
	})

	s.bus.Subscribe(core.EventTaskComplete, func(e core.Event) {
		if p := e.TaskPayload(); p != nil {
			completedTasks++
			progress := float64(completedTasks) / float64(totalTasks) * 100
			s.UpdateProgress(progress, fmt.Sprintf("Completed: %s", p.TaskID))
			s.AddLogMessage(fmt.Sprintf("✓ Completed: %s", p.TaskID))
		}
	})

	s.bus.Subscribe(core.EventTaskError, func(e core.Event) {
		if p := e.TaskPayload(); p != nil {
			s.AddLogMessage(fmt.Sprintf("✗ Error in %s: %v", p.TaskID, p.Error))
		}
	})

	// Run tasks
	go func() {
		err := s.taskRunner.Run()
		if err != nil {
			s.ctx.Set("install.complete", false)
			s.ctx.Set("install.failed", true)
			s.AddLogMessage("\n" + fmt.Sprintf(tr(s.ctx, "msg.install.failed", "Installation failed: %v"), err))
			s.UpdateProgress(0, tr(s.ctx, "status.failed", "Installation Failed"))
			s.isComplete = true
			s.ctx.Set("step.failed", true)
			s.ctx.Set("step.failed_id", s.step.ID)
			if s.bus != nil {
				s.bus.PublishStepFailure(s.step.ID)
			}
		} else {
			s.ctx.Set("install.complete", true)
			s.ctx.Set("install.failed", false)
			s.UpdateProgress(100, tr(s.ctx, "status.complete", "Installation Complete"))
			s.AddLogMessage("\n" + tr(s.ctx, "msg.success", "Installation completed successfully!"))
			s.isComplete = true
			s.ctx.Set("step.failed", false)
			s.ctx.Set("step.failed_id", "")
		}
	}()
}

// Validate validates the progress screen (checks if complete).
func (s *ProgressScreen) Validate() error {
	if !s.isComplete {
		return errors.New(tr(s.ctx, "msg.install.in_progress", "Installation is still in progress"))
	}
	return nil
}

// Collect collects data from the progress screen.
func (s *ProgressScreen) Collect(ctx *core.InstallContext) error {
	if _, ok := ctx.Get("install.complete"); !ok {
		ctx.Set("install.complete", s.isComplete)
	}
	return nil
}

// Cleanup cleans up the progress screen resources.
func (s *ProgressScreen) Cleanup() {
	s.mu.Lock()
	s.active = false
	s.logText = nil
	s.statusLabel = nil
	s.progressVar = nil
	s.progressBar = nil
	runner := s.taskRunner
	complete := s.isComplete
	s.mu.Unlock()

	// Cancel any running tasks if the screen is closed
	if runner != nil && !complete {
		runner.Cancel()
	}
}

// Type returns the screen type identifier.
func (s *ProgressScreen) Type() string {
	return "progress"
}
