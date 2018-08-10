package projects

import (
	"github.com/mxplusb/go-euler/projects/helpers"
	)

func Problem3(x int) int {
	largestPrimeFactor := 0
	for i := 1; i < x; i++ {
		if x % i == 0 {
			// TODO (mxplusb): write some timeout logic and implement a quadratic sieve.
			if helpers.MillerRabinIsPrime(i, 4) {
				//fmt.Printf("Found prime: %d\n", i)
				largestPrimeFactor = i
			}
		}
	}
	return largestPrimeFactor
}
