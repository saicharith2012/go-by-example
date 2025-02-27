package main

import (
	"errors"
	"fmt"
)

// Error handling is a crucial aspect of robust programming.
// Go has a specific way of handling errors. Unlike some languages that use exceptions,
// Go uses explicit error values. Functions in Go often return an error as the last return value.
// By convention, the error type is used to represent errors.
// The error type is itself an interface.

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42") // sentinent value to indicate failure and error value created using errors.New() with an error message inside it
	}
	return arg + 3, nil // nil indicates 'no error' or 'success'
}

// custom error type to provide more structured information about an error
type argError struct {
	arg  int    // to store the input argument that caused the error
	prob string // detailed description of the problem
}

// Error() method on the *argError pointer type
// Here, receiver is the pointer to the argError struct
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// Any type that has an Error() string method is considered to be an error.
// By implementing the Error() string method, *argError now implements the built-in error interface.

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, prob: "can't work with it"}
		// While you could technically return a value of type argError directly, 
		// it's common and often more efficient to return a pointer to a struct in Go, 
		// especially for errors. It avoids unnecessary copying of the struct value. 
		// In this example, it's not strictly necessary for functionality, but it's idiomatic Go and good practice.
		//  Because *argError implements the error interface, this &argError{...} value can be returned as an error type.
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
			// The Error() method of an error value (like our *argError value) is automatically called by Go 
			// in certain situations when Go needs to get a string representation of the error.

			// The Error() method's implementation is what determines how the information within your custom error type 
			// (like arg and prob in argError) is formatted into a user-readable error string.
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// Type assertion to access argError details
	// check if error is of the type argError
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
