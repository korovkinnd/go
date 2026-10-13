//go:build !solution

package testhelper

import "strings"

func assertEqual(t testingT, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertContains(t testingT, s, substr string) {
	t.Helper()

	if !strings.Contains(s, substr) {
		t.Errorf("%q does not contain %q", s, substr)
	}
}
