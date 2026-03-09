package main

import "testing"

func TestResolveUpstreamScriptForGaruda(t *testing.T) {
	name, upstreamID, upstreamVersion, ok := resolveUpstreamScript("garuda", "", map[string]string{})
	if !ok {
		t.Fatal("expected garuda to resolve to arch upstream")
	}
	if name != "arch_rolling.sh" || upstreamID != "arch" || upstreamVersion != "rolling" {
		t.Fatalf("unexpected garuda upstream mapping: %q %q %q", name, upstreamID, upstreamVersion)
	}
}
