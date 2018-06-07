package main

import (
	"fmt"
	. "./intset"
)

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(1)
	fmt.Println(y.String())
}
