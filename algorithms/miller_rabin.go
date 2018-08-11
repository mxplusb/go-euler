package algorithms

import (
	"math/rand"
	"github.com/mxplusb/go-euler/helpers"
)

// MillerRabinTest is called for all k trials. It returns false if n is composite and returns false if n is probably
// prime. d is an odd number such that  d*2^r = n-1 for some r >= 1
func MillerRabinTest(d, n int) bool {
	// random number in [2..n-2]
	var a = 2 + rand.Int()%(n-4)

	// compute a^d % n
	var x = helpers.ModularExponentiation(a, d, n)
	if x == 1 || x == n-1 {
		return true
	}

	// Keep squaring x while one of the following doesn't
	// happen
	// (i)   d does not reach n-1
	// (ii)  (x^2) % n is not 1
	// (iii) (x^2) % n is not n-1
	for d != n-1 {
		x = (x * x) % n
		d *= 2
		if x == 1 {
			return false
		}
		if x == n-1 {
			return true
		}
	}
	return false
}

// It returns false if n is composite and returns true if n is probably prime. k is an input parameter that determines
// accuracy level. Higher value of k indicates more accuracy.
func MillerRabinIsPrime(n, k int) bool {
	// Edge cases
	if n <= 1 || n == 4 {
		return false
	}
	if n <= 3 {
		return true
	}

	// Find r such that n = 2^d * r + 1 for some r >= 1
	var d = n - 1
	for d%2 == 0 {
		d /= 2
	}

	for i := 0; i < k; i++ {
		if !MillerRabinTest(d,n) {
			return false
		}
	}
	return true
}
