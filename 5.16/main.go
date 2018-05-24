/* 5.16 variadic strings.Join f-n */
package main

import "fmt"

func main() {
	fmt.Println(Join(";", "1", "2", "3"))
}

func Join(sep string, a ...string) string {
	res := ""
	if len(a) == 0 {
		return ""
	}
	for _, v := range a[:len(a)-1] {
		res = res + v + sep
	}
	res = res + a[len(a) - 1]
	return res
}
