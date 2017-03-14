// We can use channels to synchronize execution across goroutines.
package main

import (
	"fmt"
	"time"
)

// The done channel will be used to notify another goroutine that
// this function's work is done
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// send a value to notify that we're done
	done <- true
}

func main() {

	// start a wroker goroutine, giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// block until we receive a notification from the worker on the channel
	<-done
}
