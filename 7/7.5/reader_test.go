package reader

import (
	"bytes"
	"strings"
	"testing"

)

func TestNewReader(t *testing.T) {
	s := "limit reader"
	b := &bytes.Buffer{}
	l := LimitReader(strings.NewReader(s), 5)
	n, _ := b.ReadFrom(l)
	if n != 5 {
		t.Logf("n=%d err", n)
		t.Fail()
	}
	if b.String() != "limit" {
		t.Logf("%s != %s", b.String(), s)
	}
}

