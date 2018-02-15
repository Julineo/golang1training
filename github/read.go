package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetIssue(owner string, repo string, number string) (*Issue, error) {
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("read query failed: %s: %s", url, resp.Status)
	}
	
	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		fmt.Println(issue)
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &issue, nil
}