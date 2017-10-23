package main

import (
	"bufio"
	"fmt"
	//"io"
	"os"
	//"unicode"
	"strings"
)

func main() {
	//countsWords := make(map[string]int)

	scanner := bufio.NewScanner(strings.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text)
	}
}
