package main

import "fmt"

// structs --> structures --> composite datatypes.
// used to group together related data fields that have different types.
// i.e. to create custom data types and organize data in a meaningful way.

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20}) // creating a person instance using struct literal in a positional syntax -- order matters

	fmt.Println(person{name: "Alice", age: 30}) // named syntax - order doesnt matter

	fmt.Println(person{name: "Fred"}) // other fields that arent passed are initialised with zero values.

	fmt.Println(&person{name: "Ann", age: 40}) // pointer to the person struct instance

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// creating a pointer to s
	sp := &s

	// Even though sp is a pointer, Go allows you to use the dot operator .
	// directly to access fields of the struct that sp points to. 
	// Go automatically dereferences the pointer for you in this case. 
	// This is a convenience feature in Go. 
	// (Technically, it's shorthand for (*sp).age, but Go lets you write sp.age directly).
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}


// Structs are a cornerstone of data modeling in Go.
// They allow you to create custom types that represent real-world entities or data structures in your programs.