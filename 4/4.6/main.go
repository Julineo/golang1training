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
	bs = bytes.Replace(bs, []byte("  "), []byte(" "), -1)
	fmt.Println(bs)
}