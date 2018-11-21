/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	/*	var n int
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n)
	*/
	//initial slice
	sl := []int{10, 45, 23, 126, 432, 349, 2, 86, 23, 123, 865, 90, 3980, 21, 345, 15, 3, 789, 23, 0, 32, 34, 33, 133, 456}

	//4 slices
	sl1, sl2, sl3, sl4 := []int{}, []int{}, []int{}, []int{}

	//dividing slice to 4 parts:
	steps := len(sl) / 4
	counter := 0
	_ = counter

	for i, v := range sl {
		switch i / steps {
		case 0:
			sl1 = append(sl1, v)
		case 1:
			sl2 = append(sl2, v)
		case 2:
			sl3 = append(sl3, v)
		default:
			sl4 = append(sl4, v)
		}
	}
	fmt.Println(sl1)
	sortSl(sl1)
	fmt.Println(sl)
	fmt.Println(sl1)

}

func sortSl(sl []int) {
	sort.Ints(sl)
}
