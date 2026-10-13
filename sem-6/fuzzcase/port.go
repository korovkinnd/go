//go:build !solution

package fuzzcase

import "strconv"

// ParsePort parses a decimal port in the range 1 through 65535.
func ParsePort(s string) (uint16, error) {
	if s == "" {
		return 0, ErrPort
	}

	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, ErrPort
		}
	}

	n, err := strconv.Atoi(s)
	if err != nil || n < 1 || n > 65535 {
		return 0, ErrPort
	}

	return uint16(n), nil
}
