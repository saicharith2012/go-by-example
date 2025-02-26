package main

import "fmt"

// A closure is a function value that "closes over" or captures variables from its surrounding scope.
// This means that even after the outer function has finished executing,
// the closure function can still access and modify the variables from that outer scope.
// Closures are very useful for creating functions that maintain state
// or for creating more flexible and dynamic function behavior.

// defining functions that return other functions.
// The returned function is what we call a closure.
func intSeq() func() int {
	i := 0 // this i is local to the scope of intSeq function.

	// returning an anonymous function
	return func() int {

		// This anonymous function "closes over" the variable `i` from the outer `intSeq` function's scope.
		// This means it can access and modify `i` even after `intSeq` has finished executing.
		i += 1
		return i
	}
}

func main() {
	newInt := intSeq() // the closure is returned and assigned into newInt

	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())

	// 'i' is persisting its value between calls to `nextInt()`.
	// This is because 'nextInt' is a closure, and it "remembers" the 'i' from its creation scope.

	// creating a new closure
	newInts := intSeq()
	fmt.Println(newInts())

}

// Closures are used in various scenarios in Go, such as:

// Creating stateful helper functions.
// Implementing function factories (like intSeq which generates integer sequence functions).
// In event handling and callbacks.
// In concurrent programming (goroutines).
