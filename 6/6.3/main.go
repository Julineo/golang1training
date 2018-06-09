package main

import (
	"fmt"
	. "./intset"
)

func main() {
	var x, y IntSet
	x.AddAll(1,4,7,500)
	y.AddAll(0,4,10,100,500,1000)
	fmt.Println(x)
	fmt.Println(x.String())
	fmt.Println(y)
	fmt.Println(y.String())

	x.DifferenceWith(&y)
//	y.DifferenceWith(&x)

	fmt.Println(x)
	fmt.Println(x.String())
}
