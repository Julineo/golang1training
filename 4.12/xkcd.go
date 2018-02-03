package main

import (
	"os"
)

var usage string = `usage:
xkcd [search words]
`
func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

func main {

}