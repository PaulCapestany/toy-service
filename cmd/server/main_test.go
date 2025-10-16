package main

import "testing"

func TestResolveAddr(t *testing.T) {
	t.Run("defaultPort", func(t *testing.T) {
		t.Setenv("PORT", "")
		if got := resolveAddr(); got != ":8080" {
			t.Fatalf("expected default :8080, got %s", got)
		}
	})

	t.Run("numericPort", func(t *testing.T) {
		t.Setenv("PORT", "9090")
		if got := resolveAddr(); got != ":9090" {
			t.Fatalf("expected :9090, got %s", got)
		}
	})

	t.Run("prefixedPort", func(t *testing.T) {
		t.Setenv("PORT", ":7070")
		if got := resolveAddr(); got != ":7070" {
			t.Fatalf("expected :7070, got %s", got)
		}
	})

	t.Run("trimmedWhitespace", func(t *testing.T) {
		t.Setenv("PORT", "  9800  ")
		if got := resolveAddr(); got != ":9800" {
			t.Fatalf("expected trimmed :9800, got %s", got)
		}
	})

	t.Run("invalidPortFallsBack", func(t *testing.T) {
		t.Setenv("PORT", "abc")
		if got := resolveAddr(); got != ":8080" {
			t.Fatalf("expected fallback :8080, got %s", got)
		}
	})

	t.Run("outOfRangePortFallsBack", func(t *testing.T) {
		t.Setenv("PORT", "70000")
		if got := resolveAddr(); got != ":8080" {
			t.Fatalf("expected fallback :8080 for out-of-range port, got %s", got)
		}
	})
}
