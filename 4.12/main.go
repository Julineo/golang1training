package main

import (
	//"net/url"
	"net/http"
	"fmt"
	"encoding/json"
)

const RequestURL = "https://xkcd.com/571/info.0.json"

type ResponseResult struct {
	Month string
	Num int
	Year string
}

func main() {

	//q := url.QueryEscape(strings.Join(terms, " "))
	//resp, err := http.Get(IssuesURL + "?q=" + q)
	resp, err := http.Get(RequestURL)
	if err != nil {
		fmt.Errorf("Error : %s", err)
	}
	
	var result ResponseResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Errorf("Error : %s", err)
	}
	resp.Body.Close()
	fmt.Println(result)

}
