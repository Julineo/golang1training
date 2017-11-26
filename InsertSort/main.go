package main

import (
	"fmt"
	"gopl.io/ch4/treesort"
	//"math/rand"
)

func main() {
	//data := make([]int, 50)
	//for i := range data {
	//	data[i] = rand.Int() % 50
	//}
	data :=[]int{5,8,3,0,10,19,8}
	fmt.Println(data)
	treesort.Sort(data)
	fmt.Println(data)
}
