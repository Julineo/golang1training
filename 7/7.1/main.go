/*
ex 7.1: Using the ideas from ByteCounter, implement counters for words and for lines. You will find bufio.ScanWords useful.
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordsCounter int

func (d *WordsCounter) Write(p []byte) (int, error) {
	cur := 0
	for cur < len(p) {
		r, size := utf8.DecodeRune(p[cur:])
		fmt.Println(r, size)
	}
	return 0, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

	var d WordsCounter
	d.Write([]byte("  one two thre  e"))
	fmt.Println(d)
}
