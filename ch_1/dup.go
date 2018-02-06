package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//counts[input.Text()]++
		line := input.Text()
		counts[line] = counts[line] + 1
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println("%d%t%s", n, line)
		}
	}
}