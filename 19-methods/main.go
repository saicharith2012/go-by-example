package main

import "fmt"

// Methods are a way to associate functions with specific types in Go.
// Methods can be defined on struct types.
// methods are similar to functions, but they have a "receiver" that specifies the type that the method is associated with.
// methods allow to add behaviour to the custom data types.

type rect struct {
	width  int
	height int
}

// methods
// (r rect) -> receiver specification (receiver_parameter receiver_type)
// receiver parameters are used to access struct fields within the methods
func (r rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 12, height: 15}

	fmt.Println("area: ",r.area())
	fmt.Println("perimeter: ",r.perimeter())

	rp := &r
	
	fmt.Println("area: ",rp.area()) // Automatic dereferencing when calling methods on pointers to structs.
	fmt.Println("perimeter: ",rp.perimeter())
}
