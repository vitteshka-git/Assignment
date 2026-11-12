package main

import "fmt"

type Config struct {
	AppName string
	Port    string
	Env     string
	Debug   bool
}

func DevelopmentConfig() Config {
	return Config{
		AppName: "Employee API",
		Port:    "8080",
		Env:     "development",
		Debug:   true,
	}
}

func main() {
	config := DevelopmentConfig()

	fmt.Println("Application:", config.AppName)
	fmt.Println("Port:", config.Port)
	fmt.Println("Environment:", config.Env)
	fmt.Println("Debug Mode:", config.Debug)
}