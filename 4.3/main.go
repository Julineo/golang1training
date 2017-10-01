package main

import (
	"fmt"
)

const MAX int = 6

func main() {

	t := []int{0, 1, 2, 3, 4, 5}

	var ptr [MAX]*int;

	for  i := 0; i < MAX; i++ {
		ptr[i] = &t[i] /* assign the address of integer. */
	}

	fmt.Println(ptr)
	reverse(&ptr)
	for  i := 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i,*ptr[i] )
	} // "[5 4 3 2 1 0]"
}





func reverse(s *[MAX]*int) {
	fmt.Println(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}

