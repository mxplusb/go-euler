package projects

import "testing"

func TestProblem3(t *testing.T) {
	testNumber := 13195
	want := 29
	got := Problem3(testNumber)
	if want != got {
		t.Fatalf("wanted: 29, got: %d", got)
	}
}

func BenchmarkProblem3(b *testing.B) {
	b.N = 5 // this will timeout after 30 minutes, otherwise.
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		Problem3(600851475143)
	}
}
