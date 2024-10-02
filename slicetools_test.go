package slicetools

import (
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	original := []int{1, 2, 3, 4}

	fn := func(x int) int {
		return x + 1
	}

	mapped := Map(original, fn)

	if slices.Compare(mapped, []int{2, 3, 4, 5}) != 0 {
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

func TestFilter(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}

	isEven := func(x int) bool {
		return (x % 2) == 0
	}

	expected := []int{2, 4}

	if slices.Compare(Filter(xs, isEven), expected) != 0 {
		t.Fatalf("Function was not applied as expected")
	}
}

func TestFilterOneElement(t *testing.T) {
	xs := []int{1}

	isEven := func(x int) bool {
		return (x % 2) == 0
	}

	expected := make([]int, 0)

	if slices.Compare(Filter(xs, isEven), expected) != 0 {
		t.Fatalf("Function was not applied as expected")
	}
}

func TestFilterEmpty(t *testing.T) {
	xs := make([]int, 0)

	isEven := func(x int) bool {
		return (x % 2) == 0
	}

	expected := make([]int, 0)

	if slices.Compare(Filter(xs, isEven), expected) != 0 {
		t.Fatalf("Function was not applied as expected")
	}
}
