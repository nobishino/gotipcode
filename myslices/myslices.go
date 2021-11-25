package myslices

// copied from github.com/golang/exp

func Equal[E comparable](s1, s2 []E) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if v1 != v2 {
			return false
		}
	}
	return true
}

func Delete[S ~[]E, E any](s S, i, j int) S {
	return append(s[:i], s[j:]...)
}