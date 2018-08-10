package helpers

import (
	"math/rand"
)

// ModularExponentiation is just (x^y) % p
func ModularExponentiation(x, y, p int) int {
	var result = 1
	x = x % p

	for y > 0 {
		if !IntEven(y) {
			result = (result * x) % p
		}
		y = y >> 1 // y = y/2
		x = (x * x) % p
	}
	return result
}

// MillerRabinTest
func MillerRabinTest(d, n int) bool {
	// random number in [2..n-2]
	var a = 2 + rand.Int()%(n-4)

	// compute a^d % n
	var x = ModularExponentiation(a, d, n)
	if x == 1 || x == n-1 {
		return true
	}

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

func MillerRabinIsPrime(n, k int) bool {
	if n <= 1 || n == 4 {
		return false
	}
	if n <= 3 {
		return true
	}

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
