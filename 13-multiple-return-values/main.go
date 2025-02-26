package main

import "fmt"

// Go has built-in support to return multiple values. 
// This is useful to returning error information along with the result.
// Or could also be used to return multiple related pieces of data from a function.

func vals() (int, int) {
	return 3, 7 // in corresponding order
}

func main() {
	a,b := vals() // multiple assignment.
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // using the blank identifier to discard the first return value and assigning only the second one.
	fmt.Println(c)
}