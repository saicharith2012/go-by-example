package main

import "fmt"

// range - is a keyword used to iterate over data structures like arrays, slices, strings, maps, channels etc.

func main() {
	nums := []int{1,2,3,4}
	sum := 0

	// for...range loop
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3{
			fmt.Println("index: ", i)
		}
	}

	// iterating over a map
	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// iterating over a string
	// When you range over a string, it iterates over the Unicode code points (runes) in the string. 
	// It provides:
    // The starting byte index of the rune in the string.
    // The rune itself (which represents a Unicode code point).

	for i,c := range "go" {
		fmt.Println(i,c)
	}
}