package helpers

import (
	"math/big"
	)

func IntEven(n int) bool {
	return n%2 == 0
}

func UintEven(n uint) bool {
	return n%2 == 0
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}

func ArithmeticSum(first, step, count int) int {
	var last int = first + step*(count-1)
	return (first + last) * count / 2
}

func Factorial(n *big.Int) (result *big.Int) {
	result = new(big.Int)

	switch n.Cmp(&big.Int{}) {
	case -1, 0:
		result.SetInt64(1)
	default:
		result.Set(n)
		var one big.Int
		one.SetInt64(1)
		result.Mul(result, Factorial(n.Sub(n, &one)))
	}
	return
}

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

//func ModularSquareRoot(a, p int) int {
//
//}
//
//func LegendreSymbol(a, p int) int {
//	ls := math.Pow(float64(a), (float64(p)-1)/2)
//}
