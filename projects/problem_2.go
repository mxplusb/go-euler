package projects

// Problem2: by considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of
// the even-valued terms.
func Problem2(x int) int {
	if x == 0 {
		return 0
	}

	Left, Right := 0, 1
	var Sum int
	for Right < x {
		Left, Right = (Left + (Right * 2)), ((Left * 2) + (Right * 3))
		Sum = Sum + Left
	}

	return Sum
}
