package main

import (
	"fmt"
	"bytes"
	"unicode/utf8"
)

func main() {
	rs := []rune{'世','g','o',' ', ' ','l','a', '界', ' ' ,'n'}
	bs := []byte(string(rs))
	fmt.Println(rs)
	reverseUTF8(bs)
	reverse(bs)
	fmt.Println(bytes.Runes(bs))
}

//reverse runes only
func reverseUTF8(s []byte) {
	for len(s) > 0 {
		_, size := utf8.DecodeRune(s)
		if size > 1 {
			reverse(s[:size])
		}
		s = s[size:]
	}
}

//reverse whole slice
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}