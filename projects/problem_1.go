package projects

import (
	"fmt"
)

// Problem1: find the sum of all the multiples of 3 or 5 below 1000.
// x,y: multiples
// z: upper limit
func Problem1(x, y, z int) int {

	// protects against collisions for the multiples key.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Number exists!", r)
		}
	}()

	Multiples := make(map[int]struct{})

	var Sum int
	for i := 1; i < z; i++ {
		XFactor := x * i
		YFactor := y * i

		if XFactor < z {
			Multiples[XFactor] = struct{}{}
		}
		if YFactor < z {
			Multiples[YFactor] = struct{}{}
		}
	}

	for key, _ := range Multiples {
		Sum = Sum + key
	}

	return Sum
}
