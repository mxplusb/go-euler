package projects

import "testing"

func BenchmarkProblem2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem2(2147483647)
	}
}
