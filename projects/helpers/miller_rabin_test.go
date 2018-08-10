package helpers

import (
	"testing"
	"reflect"
)

func TestMillerRabinIsPrime(t *testing.T) {
	var iterations = 4
	// known primes smaller than 100
	expectedPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	receivedPrimes := make([]int, 0)
	for n := 0; n < 100; n++ {
		if MillerRabinIsPrime(n, iterations) {
			receivedPrimes = append(receivedPrimes, n)
		}
	}
	if !reflect.DeepEqual(expectedPrimes, receivedPrimes) {
		t.Fatalf("failed known primes test. got: %#v", receivedPrimes)
	}
}

func BenchmarkMillerRabinIsPrime(b *testing.B) {
	b.Log("Testing performance for all Miller-Rabin Primes in [1..1000000].")
	for i := 0; i < b.N; i++ {
		MillerRabinIsPrime(1000000, 4)
	}
}