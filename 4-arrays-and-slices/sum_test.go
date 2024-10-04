package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		// you can declare arrays like this to have the compiler count it for you
		// numbers := [...]int{1, 2, 3, 4, 5}
		// numbers := [5]int{1, 2, 3, 4, 5}

		// to make test pass, we turn this into a slice
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	// don't need this test b/c we have full test coverage

	// use go test -cover to find out how much test coverage we have
	// t.Run("collection of any size", func(t *testing.T) {

	// 	// slices are collections of any size
	// 	// basically the same syntax but without any size or ellipses
	// 	numbers := []int{1, 2, 3}

	// 	got := Sum(numbers)
	// 	want := 6

	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// cannot compare a slice to nil
	// if got != want {

	// useful for seeing if any 2 variables are the same
	// reflect.DeepEqual IS NOT TYPE SAFE!!!!
	// if we changed bob to want := "Bob", the test would still compile
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
	t.Run("safely sum slice with one element", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
