package main

import (
	"fmt"
	"os"
	"sync"
)


func processFile(fileName string, wg *sync.WaitGroup) {

	// Tell WaitGroup that this goroutine is finished.
	defer wg.Done()

	fmt.Println("Processing:", fileName)

	/*
		Read the existing file.

		We are NOT creating the file here.
		The file is already created before processing.
	*/
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(
		"Completed:",
		fileName,
		"| Size:",
		len(data),
		"bytes",
	)
}

func main() {

	/*
		CREATE SAMPLE FILES

		os.WriteFile() is used here.

		If the file does not exist:
			→ It creates the file.

		If the file already exists:
			→ It replaces the existing content.

		For a real file processor, we would normally
		process existing files instead of creating them.
	*/

	os.WriteFile("file1.txt", []byte("Hello Go"), 0644)

	os.WriteFile("file2.txt", []byte("Go Programming"), 0644)

	os.WriteFile("file3.txt", []byte("Concurrency"), 0644)

	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	var wg sync.WaitGroup

	/*
		Process each file concurrently.

		"go" creates a goroutine.

		Each file gets its own goroutine.
	*/

	for _, fileName := range files {

		wg.Add(1)

		go processFile(fileName, &wg)
	}

	/*
		Wait until all goroutines finish.
	*/
	wg.Wait()

	fmt.Println("All files processed")
}

