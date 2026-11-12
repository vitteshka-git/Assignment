package employee

import "testing"

func BenchmarkCalculateSalary(b *testing.B) {

	for i := 0; i < b.N; i++ {
		CalculateSalary(30000, 5000)
	}
}