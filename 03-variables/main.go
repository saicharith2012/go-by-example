package main

import "fmt"

func main() {
	// var can be used to declare one or more variables
	var a string = "hello" // initialization
	fmt.Println(a)

	var b, c int = 2, 3
	fmt.Println(b)
	fmt.Println(c)

	// Go will infer the types of the initialized variables.
	var d = true
	fmt.Println(d)

	// uninitialized values are zero-valued.
	var e int
	fmt.Println(e)

	// short variable declarations -- can be used only inside a function for the first time the variable is being declared.
	// cannot be used at package level.
	f := "orangutan"
	fmt.Println(f)

}
