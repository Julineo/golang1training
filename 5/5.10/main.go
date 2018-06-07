package main

import (
	"fmt"
	//"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

//managed to change slices with map, although it wasn't good idea
func topoSort(m map[string][]string) []string {
	var list []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]string)

	visitAll = func(items map[string]string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				keys2 := make(map[string]string)
				for _, key2 := range m[item] {
					keys2[key2] = key2
				}
				visitAll(keys2)
				list = append(list, item)
			}
		}
	}

	keys := make(map[string]string)
	for key := range m {
		keys[key] = key
	}

	visitAll(keys)
	return list
}
