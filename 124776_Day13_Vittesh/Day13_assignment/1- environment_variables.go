package main

import (
	"fmt"
	"os"
)

func main() {
	appName := os.Getenv("APP_NAME")
	port := os.Getenv("PORT")

	fmt.Println("Application Name:", appName)
	fmt.Println("Port:", port)
}