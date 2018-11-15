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
	eat   string
	move  string
	speak string
}

type bird struct {
	eat   string
	move  string
	speak string
}

type snake struct {
	eat   string
	move  string
	speak string
}

func (a *cow) Eat() {
	fmt.Println(a.eat)
}

func (a *cow) Move() {
	fmt.Println(a.move)
}

func (a *cow) Speak() {
	fmt.Println(a.speak)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	cow := []string{"grass", "walk", "moo"}
	bird := []string{"worms", "fly", "peep"}
	snake := []string{"mice", "slither", "hsss"}
	m := make(map[string]string)

	for {
		fmt.Println(">")
		scanner.Scan()
		v := scanner.Text()
		//map for storing name and type of anymal

		if strings.Split(v, " ")[0] == "newanymal" {
			an := strings.Split(v, " ")[1]
			at := strings.Split(v, " ")[2]
			m[an] = at

			fmt.Println("Created it!")
			fmt.Println(m)
		}

		if strings.Split(v, " ")[0] == "query" {
			an := strings.Split(v, " ")[1]
			atype := m[an]
			aq := strings.Split(v, " ")[2]
			if atype == "cow" {
				if aq == "eat" {
					fmt.Println(cow[0])
				}
				if aq == "move" {
					fmt.Println(cow[1])
				}
				if aq == "speak" {
					fmt.Println(cow[2])
				}
			}

			if atype == "bird" {
				if aq == "eat" {
					fmt.Println(bird[0])
				}
				if aq == "move" {
					fmt.Println(bird[1])
				}
				if aq == "speak" {
					fmt.Println(bird[2])
				}
			}

			if atype == "snake" {
				if aq == "eat" {
					fmt.Println(snake[0])
				}
				if aq == "move" {
					fmt.Println(snake[1])
				}
				if aq == "speak" {
					fmt.Println(snake[2])
				}
			}
		}

	}
}
