
package main

import (
	"fmt"
	"sync"
)

func process(workerID int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", workerID, "received", job)
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go process(i, jobs, &wg)
	}

	for i := 1; i <= 6; i++ {
		jobs <- i
	}

	close(jobs)
	wg.Wait()
}
