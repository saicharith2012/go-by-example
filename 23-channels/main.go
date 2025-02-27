package main

import "fmt"

// Channels are a powerful and fundamental feature in Go for concurrent programming.
// They provide a way for goroutines to communicate and synchronize with each other.
// Channels are typed conduits through which you can send and receive values of a specific type.
// They are designed to prevent race conditions and make concurrent programming safer and easier to reason about.

func main() {
	messages := make(chan string) // make function is used to create channels
	// In this case creating a channel that can trasmit values of type string
	// Channels are always typed.
	// You can only send and receive values of the specified type through a particular channel.

	// a new goroutine that sends a value onto the messages channel
	go func() {
		messages <- "ping" // send operation....channel <- value ("ping" a string in this case.)
	}()

	// Blocking Send: By default, send operations on channels are blocking.
	// This means that when a goroutine tries to send a value on a channel,
	// it will block (pause execution) until another goroutine is ready to receive a value from that same channel.
	// If no goroutine is ready to receive, the sending goroutine will wait indefinitely.
	// In this example, because the main goroutine is about to execute a receive operation, this send operation will be able to proceed.

	msg := <-messages // receive operation....variable <- channel

	// Blocking Receive: Receive operations on channels are also blocking by default.
	//  When a goroutine tries to receive a value from a channel,
	//  it will block (pause execution) until a value is available to be received on that channel.
	//  If no goroutine is sending on the channel, the receiving goroutine will wait indefinitely.
	//  In this example, because a goroutine is sending "ping" on the messages channel,
	//  this receive operation will be able to proceed as soon as the send happens (or shortly after).


	// Synchronization: Channels provide a way for goroutines to synchronize. 
	// The send operation in the goroutine and the receive operation in the main goroutine synchronize with each other. 
	// The receive operation in main will wait until a value is sent on the channel by the other goroutine. 
	// This ensures that the message is successfully passed from one goroutine to another.

	fmt.Println(msg)
}
