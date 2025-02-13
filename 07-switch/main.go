package main

import (
	"fmt"
	"time" // package to fetch the current time, week, day.
)

func main() {
	// switch statements are to control the flow of execution
	// to express multi-way branching
	// when a value needs to compared against multiple cases.
	i := 2
	fmt.Print("write ", i, " as ") // Print() doesnt skip to the next line.
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// switch cases break automatically when a case clause matches.

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // listing multiple values in a switch case clause.
		fmt.Println("Its the weekend.")
	default:
		fmt.Println("Its a weekday.")
	}

	// Condition based cases --- executed if true.
	t := time.Now() // writing cases for conditions on a pre-declared variable, no expression
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon.")
	default:
		fmt.Println("Its after noon.")
	}

	whatAmI := func(i interface{}) { 
		// anonymous function taking an argument of type interface{} 
		// ---> empty interface that holds value os any type. 
		// For function to accept arguments of different types. 


		// type switch -- used to switch on the type of the variable.
		// type assertion syntax
		// extracts the type of value stored in i and also assigns the value to the new variable 't' with that specific type defined.
		switch t := i.(type) { 
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("integer")
		default:
			fmt.Printf("Don't know type %T\n", t) // printing formatted strings, %T is the format verb that prints the type of the value.
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
