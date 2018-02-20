package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"html/template"
	"strings"
)

var data []byte

var movieList = template.Must(template.New("movielist").Parse(`
<table>
<tr style='text-align: left'>
  <th>Title</th>
  <th>Actor</th>
</tr>
{{range .Movies}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Actor.Name}}</td>
</tr>
{{end}}
</table>
`))

type DataS struct {
	Movies      []*Movie
}

type Movie struct {
	Title	string
	Actor      *Actor
}

type Actor struct {
	Name   string
}

func main() {
	data := `{"Movies":
	[
			{
				"Title": "Cool Hand Luke",
				"Actor": {
							"Name": "Paul Newman"
						}
			},
    		{
				"Title": "Bullitt",
				"Actor": {
							"Name":"Steve McQueen"
						}
    		},
    		{
				"Title": "Casablanca"
    		}
	]
}`

/*
,
				"Actor": {
							"Name":"Humphrey Bogart"
						}
*/

	fmt.Printf("%s\n", data)
	
	r := strings.NewReader(data)
	
	var result DataS
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		fmt.Println(1)
		log.Fatal(err)
	}
	
	fmt.Printf("%v\n", result)

	if err := movieList.Execute(os.Stdout, result); err != nil {
		fmt.Println(2)
		log.Fatal(err)
	}
}
