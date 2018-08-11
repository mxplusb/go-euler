package algorithms

import "math/big"

type Sieve struct {
	internalBaseStringRef string
	internalBigIntRef *big.Int
}

// Takes a new base 10 string of integers to sift through to factorize. Will use all available processors.
func NewQuadraticSieve(s string) *Sieve {
	i := new(big.Int)
	i.SetString(s, 10)
	return &Sieve{
		internalBaseStringRef: s,
		internalBigIntRef: i,
	}
}

//func (s *Sieve) Sift() string {
//	var microBase uint64
//	var nbPrimes int64
//	var i, j uint64
//	sqrtN, rem := new(big.Int), new(big.Int)
//
//}