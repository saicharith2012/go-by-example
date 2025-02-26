package main

import "fmt"

// Recursion is a programming technique where a function calls itself, either directly or indirectly. 
// It can be a powerful way to solve problems that can b ebroken down into smaller, self-similar subproblems.
// But if the recursion doesn't have a proper base case to stop it, it can lead to stack overflow errors.


func fact(n int) int {
	if n == 0 { // base case
		return 1
	}

	return n * fact(n - 1) // recursive step
}


func main() {
	fmt.Println(fact(5))

	// declaring the fib variable that holds the function beforehand since,
	//  it needs to be recursively called inside the function.
	var fib func(n int) int

	// anonymous recursive function assigned to a variable
	fib = func (n int) int {
		if n < 2 { // base case
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}