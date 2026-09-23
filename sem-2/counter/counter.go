//go:build !solution

package counter

func New(start, step int) func() int {
	current := start

	return func() int {
		current += step

		return current
	}
}
