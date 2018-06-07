package main

import (
	"fmt"
	"os"
	"bytes"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	var n int
	for i := len(s); i > 0; i-- {
		if n == 3 { 
			buf.WriteString(",")
			n = 0
		} 
		n++
		v := s[i-1:i]
		fmt.Fprintf(&buf, "%v", string(v))
	} 
	return reverse(buf.String() )
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}