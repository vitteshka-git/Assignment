/*
package main

import (
	"fmt"
	"sync"
)

var count int

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		count++
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go increase(&wg)
	go increase(&wg)

	wg.Wait()

	fmt.Println("Count:", count)
}
*/
package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {

		mutex.Lock()

		count++

		mutex.Unlock()
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go increase(&wg)
	go increase(&wg)

	wg.Wait()

	fmt.Println("Count:", count)
}

/*
PS C:\Users\vittesh.ka\Downloads\GO\Day8_Assignment> go run 8-race_detection.go
Count: 2000
PS C:\Users\vittesh.ka\Downloads\GO\Day8_Assignment> go run -race  8-race_detection.go
go: -race requires cgo; enable cgo by setting CGO_ENABLED=1
PS C:\Users\vittesh.ka\Downloads\GO\Day8_Assignment> 
*/