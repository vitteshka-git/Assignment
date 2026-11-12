package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	AppName string
	Port    string
}

func ValidateConfig(config Config) error {
	if strings.TrimSpace(config.AppName) == "" {
		return fmt.Errorf("application name cannot be empty")
	}

	port, err := strconv.Atoi(config.Port)
	if err != nil {
		return fmt.Errorf("port must be a number")
	}

	if port < 1 || port > 65535 {
		return fmt.Errorf("port must be between 1 and 65535")
	}

	return nil
}

func main() {
	config := Config{
		AppName: "Employee API",
		Port:    "8080",
	}

	err := ValidateConfig(config)
	if err != nil {
		fmt.Println("Configuration error:", err)
		return
	}

	fmt.Println("Configuration is valid")
	fmt.Println("Application:", config.AppName)
	fmt.Println("Port:", config.Port)
}