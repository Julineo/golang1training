package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countsWords := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		countsWords[scanner.Text()]++
	}
	for c, n := range countsWords {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
