/* 
ex 7.3 Write a String method for the *tree type in gopl.io/ch4/treesort (§4.4) that reveals the sequence of values in the tree.
*/
package treesort

import (
	"testing"
	"fmt"

//	"golang1training/7/7.3"
)

func TestString(t *testing.T) {

	// creates binary tree
	read := func() []tree {
		var input = []struct {
			val   int
			left  int
			right int
		}{
			{3, -1, -1},
			{6, -1, -1},
			{1, -1, -1},
			{4, 0, 1},
			{5, 2, 3},
		}

		var nodes []tree = make([]tree, len(input))

		var val, indexLeft, indexRight int
		for i := 0; i < len(input); i++ {
			val, indexLeft, indexRight = input[i].val, input[i].left, input[i].right
			nodes[i].value = val
			if indexLeft >= 0 {
				nodes[i].left = &nodes[indexLeft]
			}
			if indexRight >= 0 {
				nodes[i].right = &nodes[indexRight]
			}
		}

		return nodes
	}

	ti := read()
	fmt.Println(&ti[len(ti)-1])//root
	add(&ti[len(ti)-1], 10)
	fmt.Println(&ti[len(ti)-1])
}
