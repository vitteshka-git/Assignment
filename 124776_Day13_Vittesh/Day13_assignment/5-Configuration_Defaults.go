package main

import (
	"fmt"
	"os"
)

type Config struct {
	AppName string
	Port    string
	Env     string
}

func LoadConfig() Config {
	config := Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
		Env:     os.Getenv("APP_ENV"),
	}

	if config.AppName == "" {
		config.AppName = "Employee API"
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	if config.Env == "" {
		config.Env = "development"
	}

	return config
}

func main() {
	config := LoadConfig()

	fmt.Println("Application Name:", config.AppName)
	fmt.Println("Port:", config.Port)
	fmt.Println("Environment:", config.Env)
}