package main

import (
	"fmt"
)

func main() {
		fmt.Printf("%v", ys(0,2,8,6,-5))
}
	
func ys(x0, y0, x1, y1, x float64) float64 {
	
	return y0 + (x - x0)/(x1 - x0) * (y1 - y0)

}
