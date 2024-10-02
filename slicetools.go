package slicetools

// Applies a function, fn, to a slice, returning a new slice with the function applied to each element
func Map[X, Y any](xs []X, fn func (X) Y) []Y {
	yList := make([]Y, 0)

	for _, x := range xs {
		yList = append(yList, fn(x))
	}

	return yList
}

// Iterates over xs returning a slice containing only elements of xs where the predicate was true
func Filter[X any](xs []X, pred func (X) bool) []X {
	filteredXs := make([]X, 0)

	for _, x := range xs {
		if pred(x) {
			filteredXs = append(filteredXs, x)
		}
	}

	return filteredXs
}
