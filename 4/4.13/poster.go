package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	"strings"
	"io"
	"log"
)

var usage string = `usage:
poster [search term]
`
func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

type ResponseResultURL struct {
	Poster string
	Number int
}

func main() {

	if len(os.Args) < 3 {
		usageDie()
	}
	
	apikey := os.Args[1]
	name := strings.Join(os.Args[2:]," ")
	
	resp, err := http.Get("http://omdbapi.com/?t="+name+"&apikey="+apikey)
	//resp, err := http.Get("http://api.open-notify.org/astros.json")
	if err != nil {
		log.Fatal(err)
	}
	
	//fmt.Printf("resp %v \n",resp)

	var result ResponseResultURL
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}
	resp.Body.Close()
	
	resp, err = http.Get(result.Poster)
	if err != nil {
		log.Fatal(err)
	}
	//open file for writing
	image, err := os.Create("Posters/" + name + ".jpg")
	if err != nil {
		log.Fatal(err)
	}
	//io.Copy is good for a huge files
	size, err := io.Copy(image, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	resp.Body.Close()
	image.Close()
	fmt.Printf("%s with %v bytes downloaded\n", name, size)
}
