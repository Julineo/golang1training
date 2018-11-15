package main

import (
	"fmt"
)

type animal interface {
	Eat() string
	Move() string
	Speak() string
}

type cow struct {
	food   string
	locomotion  string
	noise string
}

type bird struct {
	food   string
	locomotion  string
	noise string
}

type snake struct {
	food   string
	locomotion  string
	noise string
}

func (a *cow) Eat() string { return a.food }
func (a *cow) Move() string { return a.locomotion }
func (a *cow) Speak() string { return a.noise }

func (a *bird) Eat() string { return a.food }
func (a *bird) Move() string { return a.locomotion }
func (a *bird) Speak() string { return a.noise }

func (a *snake) Eat() string { return a.food }
func (a *snake) Move() string { return a.locomotion }
func (a *snake) Speak() string { return a.noise }

func main() {

	m := make(map[string]animal)
	var com, name, class string

	loop:
	for {
		fmt.Println(">")
		fmt.Scanf("%s %s %s", &com, &name, &class)

		switch com {
		case "newanimal":
			switch class {
			case "cow":
				m[name] = &cow{"grass", "walk", "moo"}
				fmt.Println("Created it!")
			case "bird":
				m[name] = &bird{"worms", "fly", "peep"}
				fmt.Println("Created it!")
			case "snake":
				m[name] = &snake{"mice", "slither", "hsss"}
				fmt.Println("Created it!")
			default:
				fmt.Println("wrong animal")
			}

		case "query":
			switch class {
			case "eat":
				fmt.Println(m[name].Eat())
			case "move":
				fmt.Println(m[name].Move())
			case "speak":
				fmt.Println(m[name].Speak())
			default:
				fmt.Println("wrong action")
			}

		case "x":
			break loop

		default:
			fmt.Println("wrong command!")
		}
	}
}
