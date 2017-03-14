// The primary mechanism for managing state in Go is communication over channels
// We saw this for example with worker pools. There are a few other options for
// managing state though. Here we'll look at using the sync/atomic package for
// atomic counters accessed by multiple goroutines
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	// unsigned integer to represent our counter
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				// automic increment by giving it the
				// memory address
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// wait a second to allow some ops to accumulate
	time.Sleep(time.Second)

	// In order to safely use the counter while it's still being updated
	// by other goroutines, we extract a copy of the current value into
	// opsFinal via LoadUint64
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
