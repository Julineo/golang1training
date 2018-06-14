// ex 7.2: Write a function CountingWriter with the signature below that, given an io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 variable that at any moment contains the number of bytes written to the new Writer

package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	w     io.Writer
	count int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := ByteCounter{w, 0}
	return &c, &c.count
}

func main() {
	writer, count := CountingWriter(os.Stdout)
	fmt.Fprint(writer, "Counting writer\n")
	fmt.Println(*count)
}
