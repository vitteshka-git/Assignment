package main
import "fmt"

func main() {

	/*
		Buffered channel can store 3 values.
	*/
	ch := make(chan int, 3)

	/*
		This creates a goroutine.
		The goroutine sends values into the channel.
	*/
	go func() {

		ch <- 10
		ch <- 20
		ch <- 30

	}()

	/*
		Main goroutine receives the values.
	*/
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("Main completed")
}

