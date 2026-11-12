package main

import "fmt"

type Config struct {
	AppName string
	Port    string
	Env     string
}

func main() {
	config := Config{
		AppName: "Employee API",
		Port:    "8080",
		Env:     "development",
	}

	fmt.Println("Application Name:", config.AppName)
	fmt.Println("Port:", config.Port)
	fmt.Println("Environment:", config.Env)
}