package helpers

func Even(n int) bool {
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
