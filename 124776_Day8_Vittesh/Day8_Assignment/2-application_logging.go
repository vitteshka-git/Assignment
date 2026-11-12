package main
import "log"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {

	log.Println("Application started")

	employee := Employee{
		ID:     1,
		Name:   "Vittesh",
		Salary: 30000,
	}

	log.Println("Employee created")

	log.Printf("Employee ID: %d", employee.ID)
	log.Printf("Employee Name: %s", employee.Name)
	log.Printf("Employee Salary: %.2f", employee.Salary)

	log.Println("Application completed")
}