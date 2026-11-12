package main

import (
	"fmt"             
	"sync"             
	"sync/atomic"       
)

func main() {
	var counter int64 // Creates a shared integer.
	var wg sync.WaitGroup // Creates a WaitGroup.
	wg.Add(100) // Registers 100 goroutines.

	for i := 0; i < 100; i++ { // Starts 100 goroutines.
		go func() { // Creates a goroutine.
			defer wg.Done() // Signals completion.
			atomic.AddInt64(&counter, 1) // Atomically increments the counter.
		}()
	}

	wg.Wait() // Waits for all increments.
	fmt.Println("Counter:", atomic.LoadInt64(&counter)) // Atomically reads the value.
}
