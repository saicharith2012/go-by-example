package main

import "fmt"

// Slices -- built on top of arrays, but more flexible,
// They are dynamically sized.

func main() {
	s := make([]string, 3)
	// make() is a built-in function to create slices, maps and channels
	// the capacity of the slice is set to 3 i.e. the total space allocated for the slice
	// initialized with zero values

	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("set: ", s[2])

	// slice length
	fmt.Println("len: ", len(s))

	// append element to the slice
	// append() does not modify the original slice but returns a new slice, hence the slice needs to be reassigned.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// creating a slice of the same length as s
	c := make([]string, len(s))
	// copying elements from one slice to another
	copy(c, s)
	fmt.Println("cpy: ", c)

	// creating sub slices
	// create new slices that refer to a portion of an existing slice or array.
	// It doesn't copy the underlying data;
	// it creates a new slice header that points to a segment of the original data.
	l := s[2:5] // specifies a half-open range [2,5) last element is not included.
	fmt.Println("sl1: ", l)

	l = s[:5] // from the beginning to just before 5
	fmt.Println("sl2: ", l)

	l = s[2:] // from 2 to the end.
	fmt.Println("sl3: ", l)

	// declare and initializing slices similar to array literals.
	// using short variable declaration
	t := []string{"g", "h", "i"} 
	// no size specified in the square brackets [] (unlike arrays)
	// This is what distinguishes a slice literal from an array literal
	// Go infers the length and capacity of the slice based on the number of elements in the literal.
	fmt.Println("dcl: ", t)


	// 2d slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ { // iterating on the outer slice
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) // creating inner slice and assigning it to the current element
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD: ", twoD)
}
