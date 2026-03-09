package main

import "testing"

func TestBuildScriptCandidatesForRollingDistro(t *testing.T) {
	candidates := buildScriptCandidates("garuda", "")
	if len(candidates) != 1 || candidates[0] != "garuda_rolling.sh" {
		t.Fatalf("unexpected rolling candidates: %#v", candidates)
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
