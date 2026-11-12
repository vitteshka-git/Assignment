package main
import (
	"fmt"
	"time"
)

func main() {

	/*
		Anonymous Goroutine

		There is no function name like task().
		We directly create the function and run it
		using the "go" keyword.
	*/

	go func() {

		fmt.Println("Anonymous goroutine is running")

	}()

	// Give the goroutine time to execute.
	time.Sleep(1 * time.Second)

	fmt.Println("Main completed")
}

