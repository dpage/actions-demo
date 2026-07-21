// Package calc provides a handful of deliberately simple arithmetic helpers.
//
// The functions here exist to give the CI demo something small and obvious to
// break: a single character changed in any of them will fail a test.
package calc

import "errors"

// ErrDivideByZero is returned by Divide when the divisor is zero.
var ErrDivideByZero = errors.New("calc: division by zero")

// Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}

// Subtract returns a minus b.
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of a and b.
func Multiply(a, b int) int {
	return a * b
}

// Divide returns a divided by b, or ErrDivideByZero if b is zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// Sum returns the total of every value in the slice; an empty slice sums to
// zero.
func Sum(values []int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

// Max returns the largest value in the slice, and false if the slice is empty.
func Max(values []int) (int, bool) {
	if len(values) == 0 {
		return 0, false
	}
	best := values[0]
	for _, v := range values[1:] {
		if v > best {
			best = v
		}
	}
	return best, true
}
