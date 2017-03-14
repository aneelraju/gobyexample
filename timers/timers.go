// We often want to execute Go code at some point in the future or
// repeatedly at some interval. Go's built-in timer and
// ticker features make both of these tasks easy
package main

import (
	"fmt"
	"time"
)

func main() {
	// Timers represent a single event in the future. You tell the
	// timer how long you want to wait, and it provides a channel
	// that will be notified at that time. This timer will wait 2s
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	// If you just want to wait, you could have used time.Sleep
	// One reason a timer may be useful is that you can cancel the
	// timer before it expires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
