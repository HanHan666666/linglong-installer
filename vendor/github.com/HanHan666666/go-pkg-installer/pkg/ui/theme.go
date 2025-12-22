package ui

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

type uiPalette struct {
	appBg           string
	surface         string
	sidebar         string
	border          string
	text            string
	muted           string
	secondaryBorder string
	secondaryText   string
	tertiaryText    string
	accent          string
	accentHover     string
	accentActive    string
	buttonBg        string
	buttonHover     string
}

var (
	themeOnce      sync.Once
	currentPalette = defaultPalette()
	buttonFont     *FontFace
	useAzureTheme  bool
)

func applyTheme(ctx *core.InstallContext) {
	themeOnce.Do(func() {
		currentPalette = buildPalette(ctx)
		if safeActivateTheme("Azure light") || safeActivateTheme("azure light") {
			useAzureTheme = true
		}
		safeApplyFonts()
		safeApplyStyles(currentPalette)
		safeConfigureBackground(currentPalette.appBg)
	})
}

func applyTextStyle(text *TextWidget) {
	palette := currentPalette
	text.Configure(
		Background(palette.surface),
		Foreground(palette.text),
		Insertbackground(palette.text),
		Highlightthickness(1),
		Highlightbackground(palette.border),
		Highlightcolor(palette.accent),
		Borderwidth(1),
		Relief("solid"),
	)
}

func defaultPalette() uiPalette {
	return uiPalette{
		appBg:           "#f5f7fb",
		surface:         "#ffffff",
		sidebar:         "#f1f5f9",
		border:          "#e5e7eb",
		text:            "#111827",
		muted:           "#6b7280",
		secondaryBorder: "#d9d9d9",
		secondaryText:   "#595959",
		tertiaryText:    "#8c8c8c",
		accent:          "#1677ff",
		accentHover:     "#4096ff",
		accentActive:    "#0958d9",
		buttonBg:        "#f3f4f6",
		buttonHover:     "#e5e7eb",
	}
}

func buildPalette(ctx *core.InstallContext) uiPalette {
	palette := defaultPalette()
	if ctx == nil {
		return palette
	}

	accent := strings.TrimSpace(ctx.RenderOrDefault("theme.primaryColor", ""))
	if normalized, ok := normalizeHex(accent); ok {
		palette.accent = normalized
		palette.accentHover = adjustColor(normalized, 12)
		palette.accentActive = adjustColor(normalized, -18)
	}

	return palette
}

func applyFonts() {
	family := pickFontFamily()
	applyFont("TkDefaultFont", family, 10, "")
	applyFont("TkTextFont", family, 10, "")
	applyFont("TkHeadingFont", family, 14, "bold")
	applyFont("TkSmallCaptionFont", family, 9, "")
	if family != "" {
		buttonFont = NewFont(Family(family), Size(10), Weight("bold"))
	} else {
		buttonFont = NewFont(Size(10), Weight("bold"))
	}
}

func applyFont(name, family string, size int, weight string) {
	opts := []any{Size(size)}
	if family != "" {
		opts = append(opts, Family(family))
	}
	if weight != "" {
		opts = append(opts, Weight(weight))
	}
	FontConfigure(name, opts...)
}

func pickFontFamily() string {
	preferred := []string{
		"Noto Sans",
		"Inter",
		"Segoe UI",
		"Helvetica",
		"DejaVu Sans",
	}

	available := safeFontFamilies()
	if len(available) == 0 {
		return ""
	}

	lookup := make(map[string]struct{}, len(available))
	for _, name := range available {
		lookup[strings.ToLower(name)] = struct{}{}
	}
	for _, name := range preferred {
		if _, ok := lookup[strings.ToLower(name)]; ok {
			return name
		}
	}
	return ""
}

func safeActivateTheme(name string) (ok bool) {
	defer func() {
		if recover() != nil {
			ok = false
		}
	}()
	if err := ActivateTheme(name); err != nil {
		return false
	}
	return true
}

func safeApplyFonts() {
	defer func() {
		_ = recover()
	}()
	applyFonts()
}

func safeApplyStyles(p uiPalette) {
	defer func() {
		_ = recover()
	}()
	applyStyles(p)
}

func safeConfigureBackground(color string) {
	defer func() {
		_ = recover()
	}()
	App.Configure(Background(color))
}

func safeFontFamilies() (families []string) {
	defer func() {
		if recover() != nil {
			families = nil
		}
	}()
	return FontFamilies()
}

func applyStyles(p uiPalette) {
	StyleConfigure("TFrame", Background(p.surface))
	StyleConfigure("TLabel", Background(p.surface), Foreground(p.text))
	StyleConfigure("TSeparator", Background(p.border))

	StyleConfigure("TButton", Padding("8 4"), Focussolid(false))
	StyleMap("TButton", Foreground, "disabled", p.muted)

	StyleConfigure("Primary.TButton",
		Background(p.accent),
		Foreground("#ffffff"),
		Borderwidth(0),
		Padding("12 4"),
		Focuscolor(p.accent),
		Focussolid(false),
	)
	if buttonFont != nil {
		StyleConfigure("Primary.TButton", Font(buttonFont))
	}
	StyleMap("Primary.TButton",
		Background, "active", p.accentHover, "pressed", p.accentActive,
		Foreground, "disabled", p.muted,
	)

	StyleConfigure("Accent.TButton",
		Padding("12 4"),
		Foreground("#ffffff"),
		Focussolid(false),
	)
	if buttonFont != nil {
		StyleConfigure("Accent.TButton", Font(buttonFont))
	}
	StyleMap("Accent.TButton",
		Foreground, "active", "#ffffff", "disabled", p.muted,
	)

	StyleConfigure("Secondary.TButton",
		Background(p.surface),
		Foreground(p.secondaryText),
		Borderwidth(1),
		Bordercolor(p.secondaryBorder),
		Padding("10 4"),
		Focussolid(false),
	)
	StyleMap("Secondary.TButton",
		Background, "active", p.buttonHover, "pressed", p.buttonHover,
		Foreground, "disabled", p.muted,
	)

	StyleConfigure("Tertiary.TButton",
		Background(p.surface),
		Foreground(p.tertiaryText),
		Borderwidth(0),
		Padding("6 3"),
		Focussolid(false),
	)
	StyleMap("Tertiary.TButton",
		Foreground, "active", p.secondaryText, "disabled", p.muted,
	)

	StyleConfigure("Toolbutton",
		Padding("6 3"),
		Foreground(p.tertiaryText),
		Focussolid(false),
	)
	StyleMap("Toolbutton",
		Foreground, "active", p.secondaryText, "disabled", p.muted,
	)

	StyleConfigure("TEntry",
		Fieldbackground(p.surface),
		Foreground(p.text),
		Bordercolor(p.border),
		Padding("6 4"),
	)
	StyleMap("TEntry", Bordercolor, "focus", p.accent)

	StyleConfigure("TCombobox",
		Fieldbackground(p.surface),
		Foreground(p.text),
		Bordercolor(p.border),
		Padding("6 4"),
	)
	StyleMap("TCombobox", Bordercolor, "focus", p.accent)

	StyleConfigure("TCheckbutton", Background(p.surface), Foreground(p.text), Padding("4 2"))
	StyleConfigure("TRadiobutton", Background(p.surface), Foreground(p.text), Padding("4 2"))

	StyleConfigure("TProgressbar", Background(p.accent), Troughcolor(p.border), Bordercolor(p.border))
	StyleConfigure("TScrollbar", Arrowcolor(p.muted), Troughcolor(p.border))

	StyleConfigure("Main.TFrame", Background(p.appBg))
	StyleConfigure("Content.TFrame", Background(p.surface))
	StyleConfigure("Nav.TFrame", Background(p.surface))
	StyleConfigure("Sidebar.TFrame", Background(p.sidebar))

	StyleConfigure("SidebarTitle.TLabel", Background(p.sidebar), Foreground(p.text))
	StyleConfigure("SidebarStep.TLabel", Background(p.sidebar), Foreground(p.muted))
	StyleConfigure("SidebarActive.TLabel", Background(p.sidebar), Foreground(p.text))
	StyleConfigure("SidebarDone.TLabel", Background(p.sidebar), Foreground(p.accent))
	StyleConfigure("SidebarDisabled.TLabel", Background(p.sidebar), Foreground(p.border))
}

func navButtonStyles() (primary, secondary, tertiary string) {
	if useAzureTheme {
		return "Accent.TButton", "TButton", "Toolbutton"
	}
	return "Primary.TButton", "Secondary.TButton", "Tertiary.TButton"
}

func normalizeHex(value string) (string, bool) {
	r, g, b, ok := parseHexColor(value)
	if !ok {
		return "", false
	}
	return fmt.Sprintf("#%02x%02x%02x", r, g, b), true
}

func adjustColor(value string, delta int) string {
	r, g, b, ok := parseHexColor(value)
	if !ok {
		return value
	}
	r = clampColor(r + delta)
	g = clampColor(g + delta)
	b = clampColor(b + delta)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func clampColor(value int) int {
	if value < 0 {
		return 0
	}
	if value > 255 {
		return 255
	}
	return value
}

func parseHexColor(value string) (int, int, int, bool) {
	s := strings.TrimSpace(value)
	if s == "" {
		return 0, 0, 0, false
	}
	if strings.HasPrefix(s, "#") {
		s = strings.TrimPrefix(s, "#")
	}
	if len(s) == 3 {
		s = fmt.Sprintf("%c%c%c%c%c%c", s[0], s[0], s[1], s[1], s[2], s[2])
	}
	if len(s) != 6 {
		return 0, 0, 0, false
	}

	r, err := strconv.ParseInt(s[0:2], 16, 0)
	if err != nil {
		return 0, 0, 0, false
	}
	g, err := strconv.ParseInt(s[2:4], 16, 0)
	if err != nil {
		return 0, 0, 0, false
	}
	b, err := strconv.ParseInt(s[4:6], 16, 0)
	if err != nil {
		return 0, 0, 0, false
	}
	return int(r), int(g), int(b), true
}
