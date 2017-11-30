//Build a tool that lets users create, read, update, and delete GitHub issues from the command line, invoking their preferred text editor when substantial text input is required.
package main

import (
	"fmt"
	"log"
	"os"

	"golang1training/4.11/github"
	//"gopl.ex/4.11/github"
)

func main() {
	fmt.Println(os.Args[1:])
	mode := os.Args[1]
	//create functionality
	if mode == "create" {
		fmt.Println("Create")
		github.CreateIssue(os.Args[1:])
	}
	
	//search functionality
	if mode == "search" {
		result, err := github.SearchIssues(os.Args[1:])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n %s\n Time:%v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		
		}
	}
}

