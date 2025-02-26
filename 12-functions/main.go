package main

import "fmt"
// functions are the fundamental building blocks.
// They allow you to modularize your code.
// thus making it organized, reusable and easier to understand.
// Functions are blocks of code that perform specific tasks.
// You can call or invoke a function from different parts of your program and
// Functions also return values.


// function definition - > func func_name(parameter list) returntype {...}
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res:= plus(1,2)
	fmt.Println("1+2=", res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3=", res)

}
