package main

import (
	//"net/url"
	"net/http"
	"fmt"
	"encoding/json"
	"encoding/gob"
	"strconv"
	"os"
	//"bufio"
	"strings"
)

//We'll use this const for figuring out the max number of comics.
const RequestLastURL = "https://xkcd.com/info.0.json"

type ResponseResultLast struct {
	Num int
}

type ResponseResult struct {
	Num int
	Transcript string
	Alt string
}

func saveIdx(file string, idx map[string][]int) {
f, err := os.Create(file)
        if err != nil {
                panic("cant open file")
        }
        defer f.Close()

        enc := gob.NewEncoder(f)
        if err := enc.Encode(idx); err != nil {
                panic("can't encode")
        }
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

func main() {

	resp, err := http.Get(RequestLastURL)
	if err != nil {
		fmt.Errorf("Error : %s", err)
	}
		
	var resultLast ResponseResultLast
	var result ResponseResult
	if err := json.NewDecoder(resp.Body).Decode(&resultLast); err != nil {
		resp.Body.Close()
		fmt.Errorf("Error : %s", err)
	}
	resp.Body.Close()

	var ResponseResults[]ResponseResult
	//Loop through all comics. Read the data we need.
	//for i := 1; i <= 100; i++ {//for testing purposes
	for i := 1; i <= resultLast.Num; i++ {
		RequestURL := "https://xkcd.com/" + strconv.FormatInt(int64(i), 10) + "/info.0.json"
		
		resp, err := http.Get(RequestURL)
		if err != nil {
			fmt.Errorf("Error: %s", err)
		}
		
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			fmt.Errorf("Error : %s", err)
		}
		resp.Body.Close()
		ResponseResults = append(ResponseResults, result)
		
		fmt.Println(RequestURL)
	}
	
	//Build the index.
	idx := make(map[string][]int)
	for _, result := range ResponseResults {
		for _, word := range strings.Split(result.Alt," ") {
			idx[word] = append(idx[word], result.Num)
		}
	}
	for _, result := range ResponseResults {
		for _, word := range strings.Split(result.Transcript," ") {
			idx[word] = append(idx[word], result.Num)
		}
	}
	
	
	//Save to a file
	saveIdx("idx", idx)
	//var idx2 map[string][]int
	//idx2 = loadIdx("idx")
}
