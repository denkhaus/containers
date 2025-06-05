package containers

type Tuple[T1 any, T2 any] struct {
	T1Val T1
	T2Val T2
}

// Set sets the value of the tuple.
//
// This is a convenience function for updating the entire tuple at once.
func (p *Tuple[T1, T2]) Set(v1 T1, v2 T2) {
	p.T1Val = v1
	p.T2Val = v2
}

// Get1 returns the value of the first element of the tuple.

func (p *Tuple[T1, T2]) Get1() T1 {
	return p.T1Val
}

// Get2 returns the value of the second element of the tuple.

func (p *Tuple[T1, T2]) Get2() T2 {
	return p.T2Val
}

// NewTupple creates a new instance of Tuple with the given values.
//
// Example:
//
//	t := NewTupple(1, "hello")
//	t.Get1() // returns 1
//	t.Get2() // returns "hello"
func NewTupple[T1 any, T2 any](v1 T1, v2 T2) *Tuple[T1, T2] {
	return &Tuple[T1, T2]{T1Val: v1, T2Val: v2}
}
