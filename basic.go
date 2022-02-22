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

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	if a < b {
		return b
	}

	return a
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	if a > b {
		return b
	}

	return a
}

func MaxValue[T constraints.Ordered](a []T) T {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}

func MinValue[T constraints.Ordered](a []T) T {
	m := a[0]
	for _, v := range a {
		if m > v {
			m = v
		}
	}
	return m
}
