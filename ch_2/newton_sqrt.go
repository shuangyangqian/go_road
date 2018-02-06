package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for Abs(z*z-x) > 1e-9 {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

func main() {
	fmt.Println(Abs(-20))
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
