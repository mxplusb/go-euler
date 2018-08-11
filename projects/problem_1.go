package projects

import (
	"github.com/mxplusb/go-euler/helpers"
)

// Problem1: find the sum of all the multiples of 3 or 5 below 1000.
// x,y: multiples
// z: upper limit
func Problem1(x, y, z int) int {
	var sumX int = helpers.ArithmeticSum(0, x, (z-1)/x+1)
	var sumY int = helpers.ArithmeticSum(0, y, (z-1)/y+1)
	var lcm int = helpers.LCM(x, y)
	var sumBoth int = helpers.ArithmeticSum(0, lcm, (z-1)/lcm+1)
	return sumX + sumY - sumBoth
}
