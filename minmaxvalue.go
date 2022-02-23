package containers

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MinMaxValue[T constraints.Ordered] struct {
	minValue T
	curValue T
	maxValue T
}

func NewMinMaxValue[T constraints.Ordered](v T) *MinMaxValue[T] {
	return &MinMaxValue[T]{
		minValue: v,
		curValue: v,
		maxValue: v,
	}
}

func (p *MinMaxValue[T]) Update(v T) {
	p.minValue = Min(p.minValue, v)
	p.curValue = v
	p.maxValue = Max(p.maxValue, v)
}

func (p *MinMaxValue[T]) Min() T {
	return p.minValue
}

func (p *MinMaxValue[T]) Cur() T {
	return p.curValue
}

func (p *MinMaxValue[T]) Max() T {
	return p.maxValue
}

func (p *MinMaxValue[T]) Format(format string, a ...any) string {
	return fmt.Sprintf(format, append([]any{p.Min(), p.Cur(), p.Max()}, a...)...)
}
