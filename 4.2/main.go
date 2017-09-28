package main

import ("fmt"
		"crypto/sha256"
		"crypto/sha512"
		"bufio"
		"os"
	)

func main() {
	flags := os.Args[1:]
	var flag string
	
	if len(flags) == 0 {
		flag = ""
	} else {
		flag = flags[0]
	}
	
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("%v\n", flag)
	for input.Scan() {
		switch flag {
			case "SHA384":
				c1 := sha512.Sum384([]byte(input.Text()))
				fmt.Printf("%x\n", c1)
			case "SHA512":
				c1 := sha512.Sum512([]byte(input.Text()))
				fmt.Printf("%x\n", c1)
			default: 
				c1 := sha256.Sum256([]byte(input.Text()))
				fmt.Printf("%x\n", c1)
		}
	}
}