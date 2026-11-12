package main

import (
	"fmt"
	"log"

	"day8_employee_service/employee"
)

func main() {

	log.Println("[INFO] Employee application started")

	e := employee.Employee{
		ID:     1,
		Name:   "Vittesh",
		Salary: 30000,
	}

	err := employee.ValidateEmployee(e)

	if err != nil {
		log.Println("[ERROR]", err)
		return
	}

	fmt.Println("Employee ID:", e.ID)
	fmt.Println("Employee Name:", e.Name)
	fmt.Println("Employee Salary:", e.Salary)

	totalSalary := employee.CalculateSalary(
		e.Salary,
		5000,
	)

	fmt.Println("Total Salary:", totalSalary)

	log.Println("[INFO] Employee application completed")
}