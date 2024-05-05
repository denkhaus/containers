package containers

import (
	"golang.org/x/exp/slices"
)

type Slice[T any] struct {
	d []T
}

func NewSlice[T any](v ...T) *Slice[T] {
	return &Slice[T]{d: v}
}

func (p *Slice[T]) Reset() {
	p.d = []T{}
}

func (p *Slice[T]) Append(v T) *Slice[T] {
	p.d = append(p.d, v)
	return p
}

func (p *Slice[T]) Prepend(v T) *Slice[T] {
	p.InsertAt(0, v)
	return p
}

func (p *Slice[T]) InsertAt(idx int, v T) *Slice[T] {
	p.d, p.d[0] = append(p.d[:idx+1], p.d[idx:]...), v
	return p
}

func (p *Slice[T]) Remove(idx int) *Slice[T] {
	p.d = p.d[:idx+copy(p.d[idx:], p.d[idx+1:])]
	return p
}

func (p *Slice[T]) Clone() *Slice[T] {
	pp := &Slice[T]{d: make([]T, len(p.d))}
	copy(pp.d, p.d)
	return pp
}

func (p *Slice[T]) Sort(fn func(x, y T) int) *Slice[T] {
	slices.SortStableFunc(p.d, fn)
	return p
}

func (p *Slice[T]) Exists(fn func(val T) bool) bool {
	for _, item := range p.d {
		if fn(item) {
			return true
		}
	}

	return false
}

func (p *Slice[T]) First() T {
	if len(p.d) == 0 {
		var t T
		return t
	}

	return p.d[0]
}

func (p *Slice[T]) Last() T {
	nCount := len(p.d)
	if nCount == 0 {
		var t T
		return t
	}

	return p.d[nCount-1]
}

func (p *Slice[T]) GetAt(idx int) T {
	nCount := len(p.d)
	if nCount == 0 || idx >= nCount {
		var t T
		return t
	}

	return p.d[idx]
}

func (p *Slice[T]) Values() []T {
	return p.d
}

func (p *Slice[T]) Len() int {
	return len(p.d)
}

func (p *Slice[T]) Enumerate(fn func(item T) error) error {
	for _, item := range p.d {
		if err := fn(item); err != nil {
			return err
		}
	}

	return nil
}

// Take takes the first n items of []Elem
func (p *Slice[T]) Take(nCount int) *Slice[T] {
	nCount = Min(nCount, len(p.d))
	return &Slice[T]{d: p.d[:nCount]}
}

// Map turns a []Elem1 to a []Elem2 using a mapping function.
func (p *Slice[T]) Map(fn func(T) T) *Slice[T] {
	data := make([]T, len(p.d))
	for i, v := range p.d {
		data[i] = fn(v)
	}
	return &Slice[T]{d: data}
}

// Reduce reduces a []T to a single value of type T using a reduction function.
func (p *Slice[T]) Reduce(initializer T, f func(T, T) T) T {
	r := initializer
	for _, v := range p.d {
		r = f(r, v)
	}
	return r
}

func (p *Slice[T]) Select(fn func(T) bool) *Slice[T] {
	var data []T
	for _, item := range p.d {
		if fn(item) {
			data = append(data, item)
		}
	}

	return &Slice[T]{d: data}
}
