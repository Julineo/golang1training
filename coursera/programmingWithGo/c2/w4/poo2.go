package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Anymal interface {
	Eat()
	Move()
	Speak()
}

type cow struct {
	eat string
	move string
	speak string
}

type bird struct {
	eat string
	move string
	speak string
}

type snake struct {
	eat string
	move string
	speak string
}

func (a *Anymal) Eat() {
	fmt.Println(a.food)
}

func (a *Anymal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Anymal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	cow := Anymal{"grass", "walk", "moo"}
	bird := Anymal{"worms", "fly", "peep"}
	snake := Anymal{"mice", "slither", "hsss"}

	for {
		fmt.Println(">")
		scanner.Scan()
		v := scanner.Text()

		if strings.Split(v, " ")[1] == "newanimal" {

			fmt.Println("Created it!")
		}

		if strings.Split(v, " ")[1] == "query" {

		}

		an := strings.Split(v, " ")[1]
		ac := strings.Split(v, " ")[2]



	}
}
