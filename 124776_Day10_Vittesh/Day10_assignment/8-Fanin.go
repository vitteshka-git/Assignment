
package main

import (
	"fmt"
	"sync"
)

func produce(start int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i < start+3; i++ {
		output <- i
	}
}

func main() {
	output := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go produce(1, output, &wg)
	go produce(4, output, &wg)

	go func() {
		wg.Wait()
		close(output)
	}()

	for value := range output {
		fmt.Println(value)
	}
}
