package main

import (
	"fmt"
	. "./intset"
)

func main() {
	var x IntSet
	x.AddAll(1,4,7,500)
	fmt.Println(x.Elems())
	fmt.Println(x.String())
	y := x.Copy()
	fmt.Println(y)
}
