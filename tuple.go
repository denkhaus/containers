package containers

type Tuple[T1 any, T2 any] struct {
	v1 T1
	v2 T2
}

func (p *Tuple[T1, T2]) Set(v1 T1, v2 T2) {
	p.v1 = v1
	p.v2 = v2
}

func (p *Tuple[T1, T2]) Get1() T1 {
	return p.v1
}

func (p *Tuple[T1, T2]) Get2() T2 {
	return p.v2
}
