// Timers are for when you want to do something once in the future
// Tickers are for when you want to do something repeatedly at regular
// intervals
package main

import (
	"fmt"
	"time"
)

func main() {

	// Tickers use a similar mechanisms to timers: a channel that is sent values
	// Here we'll use the range builtin on the channel to iterate over the values
	// as they arrive every 500ms
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Tickers can be stopped like timers. Once a ticker is stopped it won't
	// receive any more values on its channel
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
