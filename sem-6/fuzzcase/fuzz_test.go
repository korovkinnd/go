//go:build !solution

package fuzzcase

import (
	"errors"
	"strconv"
	"testing"
)

func FuzzParsePort(f *testing.F) {
	// TODO: добавьте seeds и проверку свойства через f.Fuzz.
	for _, s := range []string{"80", "65535", "0", "", "http", "00080", "65536"} {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		got, err := ParsePort(s)
		if err != nil {
			if got != 0 {
				t.Errorf("ParsePort(%q) = %d with error, want 0", s, got)
			}

			if !errors.Is(err, ErrPort) {
				t.Errorf("ParsePort(%q) error = %v, want ErrPort", s, err)
			}

			return
		}

		if s == "" {
			t.Fatalf("ParsePort(%q) = %d, nil, want ErrPort for empty input", s, got)
		}

		for i := 0; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				t.Fatalf("ParsePort(%q) = %d, nil, want ErrPort for non-digit byte %q", s, got, s[i])
			}
		}

		if got < 1 {
			t.Fatalf("ParsePort(%q) = %d, nil, want port in 1..65535", s, got)
		}

		canonical := strconv.FormatUint(uint64(got), 10)

		again, err := ParsePort(canonical)
		if err != nil || again != got {
			t.Errorf("ParsePort(%q) = %d, %v, want %d, nil", canonical, again, err, got)
		}
	})
}

func TestParsePortRegression(t *testing.T) {
	for _, s := range []string{"+1", "+80", "-0", "+0"} {
		got, err := ParsePort(s)
		if got != 0 || !errors.Is(err, ErrPort) {
			t.Errorf("ParsePort(%q) = %d, %v, want 0, ErrPort", s, got, err)
		}
	}
}
