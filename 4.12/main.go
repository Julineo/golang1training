package main

import (
	//"net/url"
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
	"log"
	"os"
	"bufio"
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

func main() {

	//q := url.QueryEscape(strings.Join(terms, " "))
	//resp, err := http.Get(IssuesURL + "?q=" + q)
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
	
	//Write data to json format
	data, err := json.MarshalIndent(ResponseResults, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	
	//writing to a file
	f, err := os.OpenFile("data", os.O_APPEND|os.O_WRONLY, 0600)//try to open file
	if err != nil {
		f, err = os.Create("data")//create new file if can't open
	}
	if err != nil {
		fmt.Errorf("Error : %s", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	_, err = fmt.Fprint(w, data)
	if err != nil {
		fmt.Errorf("Error : %s", err)
	}
	w.Flush()

}
