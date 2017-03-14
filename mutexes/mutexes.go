// For more complex state we can use a mutex to safely access data
// across multiple goroutines
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var state = make(map[int]int)

	// This mutex will synchronize access to state
	var mutex = &sync.Mutex{}

	var readOps uint64 = 0
	var writeOps uint64 = 0

	// 100 goroutines to execute repeated reads against the state
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// For each read we pick a key to access, Lock() the mutex to ensure
				// exclusive access to the state, read the value at the choosen key,
				// Unlock() the mutex, and increment the readOps count
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We'll also start 10 goroutines to simulate writes, using the same pattern we di for reads
	for w := 0; w < 10; w++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)
			mutex.Lock()
			state[key] = val
			mutex.Unlock()
			atomic.AddUint64(&writeOps, 1)
			time.Sleep(time.Millisecond)
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
