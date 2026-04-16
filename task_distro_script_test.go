package main

import (
	"path/filepath"
	"testing"
)

func TestBuildScriptCandidatesForRollingDistro(t *testing.T) {
	candidates := buildScriptCandidates("garuda", "")
	if len(candidates) != 1 || candidates[0] != "garuda_rolling.sh" {
		t.Fatalf("unexpected rolling candidates: %#v", candidates)
	}
}

func TestBuildScriptCandidatesForUOS25(t *testing.T) {
	// UOS support is keyed by the exact os-release version, so keep the
	// generated filename contract stable when adding new release scripts.
	candidates := buildScriptCandidates("uos", "25")
	if len(candidates) != 1 {
		t.Fatalf("unexpected UOS 25 candidates: %#v", candidates)
	}
	if candidates[0] != "uos_25.sh" {
		t.Fatalf("unexpected UOS 25 candidate order: %#v", candidates)
	}
}

func TestResolveScriptForUOS25(t *testing.T) {
	// The repo must keep a concrete uos_25.sh file so the detect step can mark
	// UOS 25 as supported without relying on a risky generic fallback.
	path, meta, err := resolveScript("uos_25.sh", "scripts/distros")
	if err != nil {
		t.Fatalf("expected uos_25.sh to resolve: %v", err)
	}
	if filepath.Base(path) != "uos_25.sh" {
		t.Fatalf("unexpected resolved script path: %q", path)
	}
	if meta.RepoName == "" || meta.NextSteps == "" {
		t.Fatalf("expected UOS 25 META to be populated: %#v", meta)
	}
}

func TestResolveUpstreamScriptForGaruda(t *testing.T) {
	name, upstreamID, upstreamVersion, ok := resolveUpstreamScript("garuda", "", map[string]string{})
	if !ok {
		t.Fatal("expected garuda to resolve to arch upstream")
	}
	if name != "arch_rolling.sh" || upstreamID != "arch" || upstreamVersion != "rolling" {
		t.Fatalf("unexpected garuda upstream mapping: %q %q %q", name, upstreamID, upstreamVersion)
	}
}

func TestResolveUpstreamScriptForArchLikeDistro(t *testing.T) {
	name, upstreamID, upstreamVersion, ok := resolveUpstreamScript("cachyos", "", map[string]string{
		"ID_LIKE": "arch",
	})
	if !ok {
		t.Fatal("expected arch-like distro to resolve to arch upstream")
	}
	if name != "arch_rolling.sh" || upstreamID != "arch" || upstreamVersion != "rolling" {
		t.Fatalf("unexpected arch-like mapping: %q %q %q", name, upstreamID, upstreamVersion)
	}
}

func TestResolveUpstreamScriptForEvernight(t *testing.T) {
	name, upstreamID, upstreamVersion, ok := resolveUpstreamScript("evernight", "44", map[string]string{})
	if !ok {
		t.Fatal("expected evernight to resolve to fedora upstream")
	}
	if name != "fedora_44.sh" || upstreamID != "fedora" || upstreamVersion != "44" {
		t.Fatalf("unexpected evernight upstream mapping: %q %q %q", name, upstreamID, upstreamVersion)
	}
}

func TestResolveUpstreamScriptForEvernightPointRelease(t *testing.T) {
	name, upstreamID, upstreamVersion, ok := resolveUpstreamScript("evernight", "44.1", map[string]string{})
	if !ok {
		t.Fatal("expected evernight 44.x to resolve to fedora upstream")
	}
	if name != "fedora_44.sh" || upstreamID != "fedora" || upstreamVersion != "44" {
		t.Fatalf("unexpected evernight 44.x upstream mapping: %q %q %q", name, upstreamID, upstreamVersion)
	}
}

func TestResolveUpstreamScriptForUnsupportedEvernightVersion(t *testing.T) {
	if _, _, _, ok := resolveUpstreamScript("evernight", "45", map[string]string{}); ok {
		t.Fatal("did not expect evernight 45 to resolve to a supported upstream script")
	}
}
