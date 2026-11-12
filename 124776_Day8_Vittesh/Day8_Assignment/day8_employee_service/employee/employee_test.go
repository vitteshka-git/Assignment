package employee

import "testing"

func TestValidateEmployee(t *testing.T) {

	employee := Employee{
		ID:     1,
		Name:   "Vittesh",
		Salary: 30000,
	}

	err := ValidateEmployee(employee)

	if err != nil {
		t.Errorf("Expected employee to be valid")
	}
}

func TestValidateEmployeeEmptyName(t *testing.T) {

	employee := Employee{
		ID:     1,
		Name:   "",
		Salary: 30000,
	}

	err := ValidateEmployee(employee)

	if err == nil {
		t.Errorf("Expected validation error")
	}
}

func TestCalculateSalary(t *testing.T) {

	result := CalculateSalary(30000, 5000)

	expected := 35000.0

	if result != expected {
		t.Errorf(
			"Expected %.2f but got %.2f",
			expected,
			result,
		)
	}
}