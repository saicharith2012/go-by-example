package main

import "fmt"

// pointers are used to work directly with memory addresses.

// pass by value -- receives only a copy of i -- any changes to it will not affect the original
func zeroval(ival int) {
	ival = 0
}

// pass by pointer -- address of i is being passed as an argument into this pointer parameter.
func zeroptr(iptr *int) {
	*iptr = 0 // dereference operator to set the value at the memory address stored in iptr to 0.
	// changes are translated at the original location
}

func main() {
	i := 1
	fmt.Println("initial: ", i)


	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}