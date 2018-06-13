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
	i := 0
	count := 0
	prev := ' '
	for i < len(p) {
		curr, j := utf8.DecodeRune(p[i:])
		if (prev != ' ' && curr == ' ') || (prev != ' ' && i + j == len(p)) {
			count++
		}
		i += j
		prev = curr
	}
	*d = WordsCounter(count)
	return count, nil
}

type LinesCounter int

func (c *LinesCounter) Write(p []byte) (int, error) {
	count := 0
	for _,v := range p {
		if v == 10 {
			count++
		}
	}
	*c = LinesCounter(count + 1)
	return count + 1, nil
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
	d.Write([]byte("  I say 世界 to you"))
	fmt.Println(d)	// "5"

	var e LinesCounter
	s := `line1
		line2`
	e.Write([]byte(s))
	fmt.Println(e)

	s = "line1"
	e.Write([]byte(s))
	fmt.Println(e)

}
