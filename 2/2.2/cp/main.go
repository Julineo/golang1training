// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"

	"gopl.ex/2.2/tempconv"
)

func main() {
	fmt.Printf("cf: %v\n", tempconv.EqualT)//check if this is correct library
	if len(os.Args[1:]) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			targ := scanner.Text()
			convertAll(targ)
		}
	} else {
		for _, arg := range os.Args[1:] {
			convertAll(arg)
		}
	}
}

func convertAll(arg string) {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
}
