package employee

import "testing"

func TestValidateEmployee(t *testing.T) {

	employee := Employee{
		Name:   "Vittesh",
		Salary: 30000,
	}

	err := ValidateEmployee(employee)

	if err != nil {
		t.Error("Employee should be valid")
	}
}
