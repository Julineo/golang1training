package main

import (
	"testing"
	"strings"
)

func TestExpand(t *testing.T) {
	var tests = []struct {
			s string
			functn func(string) string
			want string
		}{
			{"$foolish big$foot", reverse, "ooflish bigooft"},
			{"$foolish big$foot", strings.Title, "Foolish bigFoot"},
		}

	for _, test := range tests {
		if got := expand(test.s, test.functn); got != test.want {
			t.Errorf("expand(%s) = %s" , test.s, got)
		}
	}
}
