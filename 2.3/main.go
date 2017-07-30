// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"

	"gopl.ex/2.3/popcount"
)

func main() {
	var pc [256]byte

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("pc[i]: %v\n", pc[i])
	}
	fmt.Printf("255: %b\n", byte(255))
	fmt.Printf("PopCount: %b\n", popcount.PopCount(1))
}
