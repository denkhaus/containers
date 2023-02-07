package containers

type Tuple[T1 any, T2 any] struct {
	T1Val T1
	T2Val T2
}

func (p *Tuple[T1, T2]) Set(v1 T1, v2 T2) {
	p.T1Val = v1
	p.T2Val = v2
}

func (p *Tuple[T1, T2]) Get1() T1 {
	return p.T1Val
}

func (p *Tuple[T1, T2]) Get2() T2 {
	return p.T2Val
}
