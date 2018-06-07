package main

import (
	"fmt"
	. "./intset"
)

func main() {
	var x IntSet
	x.AddAll(1,4,7)
	fmt.Println(x.String())
}
