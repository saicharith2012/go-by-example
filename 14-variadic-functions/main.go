package main

import "fmt"

// Variadic functions are functions that can be called with a variable number of arguments.
// Useful when it is not known in advance how many arguments a function might need to process.
// for eg, fmt.Println function itself can take any number of arguments to print.

// nums ...int --> variadic parameter i.e., accepting 0 or more arguments of type int.
// Inside the function, nums acts like a slice of ints ([]int).
// When the function is called, all the variadic arguments passed to it will be collected into a slice named nums.
func sum(nums ...int) {
	fmt.Println(nums)
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println("total: ",total)
}

func main() {
	sum(1,2)
	sum(1,2,3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
	// The ... after nums is called the "slice expansion" or "unpacking" operator. 
	// It tells Go to expand the elements of the slice nums and 
	// pass them as individual arguments to the sum function's variadic parameter.
}