package main

import (
	"fmt"
	"log"
	"os"
	"html/template"
	//"golang1training/github"
	"gopl.ex/github"
	//"net/http"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
  <th>Milestone</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		fmt.Printf("%+v\n", item)
	}
	/*
	handler := func(w http.ResponseWriter, r *http.Request) {
		//m, err := url.ParseQuery(r.URL.RawQuery)
		//fill := m["fill"][0]
	
		if err := issueList.Execute(w, result); err != nil {
			log.Fatal(err)
		}
	}
	
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))*/
}
