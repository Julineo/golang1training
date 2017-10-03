package main

import (
	"fmt"
	"bytes"
)

func main() {
	rs := []rune{'世','g','o',' ', ' ','l','a', ' ' ,'n'}
	bs := []byte(string(rs))
	fmt.Println(rs)
	fmt.Println(bs)
	fmt.Printf("%08b\n,%T\n",bs,bs)
	reverse(bs)
	fmt.Println(bs)
}

func reverseUTF8(s []byte) {
	for i :=0; len(s)-1; i < j; i++ {
		s[i], s[j] = byte(bytes.Runes(s)[j]), byte(bytes.Runes(s)[i])
		
	}
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}