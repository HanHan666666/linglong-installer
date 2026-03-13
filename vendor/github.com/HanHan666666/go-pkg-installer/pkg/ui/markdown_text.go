package ui

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	. "modernc.org/tk9.0"
)

// markdownSegment is the smallest rendering unit produced by the lightweight
// parser below. Plain text is represented with URL == "", while clickable
// content carries both the visible text and the target URL.
type markdownSegment struct {
	Text string
	URL  string
}

// insertMarkdownText renders a very small markdown subset into a Tk text
// widget. The goal here is intentionally narrow: support clickable links for
// installer copy without pulling in a full markdown renderer.
//
// Supported forms:
//   - Markdown links: [label](https://example.com)
//   - Bare URLs: https://example.com
//
// Everything else is inserted as plain text. Each link gets its own Tk tag so
// the click handler can keep the correct destination URL.
func insertMarkdownText(text *TextWidget, content string) {
	if text == nil || content == "" {
		return
	}

	linkCount := 0
	for _, segment := range parseMarkdownSegments(content) {
		if segment.URL == "" {
			text.Insert("end", segment.Text)
			continue
		}

		// Tk text widgets associate events with tags, not with arbitrary ranges.
		// We therefore create a unique tag per link so every clickable segment can
		// bind to its own destination URL.
		tag := fmt.Sprintf("link-%d", linkCount)
		linkCount++
		text.TagConfigure(tag, Foreground(currentPalette.accent), Underline(1))
		text.TagBind(tag, "<Enter>", func(e *Event) {
			if e != nil && e.W != nil {
				e.W.Configure(Cursor(Hand2))
			}
		})
		text.TagBind(tag, "<Leave>", func(e *Event) {
			if e != nil && e.W != nil {
				e.W.Configure(Cursor(""))
			}
		})
		text.Insert("end", segment.Text, tag)
		text.TagBind(tag, "<Button-1>", func(url string) func(*Event) {
			return func(e *Event) {
				if err := openExternalURL(url); e != nil {
					e.Err = err
				}
			}
		}(segment.URL))
	}
}

// parseMarkdownSegments scans the content left-to-right and emits a mixed
// sequence of plain-text segments and link segments. When both a markdown link
// and a bare URL are present, the earlier one wins so the parser stays stable
// and predictable without needing a full tokenizer.
func parseMarkdownSegments(content string) []markdownSegment {
	var segments []markdownSegment

	for len(content) > 0 {
		mdStart, mdEnd, mdText, mdURL, mdOK := findMarkdownLink(content)
		rawStart, rawEnd, rawURL, rawOK := findBareURL(content)

		switch {
		case mdOK && (!rawOK || mdStart <= rawStart):
			appendPlainSegment(&segments, content[:mdStart])
			segments = append(segments, markdownSegment{Text: mdText, URL: mdURL})
			content = content[mdEnd:]
		case rawOK:
			appendPlainSegment(&segments, content[:rawStart])
			segments = append(segments, markdownSegment{Text: rawURL, URL: rawURL})
			content = content[rawEnd:]
		default:
			appendPlainSegment(&segments, content)
			content = ""
		}
	}

	return segments
}

// appendPlainSegment coalesces adjacent plain-text chunks. This keeps the final
// segment list smaller and avoids unnecessary insert calls into the Tk widget.
func appendPlainSegment(segments *[]markdownSegment, text string) {
	if text == "" {
		return
	}

	if n := len(*segments); n > 0 && (*segments)[n-1].URL == "" {
		(*segments)[n-1].Text += text
		return
	}

	*segments = append(*segments, markdownSegment{Text: text})
}

// findMarkdownLink looks for the first valid markdown link in the current
// string slice. This parser intentionally supports only the simple
// [label](url) form because the installer copy does not need nested markdown,
// escaped delimiters, titles, or reference-style links.
func findMarkdownLink(content string) (start, end int, text, url string, ok bool) {
	openText := strings.IndexByte(content, '[')
	for openText >= 0 {
		closeText := strings.IndexByte(content[openText+1:], ']')
		if closeText < 0 {
			return 0, 0, "", "", false
		}
		closeText += openText + 1

		if closeText+1 >= len(content) || content[closeText+1] != '(' {
			next := strings.IndexByte(content[openText+1:], '[')
			if next < 0 {
				return 0, 0, "", "", false
			}
			openText += next + 1
			continue
		}

		closeURL := strings.IndexByte(content[closeText+2:], ')')
		if closeURL < 0 {
			return 0, 0, "", "", false
		}
		closeURL += closeText + 2

		linkText := content[openText+1 : closeText]
		linkURL := content[closeText+2 : closeURL]
		if linkText != "" && isExternalURL(linkURL) {
			return openText, closeURL + 1, linkText, linkURL, true
		}

		next := strings.IndexByte(content[openText+1:], '[')
		if next < 0 {
			return 0, 0, "", "", false
		}
		openText += next + 1
	}

	return 0, 0, "", "", false
}

// findBareURL finds the next plain http/https URL that is not already consumed
// as markdown. The parser trims trailing sentence punctuation so prose like
// "see https://example.com." still produces a correct clickable URL.
func findBareURL(content string) (start, end int, url string, ok bool) {
	schemes := []string{"https://", "http://"}

	bestStart := -1
	bestScheme := ""
	for _, scheme := range schemes {
		idx := strings.Index(content, scheme)
		if idx >= 0 && (bestStart < 0 || idx < bestStart) {
			bestStart = idx
			bestScheme = scheme
		}
	}
	if bestStart < 0 {
		return 0, 0, "", false
	}

	end = bestStart + len(bestScheme)
	for end < len(content) && !isURLTerminator(content[end]) {
		end++
	}

	url = strings.TrimRight(content[bestStart:end], ".,;:!?")
	if !isExternalURL(url) {
		return 0, 0, "", false
	}

	return bestStart, bestStart + len(url), url, true
}

// isURLTerminator defines characters that end a bare URL when scanning prose.
// This is intentionally conservative: spaces, line breaks, and a few common
// delimiters terminate the URL, while characters such as "/" and "?" remain
// part of it.
func isURLTerminator(ch byte) bool {
	switch ch {
	case ' ', '\n', '\r', '\t', '<', '>', '"', '\'':
		return true
	default:
		return false
	}
}

// isExternalURL limits link handling to explicit http/https destinations. That
// keeps the behavior simple and avoids turning arbitrary markdown-like text into
// clickable shell-open targets.
func isExternalURL(raw string) bool {
	return strings.HasPrefix(raw, "https://") || strings.HasPrefix(raw, "http://")
}

// openExternalURL dispatches the URL to the platform's default browser. The
// installer is Linux-first, but the helper keeps the common macOS/Windows
// commands so the vendored UI code remains portable.
func openExternalURL(raw string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", raw)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", raw)
	default:
		cmd = exec.Command("xdg-open", raw)
	}

	return cmd.Start()
}
