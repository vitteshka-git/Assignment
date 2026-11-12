
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	done := make(chan string, 1)

	go func() {
		time.Sleep(3 * time.Second)
		done <- "Work completed"
	}()

	select {
	case result := <-done:
		fmt.Println(result)

	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}
}
