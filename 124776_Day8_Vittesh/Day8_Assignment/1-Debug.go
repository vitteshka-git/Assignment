package main 

import "fmt"

type Employee struct {
	ID int
	Name string
	Salary float64
}

func calculateBonus(salary float64) float64 {
	bonus := salary * 0.10

	return bonus
}

func updatesalary(employee *Employee, increment float64) {
	employee.Salary = employee.Salary + increment
}

func main() {
	employee := []Employee {
		{
		ID : 101,
		Name : "Vittesh",
		Salary: 30000,
	},
	{
		ID : 102,
		Name : "Dhanu",
		Salary: 30000,
	},
}


fmt.Println("Employee management Application")

for _, employee := range employee {
	fmt.Println("Employee ID:", employee.ID)
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Employee Salary:", employee.Salary)

	bonus := calculateBonus(employee.Salary)

	fmt.Println("Bonus:", bonus)

	updatesalary(&employee, 5000)

	fmt.Println("updated salary:", employee.Salary)
}

}