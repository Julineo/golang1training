package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
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
		case unicode.IsNumber(r):
			fmt.Println("IsNumber")
		case unicode.IsLetter(r):
			fmt.Println("IsLetter")
		default:
			fmt.Println("default")
		}
	}
}
