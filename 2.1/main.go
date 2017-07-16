//testing package usage
package main

import (
	"fmt"
	"gopl.ex/2.1/tempconv"
)

func main() {
	fmt.Printf("Cold! %v\n", tempconv.BoilingK)
	fmt.Printf("Kelvin! %v\n",tempconv.CToK(0))

}
