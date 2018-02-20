package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"os"
	"html/template"
)

var data []byte

var movieList = template.Must(template.New("movielist").Parse(`
<table>
<tr style='text-align: left'>
  <th>Title</th>
  <th>released</th>
  <th>color</th>
</tr>
{{range .Movies}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.released}}</td>
  <td>{{.color}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	data := `{"Movies":
	[
    		{
        	"Title": "Cool Hand Luke",
        	"released": 1967,
        	"color": true,
        	"Actors": [
            	"Paul Newman"
        			]
    		},
    		{
        	"Title": "Bullitt",
        	"released": 1968,
        	"color": true,
        	"Actors": [
            	"Steve McQueen",
            	"Jacqueline Bisset"
        			]
    		},
    		{
        	"Title": "Casablanca",
        	"released": 1942,
        	"Actors": [
            	"Humphrey Bogart",
            	"Ingrid Bergman"
        			]
    		}
	]
}`
	fmt.Printf("%s\n", data)

	if err := movieList.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

/*
{"movies":
	[
    		{
        	"Title": "Casablanca",
        	"released": 1942,
        	"Actors": [
            	"Humphrey Bogart",
            	"Ingrid Bergman"
        			]
    		},
    		{
        	"Title": "Cool Hand Luke",
        	"released": 1967,
        	"color": true,
        	"Actors": [
            	"Paul Newman"
        			]
    		},
    		{
        	"Title": "Bullitt",
        	"released": 1968,
        	"color": true,
        	"Actors": [
            	"Steve McQueen",
            	"Jacqueline Bisset"
        			]
    		}
	]
}
*/
