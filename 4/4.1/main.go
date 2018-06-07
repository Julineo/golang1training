package main

import "fmt"

import ("crypto/sha256"
	//"golang1training/4.1/popcount"//this is just for practice
	)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
	
    n := 0
	for i := range c1 {
		for j := uint(0); j < 64; j++ {
			//fmt.Printf("%8b\n", c1[i])
			//fmt.Printf("%8b\n", c2[i])
			//fmt.Printf("%8b\n", 1<<j)
			if c1[i]&(1<<j) != c2[i]&(1<<j) {
				n++
			}
		}
	}
	fmt.Println(n)
	//fmt.Printf("%v", popcount.PopCount(255))
}