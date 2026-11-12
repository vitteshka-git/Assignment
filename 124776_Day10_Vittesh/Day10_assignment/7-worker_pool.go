
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "processed job", job)
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for job := 1; job <= 5; job++ {
		jobs <- job
	}

	close(jobs)
	wg.Wait()

	fmt.Println("All jobs completed")
}
