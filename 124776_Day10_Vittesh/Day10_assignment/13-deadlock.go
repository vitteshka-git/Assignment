import (
	"fmt"             
	"sync"             
)

func main() {
	var mu sync.Mutex // Creates a mutex.
	done := make(chan bool) // Creates a channel for completion.

	go func() { // Starts a goroutine.
		mu.Lock() // Acquires the lock.
		fmt.Println("Goroutine acquired lock") // Prints lock acquisition.
		mu.Unlock() // Releases the lock before waiting.
		done <- true // Signals completion.
	}()

	<-done // Waits for the goroutine to finish.
	fmt.Println("No deadlock: lock was released") // Explains the safe pattern.
}
