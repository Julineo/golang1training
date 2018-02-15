package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"net/url"
	"strings"
	"bytes"
	"os"
	"log"
)

// CreateIssue creates the GitHub issue.
//https://developer.github.com/v3/issues/#create-an-issue

func EditIssue(owner, repo, number string, fields map[string]string) (*Issue, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(fields)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	req, err := http.NewRequest("PATCH", url, buf)
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASS"))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to edit issue: %s", resp.Status)
	}
	var issue Issue
	if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil

}

func CreateIssue(owner string, repo string, number string) {

	_, err := EditIssue(owner, repo, number, map[string]string{"state": "open"})
	if err != nil {
		log.Fatal(err)
	}

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
