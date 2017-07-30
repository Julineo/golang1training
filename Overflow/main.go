package main

import (
	"fmt"
)

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	fmt.Printf("u: %b\n", u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)
	fmt.Printf("i: %b\n", i)
}
