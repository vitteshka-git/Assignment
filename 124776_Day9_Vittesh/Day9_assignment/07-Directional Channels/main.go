package main

import "fmt"

/*
	DIRECTIONAL CHANNELS

	chan int
		→ Can send and receive

	chan<- int
		→ Send only

	<-chan int
		→ Receive only
*/

// This function can ONLY send to the channel.
func sendData(ch chan<- int) {

	fmt.Println("Sending data...")

	ch <- 100
}

// This function can ONLY receive from the channel.
func receiveData(ch <-chan int) {

	value := <-ch

	fmt.Println("Received:", value)
}

func main() {

	// Normal channel: can send and receive.
	ch := make(chan int)

	// Create a goroutine for sending.
	go sendData(ch)

	// Receive data.
	receiveData(ch)
}

