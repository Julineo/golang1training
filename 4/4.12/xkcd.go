package main

import (
	"os"
	"fmt"
	"encoding/gob"
	"strings"
	"strconv"
	"net/http"
	"encoding/json"
)

var usage string = `usage:
xkcd [search term]
`
func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

type ResponseResultTranscript struct {
	Transcript string
}

func loadIdx(file string) (idx map[string][]int) {
        f, err := os.Open(file)
        if err != nil {
                panic("cant open file")
        }
        defer f.Close()

        enc := gob.NewDecoder(f)
        if err := enc.Decode(&idx); err != nil {
                panic("can't decode")
        }

        return idx
}

func loadTranscript(number int) {
	var result ResponseResultTranscript
	RequestURL := "https://xkcd.com/" + strconv.Itoa(number) + "/info.0.json"
		
	resp, err := http.Get(RequestURL)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
		
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Errorf("Error : %s", err)
	}
	resp.Body.Close()
	fmt.Println(result.Transcript, "\n")
}

func main() {
	var idx map[string][]int
	idx = loadIdx("idx")
	
	if len(os.Args) < 2 {
		usageDie()
	}
	args := strings.Join(os.Args[1:]," ")
	for _, number := range idx[args] {
		//printing link
		fmt.Println("https://xkcd.com/" + strconv.Itoa(number) + "/", "\n")
		//printing transcript
		loadTranscript(number)
	}
	
	//https://xkcd.com/1403/info.0.json
}