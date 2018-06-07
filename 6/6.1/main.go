package main

import (
	"fmt"
	. "./intset"
)

func main() {
	var x IntSet
	x.Add(1)
	x.Add(64)
	x.Add(500)
	fmt.Println(x)
	fmt.Println(x.String())
	fmt.Println(x.Len())
}
