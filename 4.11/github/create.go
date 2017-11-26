package github

import (
	//"encoding/json"
	"fmt"
	//"net/http"
	//"net/url"
	//"strings"
)

// CreateIssue creates the GitHub issue.
//https://developer.github.com/v3/issues/#create-an-issue
func CreateIssue(terms []string) {

	fmt.Println("CreateIssue will be here")
	//q := url.QueryEscape(strings.Join(terms, " "))
	//resp, err := http.Get(IssuesURL + "?q=" + q)
	//if err != nil {
	//	return nil, err
	//}
	//if resp.StatusCode != http.StatusOK {
	//	resp.Body.Close()
	//	return nil, fmt.Errorf("search query failed: %s", resp.Status)
	//}

	//var result IssuesSearchResult
	//if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	//	resp.Body.Close()
	//	return nil, err
	//}
	//resp.Body.Close()
	//return &result, nil
}
