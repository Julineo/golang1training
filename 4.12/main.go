package main

import (
	//"net/url"
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
)

//We'll use this const for figuring out the max number of comics.
const RequestLastURL = "https://xkcd.com/info.0.json"

type ResponseResultLast struct {
	Num int
}

type ResponseResult struct {
	Month string
	Num int
	Year string
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
	//Loop through all comics
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
	fmt.Println(ResponseResults)

}
