package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
	"github.com/HanHan666666/go-pkg-installer/pkg/ui"
	tk "modernc.org/tk9.0"
)

// LinglongVersionScreen checks Linglong version and prompts to update if too low.
type LinglongVersionScreen struct {
	step         *core.StepConfig
	minVersion   string
	contentLabel *tk.TLabelWidget
}

func NewLinglongVersionScreen(step *core.StepConfig) ui.ScreenRenderer {
	return &LinglongVersionScreen{step: step}
}

func (s *LinglongVersionScreen) Render(parent *tk.TFrameWidget, ctx *core.InstallContext, bus *core.EventBus) error {
	_ = bus

	titleText := ""
	if s.step != nil {
		titleText = s.step.Title
		if s.step.Screen != nil && s.step.Screen.Title != "" {
			titleText = s.step.Screen.Title
		}
	}
	if titleText == "" {
		titleText = "检查玲珑版本"
	}
	titleText = ctx.Render(titleText)

	title := parent.TLabel(tk.Txt(titleText), tk.Font("TkHeadingFont"))
	tk.Pack(title, tk.Pady("10"), tk.Side("top"))

	if s.step != nil && s.step.Screen != nil && s.step.Screen.Description != "" {
		desc := ctx.Render(s.step.Screen.Description)
		descLabel := parent.TLabel(tk.Txt(desc), tk.Wraplength("600"))
		tk.Pack(descLabel, tk.Pady("6"), tk.Side("top"))
	}

	s.contentLabel = parent.TLabel(tk.Txt(""), tk.Wraplength("600"), tk.Justify("left"), tk.Anchor("w"))
	tk.Pack(s.contentLabel, tk.Pady("8"), tk.Side("top"), tk.Fill("x"))

	s.minVersion = "1.9.0"
	if val, ok := ctx.Get("linglong.min_version"); ok {
		if str := strings.TrimSpace(fmt.Sprintf("%v", val)); str != "" {
			s.minVersion = str
		}
	}
	ctx.Set("linglong.min_version", s.minVersion)

	version, ok := detectLinglongVersion(s.minVersion)
	ctx.Set("linglong.version", version)
	ctx.Set("linglong.version_ok", ok)

	if s.step != nil && s.step.Screen != nil && s.step.Screen.Content != "" {
		content := ctx.Render(s.step.Screen.Content)
		s.contentLabel.Configure(tk.Txt(content))
	}

	if !ok {
		prompt := fmt.Sprintf("检测到当前玲珑版本 %s 低于最低要求 %s。\n是否现在更新？", version, s.minVersion)
		result := tk.MessageBox(
			tk.Icon("question"),
			tk.Msg(prompt),
			tk.Title("玲珑版本过低"),
			tk.Type("yesno"),
		)
		if result != "yes" {
			tk.Destroy(tk.App)
			os.Exit(0)
		}
	}

	return nil
}

func (s *LinglongVersionScreen) Validate() error {
	return nil
}

func (s *LinglongVersionScreen) Collect(ctx *core.InstallContext) error {
	return nil
}

func (s *LinglongVersionScreen) Cleanup() {}

func (s *LinglongVersionScreen) Type() string {
	return "linglongVersion"
}

func detectLinglongVersion(minVersion string) (string, bool) {
	output, err := exec.Command("ll-cli", "--version").CombinedOutput()
	if err != nil {
		return "未检测到", false
	}

	raw := strings.TrimSpace(string(output))
	version, ok := extractVersion(raw)
	if !ok {
		return "未检测到", false
	}

	if versionLess(version, minVersion) {
		return version, false
	}
	return version, true
}

func extractVersion(raw string) (string, bool) {
	re := regexp.MustCompile(`\d+(?:\.\d+){1,2}`)
	match := re.FindString(raw)
	if match == "" {
		return "", false
	}

	parts := strings.Split(match, ".")
	if len(parts) < 2 {
		return "", false
	}

	nums := make([]int, 0, 3)
	for _, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			return "", false
		}
		nums = append(nums, n)
	}
	for len(nums) < 3 {
		nums = append(nums, 0)
	}

	return fmt.Sprintf("%d.%d.%d", nums[0], nums[1], nums[2]), true
}

func versionLess(a, b string) bool {
	av, ok := parseVersion(a)
	if !ok {
		return true
	}
	bv, ok := parseVersion(b)
	if !ok {
		return false
	}

	for i := 0; i < 3; i++ {
		if av[i] < bv[i] {
			return true
		}
		if av[i] > bv[i] {
			return false
		}
	}
	return false
}

func parseVersion(value string) ([3]int, bool) {
	var out [3]int
	version, ok := extractVersion(value)
	if !ok {
		return out, false
	}
	parts := strings.Split(version, ".")
	if len(parts) < 2 {
		return out, false
	}

	for i := 0; i < 3 && i < len(parts); i++ {
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			return out, false
		}
		out[i] = n
	}
	return out, true
}
