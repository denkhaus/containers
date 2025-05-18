package containers

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Diff[T Number](a, b T) T {
	return a - b
}

func Sum[T Number](a, b T) T {
	return a + b
}

func Sub[T Number](a, b T) T {
	return a - b
}

func Div[T Number](a, b T) T {
	return a / b
}

func Mult[T Number](a, b T) T {
	return a * b
}

// Max returns the larger of a and b.
// If a == b, a is returned.
//
// This function is intended to be used with types that support
// the constraints.Ordered type constraint, such as int, float64,
// and string.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	if a < b {
		return b
	}

	return a
}

// Min returns the smaller of a and b.
// If a == b, a is returned.
// This function is intended to be used with types that support
// the constraints.Ordered type constraint, such as int, float64,
// and string.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	if a > b {
		return b
	}

	return a
}

// MaxValue returns the largest value in a slice.
// Assumes the slice is non-empty.
func MaxValue[T constraints.Ordered](a []T) T {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}

// MinValue returns the smallest value in a slice.
// If the slice is empty, it returns the zero value of type T.
func MinValue[T constraints.Ordered](a []T) T {
	m := a[0]
	for _, v := range a {
		if m > v {
			m = v
		}
	}
	return m
}

// Between checks if val is between min and max, inclusive.
func BetweenInclusive[T constraints.Ordered](min, max, val T) bool {
	return min <= val && val <= max
}

// Between checks if val is between min and max, exclusive.
func BetweenExclusive[T constraints.Ordered](min, max, val T) bool {
	return min < val && val < max
}
