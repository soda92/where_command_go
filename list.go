package main

type List[E any] []E

func (s List[E]) Map(f func(E)) {
	for _, v := range s {
		f(v)
	}
}
