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

func TestMapWithIndex(t *testing.T) {
	original := []int{1, 2, 3, 4}

	fn := func(i int, x int) int {
		return i
	}

	mapped := MapWithIndex(original, fn)

	if slices.Compare(mapped, []int{0, 1, 2, 3}) != 0 {
		t.Fatalf("Function was not applied as expected: actual slice %v", mapped)
	}
}

func TestMapWithIndexOneElement(t *testing.T) {
	original := []int{1}

	fn := func(i int, x int) int {
		return i
	}

	mapped := MapWithIndex(original, fn)

	if slices.Compare(mapped, []int{0}) != 0 {
		t.Fatalf("Function was not applied as expected: actual slice %v", mapped)
	}
}

func TestMapWithIndexEmpty(t *testing.T) {
	original := make([]int, 0)

	fn := func(i int, x int) int {
		return i
	}

	mapped := MapWithIndex(original, fn)

	if slices.Compare(mapped, make([]int, 0)) != 0 {
		t.Fatalf("Function was not applied as expected: actual slice %v", mapped)
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

func TestTail(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}

	expected := []int{2, 3, 4, 5}

	if slices.Compare(Tail(xs), expected) != 0 {
		t.Fatalf("Function was not applied as expected")
	}
}

func TestTailOneElement(t *testing.T) {
	xs := []int{1}

	expected := make([]int, 0)

	if slices.Compare(Tail(xs), expected) != 0 {
		t.Fatalf("Function was not applied as expected")
	}
}

func TestTailEmpty(t *testing.T) {
	xs := make([]int, 0)

	expected := make([]int, 0)

	if slices.Compare(Tail(xs), expected) != 0 {
		t.Fatalf("Function was not applied as expected")
	}
}

func TestHead(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}

	expected := 1

	if Head(xs) != expected {
		t.Fatalf("Function was not applied as expected")
	}
}

func TestHeadEmpty(t *testing.T) {
	xs := make([]int, 0)

	expected := 0 // The default value for an int

	if Head(xs) != expected {
		t.Fatalf("Function was not applied as expected")
	}
}

func TestFoldl(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}

	fn := func(total, val int) int {
		return total + val
	}

	expected := 15

	actual := Foldl(xs, 0, fn)

	if actual != expected {
		t.Fatalf("Function was not applied as expected, wanted %v got %v", expected, actual)
	}
}

func TestFoldlComplex(t *testing.T) {
	xs := []int{-20, 5, 2}

	fn := func(total, val int) int {
		return total - val
	}

	expected := 13

	actual := Foldl(xs, 0, fn)

	if actual != expected {
		t.Fatalf("Function was not applied as expected, wanted %v got %v", expected, actual)
	}
}

func TestFoldlOneElement(t *testing.T) {
	xs := []int{2}

	fn := func(total, val int) int {
		return total + (val * 2)
	}

	expected := 4

	actual := Foldl(xs, 0, fn)

	if actual != expected {
		t.Fatalf("Function was not applied as expected, wanted %v got %v", expected, actual)
	}
}

func TestFoldlEmpty(t *testing.T) {
	xs := make([]int, 0)

	fn := func(total, val int) int {
		return total + (val * 2)
	}

	expected := 10 // Should be the init value

	actual := Foldl(xs, 10, fn)

	if actual != expected {
		t.Fatalf("Function was not applied as expected, wanted %v got %v", expected, actual)
	}
}
