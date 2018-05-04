package main

import (
	"fmt"
	"sort"
)

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

var prereqs2 = map[string][]string{
	"algorithms": {"data structures"},
	"linear algebra":	{"calculus"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming","algorithms"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	for i, course := range topoSort(prereqs2) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}


}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	index := make(map[string]int)
	graph := make([][]int,2)
	fmt.Println(graph)
	var visitAll func(items []string)
	var idx int

	visitAll = func(items []string) {
		fmt.Println("items: ", items)
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				fmt.Println("add to order:", item)
				order = append(order, item)
				index[item] = idx
				idx++
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	fmt.Println("keys: ", keys)
	visitAll(keys)
	fmt.Println(index)
	return order
}
