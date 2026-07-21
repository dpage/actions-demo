package calc

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"two positives", 2, 3, 5},
		{"positive and negative", 10, -4, 6},
		{"two zeroes", 0, 0, 0},
		{"two negatives", -7, -8, -15},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Add(tc.a, tc.b); got != tc.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"positive result", 9, 4, 5},
		{"negative result", 4, 9, -5},
		{"subtract zero", 6, 0, 6},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Subtract(tc.a, tc.b); got != tc.want {
				t.Errorf("Subtract(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"two positives", 6, 7, 42},
		{"by zero", 12, 0, 0},
		{"by one", 12, 1, 12},
		{"negative operand", -3, 5, -15},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Multiply(tc.a, tc.b); got != tc.want {
				t.Errorf("Multiply(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(84, 2)
	if err != nil {
		t.Fatalf("Divide(84, 2) returned unexpected error: %v", err)
	}
	if got != 42 {
		t.Errorf("Divide(84, 2) = %d, want 42", got)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(1, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Errorf("Divide(1, 0) returned error %v, want %v", err, ErrDivideByZero)
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		name   string
		values []int
		want   int
	}{
		{"empty slice", nil, 0},
		{"single value", []int{5}, 5},
		{"several values", []int{1, 2, 3, 4}, 10},
		{"mixed signs", []int{10, -3, -7}, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Sum(tc.values); got != tc.want {
				t.Errorf("Sum(%v) = %d, want %d", tc.values, got, tc.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	if _, ok := Max(nil); ok {
		t.Error("Max(nil) reported ok = true, want false")
	}

	cases := []struct {
		name   string
		values []int
		want   int
	}{
		{"single value", []int{3}, 3},
		{"largest last", []int{1, 2, 9}, 9},
		{"largest first", []int{9, 2, 1}, 9},
		{"all negative", []int{-9, -2, -5}, -2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := Max(tc.values)
			if !ok {
				t.Fatalf("Max(%v) reported ok = false, want true", tc.values)
			}
			if got != tc.want {
				t.Errorf("Max(%v) = %d, want %d", tc.values, got, tc.want)
			}
		})
	}
}
