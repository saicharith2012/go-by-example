package main

import (
	"fmt"
	"time"
	// "time"
)

// Goroutines are a lightweight concurrency feature in Go.
// They are functions that execute concurrently with other goroutines within the same address space.
// Think of them as lightweight threads.
// Goroutines make it very easy to write concurrent and parallel programs in Go.
// They are a fundamental part of Go's concurrency model and are essential for building efficient and scalable applications.

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct") // regular function call...
	// that runs directly in the main goroutine...gets executed completely before moving to the next line. ie. sequentially

	// The go keyword is used to start a new goroutine. Just put go before a function call.
	go f("goroutine") // launches a new goroutine that will execute f("goroutine")

	// Concurrency: The f("goroutine") function will now run concurrently with the main function.
	// This means that the main function doesn't wait for f("goroutine") to finish before proceeding to the next line.
	// Instead, f("goroutine") starts running in the background.

	// this block launches a new goroutine too, but it executes an anonymous function unlike the previous case where a predefined function is executed.
	go func(msg string) {
		fmt.Println(msg)
	}("going") // Immediately after defining the anonymous function, we are calling it and passing the argument "going".

	time.Sleep(time.Second) // This line pauses the execution of the main goroutine for 1 second.
	fmt.Println("done")

	// Synchronization : When you have concurrent goroutines,
	//  you often need mechanisms to synchronize them,
	//  to wait for them to complete, to communicate between them,
	// or to protect shared resources from race conditions.

	// In this simple example, time.Sleep is used as a very basic form of
	// synchronization to wait for the goroutines to complete (or at least run for some time).
	// In real-world concurrent programs, you would typically use more robust synchronization mechanisms
	// like channels (which we'll learn about soon) or wait groups to coordinate goroutines properly.

	// Main Goroutine: Every Go program has at least one goroutine, the main goroutine,
	//  which is where the main function executes.
	//  When the main function returns, the program exits,
	//  and any other goroutines that are still running are abruptly terminated.
	//  This is why time.Sleep is used here – to keep the main goroutine alive long enough to see the output from the other goroutines.
}
