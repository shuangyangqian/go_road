package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("I am a hero")
}

func my_float() {
	for x := 0; x < 100; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func my_uint() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
}
