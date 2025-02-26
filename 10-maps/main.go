package main

import "fmt"

// Maps - associate keys to values.
// store and retreive data based on a key.
// Keys in a map must be of a comparable type like strings, integers etc.
// and all values in map must be of same type.

func main() {
	m := make(map[string]int)

	m["k1"] = 34
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	v3 := m["k3"] // not assigned in the map, returns a zero value
	fmt.Println("v3: ", v3)

	fmt.Println("len:", len(m))

	// deleting a key-value pair from the map
	delete(m, "k2") // arguments - map and key
	fmt.Println("after deleting k2 from map m: ", m)
	// if the key passed into delete function doesn't exist, it does nothing, doesn't cause any error

	// checks if a key is present in the map
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)
	// In Go, when you access a map element, you can get two return values: 
	// the value itself, and a boolean value indicating if the key was actually present in the map. 
	// This is often called the "comma ok idiom".
	// The blank identifier _ is used to discard the first return value 


	// declaring maps using a map literal
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
