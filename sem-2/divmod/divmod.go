//go:build !solution

package divmod

import "errors"

var ErrDivisionByZero = errors.New("division by zero")

func DivMod(dividend, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {
		return 0, 0, ErrDivisionByZero
	}

	return dividend / divisor, dividend % divisor, nil
}
