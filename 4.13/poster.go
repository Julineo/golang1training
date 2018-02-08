package main

import (
	//"net/url"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	//"bufio"
	"strings"
)

var usage string = `usage:
xkcd [search term]
`
func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

type ResponseResultURL struct {
	Poster string
}

func main() {

	if len(os.Args) < 3 {
		usageDie()
	}
	
	apikey := os.Args[1]
	name := strings.Join(os.Args[2:]," ")
	
	fmt.Println("http://omdbapi.com/?t="+name+"&apikey="+apikey)
	
	resp, err := http.Get("http://www.omdbapi.com/?t="+name+"&"+apikey)
	if err != nil {
		fmt.Errorf("Error : %s", err)
	}
	
	fmt.Println(resp)

	var result ResponseResultURL
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Errorf("Error : %s", err)
	}
	resp.Body.Close()
	
	fmt.Println(result.Poster)
}
