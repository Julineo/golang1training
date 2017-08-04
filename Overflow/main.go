package main

import (
	"fmt"
)

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	fmt.Printf("u: %b\n", 65025)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)
	fmt.Printf("i: %b\n", i)

	//bitwise operations
	var x uint8 = 1<<1 | 1<<5
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	//bitwise shift
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", x<<5)

	//signed bitshifts
	var j int8 = +1
	fmt.Printf("%08b\n", j)
	fmt.Printf("%08b\n", j>>5)
}
