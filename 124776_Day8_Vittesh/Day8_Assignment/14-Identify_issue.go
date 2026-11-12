package main

import (
	"fmt"
)

var employees []Employee

type Employee struct {
	name   string
	salary float64
}

func addEmployee(name string, salary float64) {

	employees = append(employees, Employee{
		name:   name,
		salary: salary,
	})

	fmt.Println("Employee added")
}

func calculate(a float64, b float64) float64 {
	return a + b
}

func main() {

	addEmployee("Vittesh", 30000)

	for i := 0; i < len(employees); i++ {
		fmt.Println(employees[i])
	}

}