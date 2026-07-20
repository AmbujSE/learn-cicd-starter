package main

import "testing"

func TestGetPort(t *testing.T) {
	t.Run("uses env when set", func(t *testing.T) {
		t.Setenv("PORT", "9090")

		if got := getPort(); got != "9090" {
			t.Fatalf("expected env port 9090, got %q", got)
		}
	})

	t.Run("defaults to 8080 when env is empty", func(t *testing.T) {
		t.Setenv("PORT", "")

		if got := getPort(); got != "8080" {
			t.Fatalf("expected default port 8080, got %q", got)
		}
	})
}
