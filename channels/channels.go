// Channels are the pipes that connect concurrent goroutines. You can
// send values into channels from one goroutine and recieve those values
// into another goroutine
package main

import "fmt"

func main() {

	// create a new channel with make(chan val-type)
	messages := make(chan string)

	// send a value into a channle using channel <- syntax
	go func() { messages <- "ping" }()

	// The <- channel syntax receives a value from the channel
	msg := <-messages
	fmt.Println(msg)

	// By default sends and receives block until both the sender and
	// receiver are ready
}
