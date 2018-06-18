// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import (
	"fmt"
	//"bytes"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// reveals the sequence of values in the tree
func (t *tree) String() (res string) {
	var visit func(t *tree)
	visit = func(t *tree) {
		if t.left != nil {
			visit(t.left)
		}
		res = fmt.Sprintf("%s %d", res, t.value)
		if t.right != nil {
			visit(t.right)
		}
	}
	visit(t)
	return res
}
