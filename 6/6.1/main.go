package main

import (
	"fmt"
	. "./intset"
)

func main() {
	var x, y IntSet
	x.Add(0)
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(4)
	x.Add(5)
	x.Add(6)
	x.Add(7)
	fmt.Println(x)
	fmt.Println(x.String())

	fmt.Println(x.Len())

	x.Remove(2)
	fmt.Println(x)
	fmt.Println(x.String())

	fmt.Println("Copy: ", x.Copy())
	y = *x.Copy()

	x.Clear()
	fmt.Println(x)
	fmt.Println(x.String())
	x.Add(1)

	fmt.Println(x.String())
	fmt.Println(y.String())
}
