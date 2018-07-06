/* Ex. 7.4
The strings.NewReader function returns a value that satisfies the io.Reader interface (and others) by reading from its argument, a string. Implement a simple version of NewReader yourself, and use it to make the HTML parser (§5.2) take input from a string.
*/
package reader

import (
	"io"
)

type stringReader struct {
	s string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF // Must set EOF, otherwise it does not end
	}
	return
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}
