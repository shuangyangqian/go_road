package main

import "fmt"

func Fibonacci(x int) int {
	if x == 1 {
		return 1
	} else if x == 2 {
		return 1
	} else if x >= 3 {
		return Fibonacci(x-1) + Fibonacci(x-2)
	}
	return 1
}

func main() {
	fmt.Println(Fibonacci(12))
}
