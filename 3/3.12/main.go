package main

import (
	"fmt"
)

func main() {
	//check if strings are anagrams
	s1, s2 := "strings", "sgnirts"
	if reverse(s1) == s2 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}