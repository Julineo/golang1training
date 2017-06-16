//Echo prints its command-line arguments. And shows times of multiple executions for different approaches.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	i := 0

	start := time.Now()
	for i = 1; i <9999990; i++ {
		ex1()
	}
	resex1 := time.Since(start).Seconds()

	start = time.Now()
	for i = 1; i <9999990; i++ {
		ex2()
	}
	resex2 := time.Since(start).Seconds()

	start = time.Now()
	for i = 1; i <9999990; i++ {
		ex3()
	}
	resex3 := time.Since(start).Seconds()

	fmt.Println(resex1)
	fmt.Println(resex2)
	fmt.Println(resex3)
}

func ex1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func ex2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func ex3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

