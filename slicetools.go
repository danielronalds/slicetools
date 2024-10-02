package slicetools

// Applies a function, fn, to a slice, returning a new slice with the function applied to each element
func Map[X, Y any](xs []X, fn func (X) Y) []Y {
	yList := make([]Y, 0)

	for _, x := range xs {
		yList = append(yList, fn(x))
	}

	return yList
}
