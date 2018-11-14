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
		an := strings.Split(v, " ")[0]
		ac := strings.Split(v, " ")[1]

		if an == "cow" && ac == "eat" {
			cow.Eat()
		}
		if an == "cow" && ac == "move" {
			cow.Move()
		}
		if an == "cow" && ac == "speak" {
			cow.Speak()
		}

		if an == "bird" && ac == "eat" {
			bird.Eat()
		}
		if an == "bird" && ac == "move" {
			bird.Move()
		}
		if an == "bird" && ac == "speak" {
			bird.Speak()
		}

		if an == "snake" && ac == "eat" {
			snake.Eat()
		}
		if an == "snake" && ac == "move" {
			snake.Move()
		}
		if an == "snake" && ac == "speak" {
			snake.Speak()
		}

	}
}
