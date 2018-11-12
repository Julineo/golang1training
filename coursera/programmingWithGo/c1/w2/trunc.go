/*
Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter float: ")
	str, _ := reader.ReadString('\n')

	_, err := strconv.ParseFloat(str[:len(str)-1], 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Split(str, ".")[0])
}


/*
//versions:


func main() {
	fmt.Print("enter a floating point number: ")
	var f float64
	_, err := fmt.Scan(&f)
	if err != nil {
		panic(err)
	}

	i := int64(f)
	fmt.Printf("truncated version: %v", i)
}
*/
