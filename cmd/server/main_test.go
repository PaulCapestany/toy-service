package main

import "testing"

func TestResolveAddr(t *testing.T) {
	t.Setenv("PORT", "")
	if got := resolveAddr(); got != ":8080" {
		t.Fatalf("expected default :8080, got %s", got)
	}

	t.Setenv("PORT", "9090")
	if got := resolveAddr(); got != ":9090" {
		t.Fatalf("expected :9090, got %s", got)
	}

	t.Setenv("PORT", ":7070")
	if got := resolveAddr(); got != ":7070" {
		t.Fatalf("expected :7070, got %s", got)
	}
}

