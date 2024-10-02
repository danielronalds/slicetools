package slicetools

import (
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	original := []int{1,2,3,4}

	fn := func(x int) int {
		return x + 1
	}

	mapped := Map(original, fn)

	if slices.Compare(mapped, []int{2,3,4,5}) != 0 {
		t.Fatalf("Function was not applied as expected")
	}
}

func TestMapOneElement(t *testing.T) {
	original := []int{1}

	fn := func(x int) int {
		return x + 1
	}

	mapped := Map(original, fn)

	if slices.Compare(mapped, []int{2}) != 0 {
		t.Fatalf("Function was not applied as expected")
	}
}

func TestMapEmpty(t *testing.T) {
	original := make([]int, 0)

	fn := func(x int) int {
		return x + 1
	}

	mapped := Map(original, fn)

	if slices.Compare(mapped, make([]int, 0)) != 0 {
		t.Fatalf("Function was not applied as expected")
	}
}
