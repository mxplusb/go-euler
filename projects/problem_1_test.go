package projects

import "testing"

func BenchmarkProblem1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem1(3, 5, 1000)
	}
}
