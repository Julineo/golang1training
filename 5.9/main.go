package main

import (
//	"sort"
	"fmt"
)

func main() {
	fmt.Printf("%v", reverse("123"))
}

func expand(s string, f func(string) string) string {
	return "0"
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
