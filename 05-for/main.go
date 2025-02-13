package main

import "fmt"

func main() {
	// for --> is the only looping construct in Golang.


	// condition based loop (similar to while)
	i:= 1
	for i <=3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop with initialization, condition, post-statement.
	for j := 7; j<= 9; j++ {
		fmt.Println(j)
	}

	for { // infinite loop with a break.
		fmt.Println("looop")
		break // exits the inner most loop.
	}

	for n:=0; n<=5; n++ {
		if n%2 == 0 {
			continue // skips the current iteration
		}

		fmt.Println(n)
	}
}