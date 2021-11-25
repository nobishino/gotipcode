package calc

type Number interface {
    int | int32 | int64 | float32 | float64
}

func Max[T Number] (x, y T) T {
	if x >= y {
		return x
	}
	return y
}