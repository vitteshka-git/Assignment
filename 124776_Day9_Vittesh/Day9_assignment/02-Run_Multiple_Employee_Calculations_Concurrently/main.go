package main
import (
	"fmt"
	"time"
)

// Employee represents an employee.
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// calculateSalary performs a calculation for one employee.
func calculateSalary(emp Employee) {

	// Simulate some time-consuming work.
	time.Sleep(1 * time.Second)

	// Calculate 10% bonus.
	bonus := emp.Salary * 0.10

	// Calculate final salary.
	finalSalary := emp.Salary + bonus

	fmt.Printf(
		"Employee: %s | Salary: %.2f | Bonus: %.2f | Final Salary: %.2f\n",
		emp.Name,
		emp.Salary,
		bonus,
		finalSalary,
	)
}

func main() {

	employees := []Employee{
		{1, "vittesh", 30000},
		{2, "dhanu", 40000},
		{3, "guru", 50000},
	}

	/*
		CONCURRENT EXECUTION:

		Using "go" before a function call starts that function
		in a separate goroutine.

		So all three employee calculations can run concurrently.

		    Employee 1 ──> Goroutine 1
		    Employee 2 ──> Goroutine 2
		    Employee 3 ──> Goroutine 3

		They do NOT have to wait for each other.
	*/

	for _, emp := range employees {
		go calculateSalary(emp)
	}

	/*
		We wait for 2 seconds so that the goroutines
		have enough time to finish.

		In real applications, instead of time.Sleep(),
		use sync.WaitGroup for proper synchronization.
	*/
	time.Sleep(2 * time.Second)

	fmt.Println("All calculations completed")
}


