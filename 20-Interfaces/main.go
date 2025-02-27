package main

import (
	"fmt"
	"math"
)

// Interfaces are one of the most powerful and fundamental concepts in Go.
// They enable polymorphism and abstraction, making Go code flexible and extensible.
// In Go, an interface is a type that defines a set of method signatures.
// A type implements an interface if it provides concrete implementations for all the methods defined in that interface.
// Importantly, this implementation is implicit - you don't need to explicitly declare that a type implements an interface;
// if it has the right methods, it just does. This is known as "duck typing" or "structural typing".

type geometry interface {
	// method signatures that any type implementing this interface must have
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// Because rect now has both area() and perim() methods that match the signatures
// in the geometry interface, rect implicitly implements the geometry interface.
// Same with circle

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// this function is designed to work with any type that implements the geometry interface
func measure(g geometry) {
	fmt.Println(g)
	// The specific area() method that gets called will depend on
	// the actual type of the value passed to measure (whether it's a rect or a circle or some other type that implements geometry).
	// This is abstraction and polymorphism in action.
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimeter: ", g.perimeter())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
