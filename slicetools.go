package slicetools

import "slices"

// Applies a function, fn, to a slice, returning a new slice with the function applied to each element
func Map[X, Y any](xs []X, fn func(X) Y) []Y {
	yList := make([]Y, 0)

	for _, x := range xs {
		yList = append(yList, fn(x))
	}

	return yList
}

// Iterates over xs returning a slice containing only elements of xs where the predicate was true
func Filter[X any](xs []X, pred func(X) bool) []X {
	filteredXs := make([]X, 0)

	for _, x := range xs {
		if pred(x) {
			filteredXs = append(filteredXs, x)
		}
	}

	return filteredXs
}

// Returns a slice without it's first element
func Tail[X any](xs []X) []X {
	if len(xs) == 0 {
		return xs
	}
	return xs[1:]
}

// Returns the first element of the array
func Head[X any](xs []X) X {
	if len(xs) == 0 {
		var def X
		return def
	}
	return xs[0]
}

// Applies a function, fn, to a slice, reducing it to a value of type Y. Applies from the first
// element first to the last (left to right)
func Foldl[X, Y any](xs []X, init Y, fn func(Y, X) Y) Y {
	currXs := xs
	currVal := init

	for true {
		if len(currXs) == 0 {
			break
		}
		x := Head(currXs)
		currVal = fn(currVal, x)
		currXs = Tail(currXs)
	}

	return currVal
}

// Applies a function, fn, to a slice, reducing it to a value of type Y. Applies from the last
// element to the first (right to left)
func Foldr[X, Y any](xs []X, init Y, fn func(Y, X) Y) Y {
	currXs := xs
	// FIXME: There should be a more efficient way to do this
	slices.Reverse(currXs)

	return Foldl(currXs, init, fn)
}
