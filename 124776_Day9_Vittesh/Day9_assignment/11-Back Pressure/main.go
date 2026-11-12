package main
import (
	"fmt"
	"time"
)

func main() {

	/*
		BACKPRESSURE

		Producer → Channel → Consumer

		Producer is FAST.
		Consumer is SLOW.

		When the channel becomes FULL,
		the producer has to WAIT.

		This waiting is called BACKPRESSURE.
	*/

	// Buffer can hold only 2 values.
	ch := make(chan int, 2)

	// PRODUCER
	go func() {

		for i := 1; i <= 5; i++ {

			fmt.Println("Producing:", i)

			// If the channel is full,
			// this send will WAIT.
			ch <- i

			fmt.Println("Sent:", i)
		}

		close(ch)
	}()

	// CONSUMER
	for value := range ch {

		fmt.Println("Consuming:", value)

		// Consumer is slow.
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Done")
}

