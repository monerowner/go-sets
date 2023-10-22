package sets

func Map[K comparable, V comparable](s Set[K], f func(K) V) Set[V] {
	mapped := make(Set[V], len(s))
	for elem := range s {
		mapped.Insert(f(elem))
	}
	return mapped
}

func Filter[K comparable](s Set[K], f func(K) bool) Set[K] {
	filtered := make(Set[K], len(s))
	for elem := range s {
		if f(elem) {
			filtered.Insert(elem)
		}
	}
	return filtered
}

func FilterIP[K comparable](s Set[K], f func(K) bool) Set[K] {
	for elem := range s {
		if !f(elem) {
			s.Remove(elem)
		}
	}

	return s
}

func ReduceInit[K comparable, A any](s Set[K], f func(A, K) A, initial A) A {
	red := initial
	for elem := range s {
		red = f(red, elem)
	}
	return red
}

func Reduce[K comparable](s Set[K], f func(K, K) K) K {
	if s.Size() == 0 {
		return zero[K]()
	}

	first := true
	var red K
	for elem := range s {
		if first {
			red = elem
			first = false
		} else {
			red = f(red, elem)
		}
	}
	return red
}
