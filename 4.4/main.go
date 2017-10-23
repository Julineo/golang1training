package main

import (
	"fmt"
)

func main() {
	var r int = 2//rotate by this number
	c := []int{0, 1, 2, 3, 4, 5}
	c = append(c[r:],c[:r]...)
	fmt.Println(c)
}
