package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		vals []int
		want int
	}{
		{[]int{1,2,3}, 3},
		{[]int{-1}, -1},
		{[]int{0}, 0},
	//	{[]int{}, 0},
	}

	for _, test := range tests {
		if got := max(test.vals...); got != test.want {
			t.Errorf("max(%d) = %d", test.vals, got)
		}
	}

}
