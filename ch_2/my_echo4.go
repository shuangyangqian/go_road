package main

import (
	"flag"
	"fmt"
	"strings"
)


func main() {
	var results = fib(20)
	fmt.Println(results)
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i <n; i++ {
		x, y = y, x+y
	}
	return x
}

//求两个数的最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func my_flag() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}


func my_fmt() {
	var i int = 123
	var p *int = &i
	var q **int = &p
	var x int = *p

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(i)
	fmt.Println(*p)
	fmt.Println(*q)
	fmt.Println(x)
}
