package containers

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"golang.org/x/exp/constraints"
)

type MinMaxValue[T constraints.Ordered] struct {
	def      T
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

func (p *MinMaxValue[T]) Reset() {
	p.minValue = p.def
	p.curValue = p.def
	p.maxValue = p.def
}

func (p *MinMaxValue[T]) isUntouched() bool {
	return p.minValue == p.def && p.curValue == p.def && p.maxValue == p.def
}

func (p *MinMaxValue[T]) Update(v T) {
	if p.isUntouched() {
		p.minValue = v
		p.curValue = v
		p.maxValue = v
	}

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

func (p *MinMaxValue[T]) WriteTo(w io.Writer) error {
	enc := gob.NewEncoder(w)
	if err := enc.Encode(p.minValue); err != nil {
		return err
	}
	if err := enc.Encode(p.curValue); err != nil {
		return err
	}

	return enc.Encode(p.maxValue)
}

func (p *MinMaxValue[T]) GobEncode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := p.WriteTo(buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p *MinMaxValue[T]) ReadFrom(r io.Reader) error {
	dec := gob.NewDecoder(r)
	if err := dec.Decode(&p.minValue); err != nil {
		return err
	}
	if err := dec.Decode(&p.curValue); err != nil {
		return err
	}

	return dec.Decode(&p.maxValue)
}

func (p *MinMaxValue[T]) GobDecode(v []byte) error {
	return p.ReadFrom(bytes.NewBuffer(v))
}

func (p *MinMaxValue[T]) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"min": p.minValue,
		"max": p.maxValue,
		"cur": p.curValue,
	}
}
