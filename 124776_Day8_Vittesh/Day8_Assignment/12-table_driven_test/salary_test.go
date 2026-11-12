package main

import "testing"

func TestCalculateSalary(t *testing.T) {

	tests := []struct {
		name     string
		basic    float64
		bonus    float64
		expected float64
	}{
		{
			name:     "Normal salary",
			basic:    30000,
			bonus:    5000,
			expected: 35000,
		},
		{
			name:     "Higher salary",
			basic:    50000,
			bonus:    10000,
			expected: 60000,
		},
		{
			name:     "No bonus",
			basic:    30000,
			bonus:    0,
			expected: 30000,
		},
		{
			name:     "Small salary",
			basic:    10000,
			bonus:    2000,
			expected: 12000,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			result := CalculateSalary(
				test.basic,
				test.bonus,
			)

			if result != test.expected {
				t.Errorf(
					"Expected %.2f but got %.2f",
					test.expected,
					result,
				)
			}
		})
	}
}