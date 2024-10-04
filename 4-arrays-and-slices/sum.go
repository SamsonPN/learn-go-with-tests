package main

// compiler was complaining if we use []int or [...int]
// we must be explicit in the size of the array

// THIS IS BECAUSE THE SIZE OF THE ARRAY IS ENCODED IN ITS TYPE
// IT IS LITERALLY A PART OF IT!!!

// if you try to use [4] int, it would not compile!
// func Sum(numbers [5]int) int {

// refactored
func Sum(numbers []int) int {
	sum := 0

	// range returns [index, value]
	// we ignore index here by using the blank identifier "_"
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// variadic functions can take a variable number of arguments
func SumAll(numbersToSum ...[]int) []int {

	// len = number of elements a slice holds
	// capacity = number of elements it can hold in the underlying array
	// e.g. make([]int, 0, 5) creates a slice with length 0 and capacity 5

	// sums := make([]int, len(numbersToSum))

	// since slices have a capacity, it's better to make an empty slice here
	var sums []int

	for _, numbers := range numbersToSum {

		// append returns an updated slice
		// so you have to reassign it
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(tailsToSum ...[]int) []int {
	var sums []int

	for _, numbers := range tailsToSum {
		sum := 0

		if len(numbers) != 0 {
			sum = Sum(numbers[1:])
		}
		// slices can be sliced
		// slice[1:] = returns values from [1, len(sum) - 1]
		sums = append(sums, sum)
	}

	return sums
}
