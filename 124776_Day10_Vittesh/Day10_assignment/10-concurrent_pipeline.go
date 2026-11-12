
package main

import "fmt"

func generate() <-chan int {
	out := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func multiply(input <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for value := range input {
			out <- value * 2
		}
		close(out)
	}()

	return out
}

func main() {
	numbers := generate()
	results := multiply(numbers)

	for value := range results {
		fmt.Println(value)
	}
}
