package ui

import (
	"errors"
	"sync"
	"time"

	. "modernc.org/tk9.0"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// DetectScreen runs step tasks on entry and displays a concise status message.
type DetectScreen struct {
	mu sync.Mutex

	step *core.StepConfig
	ctx  *core.InstallContext
	bus  *core.EventBus

	descriptionLabel *TLabelWidget
	contentLabel     *TLabelWidget

	taskRunner *core.TaskRunner
	active     bool
	isComplete bool
	failed     bool
	failErr    error

	descriptionText string
	contentText     string
}

// NewDetectScreen creates a detect screen renderer.
func NewDetectScreen(step *core.StepConfig) ScreenRenderer {
	return &DetectScreen{step: step}
}

// Render creates the detect screen UI.
func (s *DetectScreen) Render(parent *TFrameWidget, ctx *core.InstallContext, bus *core.EventBus) error {
	s.ctx = ctx
	s.bus = bus
	s.active = true

	titleText := s.step.Screen.Title
	if titleText == "" {
		if s.step.Title != "" {
			titleText = s.step.Title
		} else {
			titleText = tr(ctx, "title.detect", "Check System")
		}
	}
	titleText = ctx.Render(titleText)

	title := parent.TLabel(Txt(titleText), Font("TkHeadingFont"))
	Pack(title, Pady("10"), Side("top"))

	desc := s.runningDescription()
	s.descriptionLabel = parent.TLabel(Txt(desc), Wraplength("600"), Justify("left"), Anchor("w"))
	Pack(s.descriptionLabel, Pady("6"), Side("top"), Fill("x"))

	s.contentLabel = parent.TLabel(Txt(""), Wraplength("600"), Justify("left"), Anchor("w"))
	Pack(s.contentLabel, Pady("8"), Side("top"), Fill("x"))

	go func() {
		time.Sleep(50 * time.Millisecond)
		s.runTasks()
	}()

	return nil
}

func (s *DetectScreen) runningDescription() string {
	desc := ""
	if s.step != nil && s.step.Screen != nil {
		desc = s.step.Screen.Description
	}
	if desc == "" {
		desc = tr(s.ctx, "msg.detect.running", "Detecting...")
	}
	return s.ctx.Render(desc)
}

func (s *DetectScreen) renderedContent() string {
	if s.step == nil || s.step.Screen == nil {
		return ""
	}
	if s.step.Screen.Content == "" {
		return ""
	}
	return s.ctx.Render(s.step.Screen.Content)
}

func (s *DetectScreen) runTasks() {
	s.mu.Lock()
	s.isComplete = false
	s.failed = false
	s.failErr = nil
	s.mu.Unlock()

	s.setDescription(s.runningDescription())
	s.setContent("")

	tasks := s.step.Tasks
	if len(tasks) == 0 {
		s.markSuccess()
		return
	}

	s.taskRunner = core.NewTaskRunner(s.ctx, s.bus)
	for _, task := range tasks {
		if err := s.taskRunner.QueueConfig(task); err != nil {
			s.markFailure(err)
			return
		}
	}

	if err := s.taskRunner.Run(); err != nil {
		s.markFailure(err)
		return
	}

	s.markSuccess()
}

func (s *DetectScreen) markSuccess() {
	s.mu.Lock()
	s.isComplete = true
	s.failed = false
	s.failErr = nil
	s.mu.Unlock()

	s.ctx.Set("step.failed", false)
	s.ctx.Set("step.failed_id", "")

	s.setDescription("")
	s.setContent(s.renderedContent())
}

func (s *DetectScreen) markFailure(err error) {
	s.mu.Lock()
	s.isComplete = true
	s.failed = true
	s.failErr = err
	s.mu.Unlock()

	s.ctx.Set("step.failed", true)
	s.ctx.Set("step.failed_id", s.step.ID)

	if s.bus != nil {
		s.bus.PublishStepFailure(s.step.ID)
	}

	s.setDescription("")
	s.setContent(tr(s.ctx, "msg.detect.failed", "Detection failed."))
}

func (s *DetectScreen) setDescription(text string) {
	s.mu.Lock()
	s.descriptionText = text
	label := s.descriptionLabel
	active := s.active
	s.mu.Unlock()

	if !active || label == nil {
		return
	}

	PostEvent(func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		if !s.active || s.descriptionLabel == nil {
			return
		}
		s.descriptionLabel.Configure(Txt(text))
	}, true)
}

func (s *DetectScreen) setContent(text string) {
	s.mu.Lock()
	s.contentText = text
	label := s.contentLabel
	active := s.active
	s.mu.Unlock()

	if !active || label == nil {
		return
	}

	PostEvent(func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		if !s.active || s.contentLabel == nil {
			return
		}
		s.contentLabel.Configure(Txt(text))
	}, true)
}

// Validate blocks navigation until tasks complete or fail.
func (s *DetectScreen) Validate() error {
	s.mu.Lock()
	complete := s.isComplete
	failed := s.failed
	err := s.failErr
	s.mu.Unlock()

	if !complete {
		return errors.New(tr(s.ctx, "msg.detect.in_progress", "Detection is still running"))
	}
	if failed {
		if err != nil {
			return err
		}
		return errors.New(tr(s.ctx, "msg.detect.failed", "Detection failed."))
	}
	return nil
}

// Collect collects data from the detect screen (nothing to collect).
func (s *DetectScreen) Collect(ctx *core.InstallContext) error {
	return nil
}

// Cleanup cleans up the detect screen resources.
func (s *DetectScreen) Cleanup() {
	s.mu.Lock()
	s.active = false
	runner := s.taskRunner
	complete := s.isComplete
	s.descriptionLabel = nil
	s.contentLabel = nil
	s.mu.Unlock()

	if runner != nil && !complete {
		runner.Cancel()
	}
}

// Type returns the screen type identifier.
func (s *DetectScreen) Type() string {
	return "detect"
}
