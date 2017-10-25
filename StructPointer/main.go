package main

import (
	"time"
	"fmt"
)

	type Employee struct {
		ID		int
		Name		string
		Address		string
		DoB		time.Time
		Position	string
		Salary		int
		ManagerID	int
	}

func main() {
	var dilbert Employee

	dilbert.Salary -= 5000

	fmt.Println(dilbert.Salary)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	fmt.Println(dilbert.Position)
	fmt.Println(employeeOfTheMonth.Position)
}
