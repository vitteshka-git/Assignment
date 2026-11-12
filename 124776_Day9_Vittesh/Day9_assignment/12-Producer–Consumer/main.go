package main
import (
	"fmt"
	"time"
)


func producer(ch chan int) {

	// Produce 5 values.
	for i := 1; i <= 5; i++ {

		fmt.Println("Producer: Sending", i)

		ch <- i

		// Small delay between producing values.
		time.Sleep(500 * time.Millisecond)
	}

	// Producer has no more data.
	close(ch)
}

func consumer(ch chan int) {

	/*
		range receives values until the channel
		is closed by the producer.
	*/
	for value := range ch {

		fmt.Println("Consumer: Received", value)

		// Consumer takes some time to process data.
		time.Sleep(1 * time.Second)
	}
}

func main() {

	// Create a buffered channel.
	ch := make(chan int, 2)

	// Start producer goroutine.
	go producer(ch)

	// Start consumer goroutine.
	go consumer(ch)

	/*
		Keep main alive so both goroutines
		get time to complete.
	*/
	time.Sleep(6 * time.Second)

	fmt.Println("Main completed")
}

