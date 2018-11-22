/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"fmt"
	"sort"
	"sync"
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

	//sorting 4 slices
	var wg sync.WaitGroup

	sortSl := func(sl []int) {
		sort.Ints(sl)
		fmt.Println(sl)
		wg.Done()
	}

	wg.Add(4)
	go sortSl(sl1)
	go sortSl(sl2)
	go sortSl(sl3)
	go sortSl(sl4)

	wg.Wait()
	fmt.Println(sl)

	//result slice
	fmt.Println(mergeSl(mergeSl(sl1, sl2), mergeSl(sl3, sl4)))
}

func mergeSl(sl1, sl2 []int) []int {

	i, j := 0, 0
	res := []int{}

	for i < len(sl1) && j < len(sl2) {
		if sl1[i] > sl2[j] {
			res = append(res, sl2[j])
			j++
		} else {
			res = append(res, sl1[i])
			i++
		}
	}

	if i == len(sl1) {
		for j < len(sl2) {
			res = append(res, sl2[j])
			j++
		}
	} else {
		for i < len(sl1) {
			res = append(res, sl1[i])
			i++
		}
	}
	return res
}

/*
inplace version

package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "sort"
    "sync"
    "bytes"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Printf("Please enter integers space separated on one line\nints: ")
    scanner.Scan()
    ints := strings.Fields(scanner.Text())

    isli := make([]int, 0, 1)

    for i:=0;i<len(ints);i++ {
        n,e := strconv.Atoi(ints[i])
        if e != nil {
            fmt.Printf("Error parsing input: %s\n", e)
            return
        }
        isli = append(isli, n)
    }

    chunk := 4
    delim := len(isli) / chunk
    if len(isli) % chunk > 0 {delim++}

    var wg sync.WaitGroup
    var buf bytes.Buffer
    wg.Add(chunk)


    // Loop through and concurrently sort sections of the array
    for i,x:=1,0; x<len(isli); x,i = x+delim, i+1 {
        lb := x
        rb := x + delim
        if rb >= len(isli) {rb -= rb-len(isli)}
        // Start a goroutine...
        go csort(isli[lb:rb], &buf, &wg, i)
    }

    wg.Wait()
    sort.Ints(isli)

    fmt.Print(buf.String())
    fmt.Print("Final Sorted array: ")
    fmt.Println(isli)

}

func csort(isli []int, buf *bytes.Buffer, wg *sync.WaitGroup, i int) {
    var lbuf bytes.Buffer
    //lbuf.WriteString(fmt.Sprintf("loop %d before: ", i))
    //lbuf.WriteString(fmt.Sprintln(isli))

    sort.Ints(isli)

    //lbuf.WriteString(fmt.Sprintf("loop %d after: ", i))
    lbuf.WriteString(fmt.Sprint("Interim sorted array: "))
    lbuf.WriteString(fmt.Sprintln(isli))
    buf.WriteString(lbuf.String())
    wg.Done()
}
*/
