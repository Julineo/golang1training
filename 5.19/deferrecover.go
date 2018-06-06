/*
Exercise 5.19: Use panic and recover to write a function that contains no return statement yet returns a non-zero value.
*/

package main

import "fmt"

func main() {
	fmt.Println(f(3))
}

func f(x int) (s string) {
	defer func() {
		recover()
		s = "division by zero"
	}()
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
	panic("No return")
}

