//go:build !solution

package testcleanup

import (
	"errors"
	"io/fs"
	"os"
	"testing"
)

func setupTempFile(t *testing.T, content string) string {
	t.Helper()

	f, err := os.CreateTemp("", "seminar-*.txt")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}

	path := f.Name()

	t.Cleanup(func() {
		if err := os.Remove(path); err != nil && !errors.Is(err, fs.ErrNotExist) {
			t.Errorf("remove temp file %q: %v", path, err)
		}
	})

	if _, err := f.WriteString(content); err != nil {
		_ = f.Close()
		t.Fatalf("write temp file %q: %v", path, err)
	}

	if err := f.Close(); err != nil {
		t.Fatalf("close temp file %q: %v", path, err)
	}

	return path
}
