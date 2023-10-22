package sets

type Set[K comparable] map[K]struct{}

func (s Set[K]) Insert(v K) {
	s[v] = struct{}{}
}

func (s Set[K]) Remove(v K) {
	delete(s, v)
}

func (s Set[K]) Pop() K {
	for v := range s {
		s.Remove(v)
		return v
	}

	var zero K
	return zero
}

func (s Set[K]) Pick() K {
	for v := range s {
		return v
	}

	var zero K
	return zero
}

func (s Set[K]) Has(v K) bool {
	_, ok := s[v]
	return ok
}

func (s Set[K]) Cardinality() int {
	return len(s)
}

func (s Set[K]) Clear() {
	for v := range s {
		s.Remove(v)
	}
}

func (s Set[K]) Clone() Set[K] {
	c := make(Set[K], s.Cardinality())
	for v := range s {
		c.Insert(v)
	}
	return c
}

func (s Set[K]) Union(other Set[K], additional ...Set[K]) Set[K] {
	numElements := len(s) + len(other)
	for _, other := range additional {
		numElements += len(other)
	}

	union := make(Set[K], numElements)
	for _, other := range append(additional, s, other) {
		for v := range other {
			union.Insert(v)
		}
	}
	return union
}

func (s Set[K]) Add(other Set[K], additional ...Set[K]) {
	for _, other := range append(additional, other) {
		for v := range other {
			s.Insert(v)
		}
	}
}

func (s Set[K]) Intersection(other Set[K], additional ...Set[K]) Set[K] {
	intersection := make(Set[K], len(s))
	for v := range s {
		if other.Has(v) {
			intersection.Insert(v)
		}
	}
	for _, other := range additional {
		for v := range intersection {
			if !other.Has(v) {
				intersection.Remove(v)
			}
		}
	}
	return intersection
}

func (s Set[K]) Subtract(other Set[K], additional ...Set[K]) Set[K] {
	for _, other := range append(additional, other) {
		for v := range s {
			if !other.Has(v) {
				delete(s, v)
			}
		}
	}

	return s
}

func (s Set[K]) Equals(other Set[K]) bool {
	if len(s) != len(other) {
		return false
	}

	for v := range s {
		if !other.Has(v) {
			return false
		}
	}

	return true
}

func (s Set[K]) IsSuperset(other Set[K]) bool {
	return other.IsSubset(s)
}

func (s Set[K]) IsSubset(other Set[K]) bool {
	if len(s) > len(other) {
		return false
	}

	for v := range s {
		if !other.Has(v) {
			return false
		}
	}

	return true
}

func (s Set[K]) ToSlice() []K {
	slice := make([]K, 0, len(s))
	for v := range s {
		slice = append(slice, v)
	}
	return slice
}
