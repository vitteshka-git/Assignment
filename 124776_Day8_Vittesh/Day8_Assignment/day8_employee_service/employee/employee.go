package employee

import (
	"errors"
	"log"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func ValidateEmployee(e Employee) error {

	if e.Name == "" {
		return errors.New("employee name is required")
	}

	if e.Salary <= 0 {
		return errors.New("salary must be greater than 0")
	}

	return nil
}

func CalculateSalary(basic float64, bonus float64) float64 {

	log.Println("[INFO] Calculating salary")

	return basic + bonus
}