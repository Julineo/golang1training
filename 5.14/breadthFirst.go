//5.14: Use the breadthFirst function to explore a different structure. For example, you could use the course dependencies from the topoSort example
package main

import (
	"fmt"
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

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//returns prereqs of a speciphic course
func getPrereqs(course string) []string {
	fmt.Println(course)
	list := prereqs[course]
	return list
}

func main() {
	//print all prereqs by courses
	breadthFirst(getPrereqs, []string{"networks", "calculus"})
	fmt.Println("")
	breadthFirst(getPrereqs, []string{"calculus"})
	fmt.Println("")
	breadthFirst(getPrereqs, []string{"compilers"})

}
