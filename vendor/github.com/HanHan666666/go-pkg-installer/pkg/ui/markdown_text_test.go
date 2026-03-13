package ui

import (
	"reflect"
	"testing"
)

func TestParseMarkdownSegments(t *testing.T) {
	content := "下载请见 [发布页](https://example.com/releases/latest)\n或直接访问 https://example.com/direct 。"

	got := parseMarkdownSegments(content)
	want := []markdownSegment{
		{Text: "下载请见 "},
		{Text: "发布页", URL: "https://example.com/releases/latest"},
		{Text: "\n或直接访问 "},
		{Text: "https://example.com/direct", URL: "https://example.com/direct"},
		{Text: " 。"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected segments:\n got: %#v\nwant: %#v", got, want)
	}
}

func TestParseMarkdownSegmentsIgnoresInvalidMarkdownLink(t *testing.T) {
	content := "[无效链接](ftp://example.com) https://example.com"

	got := parseMarkdownSegments(content)
	want := []markdownSegment{
		{Text: "[无效链接](ftp://example.com) "},
		{Text: "https://example.com", URL: "https://example.com"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected segments:\n got: %#v\nwant: %#v", got, want)
	}
}
