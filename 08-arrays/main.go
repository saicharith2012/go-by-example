package main

import "fmt"

// arrays -- fixed-size sequential collection of the elements of the same type.

func main() {
	var a [5]int // array declaration without explicit initialization. (by default initialized with zero values)
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	// length of the array 
	fmt.Println("length of a: ", len(a))

	// array declaration and initialization using short variable declarations
	b := [5]int{4, 5, 6, 7, 8}
	fmt.Println("dcl:", b)

	// 2d arrays
	var Twod [2][3]int
	for i := 0; i<2; i++ {
		for j := 0; j<3; j++ {
			Twod[i][j] = i + j;
		}
 	}

	fmt.Println("2d array: ", Twod)
}