package main

import (
	//"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var n []string
	n = strings.Fields(s)
	var m map[string]int
	m = make(map[string]int)
	for _, word := range n {
		_, ok := m[word]
		if ok {
			m[word] += 1
		} else {
			m[word] = 1
		}

	}
	return m
	//return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
	//fmt.Println(WordCount("thuisi is is is is "))
}
