package set

type Set[T comparable] map[T]struct{}

func New[T comparable](xs ...T) Set[T] {
	s := make(Set[T])
	for _, xs := range xs {
		s.Add(xs)
	}
	return s
}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Includes(x T) bool {
	_,ok := s[x]
	return ok
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}