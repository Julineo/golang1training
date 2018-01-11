package main

import (
	"net/url"
	"net/http"
)

const RequestURL = "https://xkcd.com/571/info.0.json"

func main() {

	//q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

}