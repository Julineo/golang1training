package main

import (
	"testing"
)

func TestJoin(t *testing.T) {
	var tests = []struct {
		vals []string
		sep string
		want string
}{
		{[]string{"1","2","3"}, "," , "1,2,3"},
		{[]string{"-1"}, " ", "-1"},
		{[]string{"asd", "asdf", ""}, "", "asdasdf"},
		{[]string{}, "", ""},
	}


	for _, test := range tests {
		if got := Join(test.sep, test.vals...); got != test.want {
			t.Errorf("Join(%v, %v) = %v", test.vals, test.sep, got)
		}
	}

}
