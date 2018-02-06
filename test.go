package main

import (
	"fmt"
)

func main() {

	var m map[string]int

	m = make(map[string]int)
	m["sqian"] = 12
	m["sqians"] = 123
	fmt.Println(m)

	// ages := map[string]int{
	// 	"alice": 32,
	// 	"joy":   33,
	// }
	// ages := make(map[string]int)
	// ages["alice"] = 21
	// ages["joy"] = 22
	// delete(ages, "joy")
	// for name, age := range ages {
	// 	fmt.Printf("%s\t%d\n", name, age)
	// }

	// a := [...]int{99: -1}
	// fmt.Println(a[0])
	// fmt.Println(len(a))
	// for i, v := range a {
	// 	fmt.Printf("%d %d\n", i, v)
	//}
}
