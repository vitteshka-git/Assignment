package main 
import ( 
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.RWMutex // create a read /write mutex
	data := "Employee Data"


	var wg sync.WaitGroup
	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go func(id int) {
			defer wg.done()
			mu.RLock()
			fmt.Println("Reader" , id "read:" , data)
			time.Sleep(100 * time.Millisecond)
			mu.RUNlock()
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond)
		mu.Lock() // Gets exclusive write access.
		data = "Updated Employee Data" // Changes shared data.
		fmt.Println("Writer updated data") // Prints update information.
		mu.Unlock() // Releases the write lock.
	}()

	wg.Wait() // Waits for all readers and writer.
	fmt.Println("Finished") // Prints when all operations complete.
	
}