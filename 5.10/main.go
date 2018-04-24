package main

import (
	"fmt"
	"sort"
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


func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	fmt.Println(keys)

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func topoSort2(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func([]string)

	visitAll = func(items []string) {
		for _, v := range items {
			if !seen[v] {
				seen[v] = true
				visitAll(m[v])
				order = append(order, v)
			}
		}
	}

	for k := range m {
		visitAll([]string{k})
	}
	return order
}
