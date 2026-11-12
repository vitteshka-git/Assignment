package employee

import "testing"

func TestCalculateSalary(t *testing.T) {

	result := CalculateSalary(30000, 5000)

	expected := 35000.0

	if result != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, result)
	}
}