package main

import (
	"fmt"
	"os"
	"bytes"
	"strings"
	"strconv"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf1 bytes.Buffer
	//var buf2 bytes.Buffer
	var s1, s2 string
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s1 = s[:dot]
		s2 = s[dot+1:]
		//fmt.Println(s1)
		fmt.Println(s2)
	}
	var n int
	for i := len(s1); i > 0; i-- {
		notSign := false
		if _, err := strconv.Atoi(v); err == nil {
			notSign = true
		}
		if n >= 3 && notSign { 
			buf1.WriteString(",")
			n = 0
		}
		v := s[i-1:i]
		n++
		fmt.Fprintf(&buf1, "%v", string(v))
	} 
	return reverse(buf1.String() )
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}