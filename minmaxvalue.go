package containers

import (
	"bytes"
	"encoding/gob"
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

func (p *MinMaxValue[T]) GobEncode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buffer)

	if err := enc.Encode(p.minValue); err != nil {
		return nil, err
	}
	if err := enc.Encode(p.curValue); err != nil {
		return nil, err
	}
	if err := enc.Encode(p.maxValue); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p *MinMaxValue[T]) GobDecode(v []byte) error {
	dec := gob.NewDecoder(bytes.NewBuffer(v))

	if err := dec.Decode(&p.minValue); err != nil {
		return err
	}
	if err := dec.Decode(&p.curValue); err != nil {
		return err
	}
	if err := dec.Decode(&p.maxValue); err != nil {
		return err
	}

	return nil
}
