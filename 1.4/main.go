//Read text files and print
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	arr := [][]string{}
	countsf := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, &arr)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, &arr)
			f.Close()
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i][1] == arr[j][1] {
				countsf[arr[i][0]]++
				countsf[arr[j][0]]++
			}

		}
	}
	for line, n := range countsf {
		if n > 2 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, arr *[][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		*arr = append(*arr, []string{f.Name(), input.Text()})
	}
	// potential errors from input.Err() ignored
}
