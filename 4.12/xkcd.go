package main

import (
	"os"
	"fmt"
	"encoding/gob"
)

var usage string = `usage:
xkcd [search words]
`
func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

func loadIdx(file string) (idx map[string][]int) {
        f, err := os.Open(file)
        if err != nil {
                panic("cant open file")
        }
        defer f.Close()

        enc := gob.NewDecoder(f)
        if err := enc.Decode(&idx); err != nil {
                panic("can't decode")
        }

        return idx
}

func main() {
	var idx map[string][]int
	idx = loadIdx("idx")
	
	fmt.Println(idx["you"])
}