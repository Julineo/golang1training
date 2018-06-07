package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	countsCat := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		switch true {
		case unicode.IsLetter(r):
			countsCat["IsLetter"]++
		case unicode.IsNumber(r):
			countsCat["IsNumber"]++
		case unicode.IsControl(r):
			countsCat["IsControl"]++
		default: 
			countsCat["Other"]++
		}
	}
	fmt.Printf("function\tcount\n")
	for cat, n := range countsCat {
		fmt.Printf("%q\t%d\n", cat, n)
	}
}
