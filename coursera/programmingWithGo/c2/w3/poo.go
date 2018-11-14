package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Anymal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Anymal) Eat() {

}

func (a *Anymal) Move() {

}

func (a *Anymal) Speak() {

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(">")
		scanner.Scan()
		v := scanner.Text()
		an := strings.Split(v, " ")[0]
		ac := strings.Split(v, " ")[1]
	}
}
