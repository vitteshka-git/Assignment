package main
import "fmt"

func main() {

	/*
		Unbuffered Channel

		make(chan int) creates an unbuffered channel.

		Capacity = 0

		Sender  --->  Channel  --->  Receiver

		The sender waits until the receiver receives
		the value.
	*/

	ch := make(chan int)

	// Send value from a goroutine
	go func() {
		fmt.Println("Sending value...")
		ch <- 100
		fmt.Println("Value sent")
	}()

	// Receive value
	value := <-ch

	fmt.Println("Received:", value)
}

