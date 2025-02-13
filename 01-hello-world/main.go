package main // package name is declared here.  "main" indicates that it a executable program and not a library.

import "fmt" // fmt is a go standard library, provides functions for formatted input and output.

// main function --> code exectuion starts here i.e. entry point.
func main() { 
	fmt.Println("O mara manishi....")
	fmt.Println("naa loki raa....")
	fmt.Println("tun tudo tuntun yanthruda...")
	fmt.Println("tun tudo tuntun yanthruda...")
	fmt.Println("raaa yanthruda!")
}

// go run main.go ---> runs the program
// go build main.go --->  builds the program to binaries that are executable.
// ./main ---> execute the binary
