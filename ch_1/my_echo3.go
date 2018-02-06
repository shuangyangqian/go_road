package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println(os.Args)
}