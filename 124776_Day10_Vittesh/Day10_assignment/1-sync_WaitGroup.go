package main

import (
	"fmt"             
	"sync"            
	"time"              
)

func main() {
	var wg sync.WaitGroup 

	wg.Add(3) /

	for i := 1; i <= 3; i++ { // Starts three worker goroutines.
		go func(worker int) { // Launches a goroutine and receives the worker number.
			defer wg.Done() // Marks this goroutine as completed when it returns.
			fmt.Println("Worker started:", worker) // Prints the worker number.
			time.Sleep(time.Second) // Simulates some work.
			fmt.Println("Worker finished:", worker) // Prints completion.
		}(i) // Passes the current loop value safely to the goroutine.
	}

	wg.Wait() 
	fmt.Println("All workers completed") // Prints after every worker finishes.
}
