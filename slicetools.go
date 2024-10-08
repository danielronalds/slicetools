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

// Applies a function, fn, to a slice, returning a new slice with the function applied to each element.
// Differs from `Map` as the function includes the index
func MapWithIndex[X, Y any](xs []X, fn func(int, X) Y) []Y {
	yList := make([]Y, 0)

	for i, x := range xs {
		yList = append(yList, fn(i, x))
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

// Same as Filter, except the predicate function is also passed the index of the element being filterd
func FilterWithIndex[X any](xs []X, pred func(int, X) bool) []X {
	filteredXs := make([]X, 0)

	for i, x := range xs {
		if pred(i, x) {
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

// Same as Foldl, except the index of the current element is passed to the function.
func FoldlWithIndex[X, Y any](xs []X, init Y, fn func(int, Y, X) Y) Y {
	currXs := xs
	currVal := init

	index := 0

	for true {
		if len(currXs) == 0 {
			break
		}
		x := Head(currXs)
		currVal = fn(index, currVal, x)
		currXs = Tail(currXs)
		index += 1
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

// Same as Foldr, except the index of the current element is passed to the function.
func FoldrWithIndex[X, Y any](xs []X, init Y, fn func(int, Y, X) Y) Y {
	currXs := xs

	slices.Reverse(currXs)

	return FoldlWithIndex(currXs, init, fn)
}
