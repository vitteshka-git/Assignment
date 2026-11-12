
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(
	ctx context.Context,
	id int,
	jobs <-chan int,
	results chan<- int,
	wg *sync.WaitGroup,
	mu *sync.Mutex,
	processed *int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker", id, "stopped")
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

			// Simulate processing time.
			select {
			case <-ctx.Done():
				fmt.Println("Worker", id, "cancelled")
				return
			case <-time.After(300 * time.Millisecond):
			}

			// Protect shared data with a mutex.
			mu.Lock()
			*processed = *processed + 1
			mu.Unlock()

			// Send the result safely.
			select {
			case results <- job * 2:
				fmt.Println("Worker", id, "processed job", job)
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	// Set a timeout of 2 seconds.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup
	var mu sync.Mutex
	processed := 0

	// Start 3 workers: fan-out.
	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go worker(
			ctx,
			i,
			jobs,
			results,
			&wg,
			&mu,
			&processed,
		)
	}

	// Send jobs to workers.
	go func() {
		defer close(jobs)

		for i := 1; i <= 10; i++ {
			select {
			case jobs <- i:
			case <-ctx.Done():
				fmt.Println("Job sending cancelled")
				return
			}
		}
	}()

	// Close results after all workers finish.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Receive all results: fan-in.
	for result := range results {
		fmt.Println("Result:", result)
	}

	mu.Lock()
	fmt.Println("Total processed:", processed)
	mu.Unlock()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Timeout reached")
	}

	fmt.Println("Program completed")
}
