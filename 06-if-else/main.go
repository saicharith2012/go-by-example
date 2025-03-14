package main

import "fmt"

func main() {
	// conditional execution

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%2 == 0 {
		fmt.Println("8 is divisible by 2.")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even.")
	}

	// with a short variable declaration before condition --->  the scope of this variable is limited to this if/else-if/else block
	if num := 9; num < 0 {
		fmt.Println(num, "is negative.")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit.")
	} else {
		fmt.Println(num, "has multiple digits.")
	}
}
