/*5.13: Modify crawl to make local copies of the pages it finds, creating directories as necessary. Ignore external links*/

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"net/http"
	"net/url"
	"io"
	"strings"

	"gopl.io/ch5/links"
)

//for test purpose. to limit working time
var ext int

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				ext++
				if ext > 200 {
					return
				}

				worklist = append(worklist, f(item)...)
			}
		}
	}
}


//modified crawl for saving data on local disk
var local []string
func crawl(urla string) []string {
	//fmt.Println(urla)
	//fmt.Println(path.Base(urla))
	u, err := url.Parse(urla)
	if err != nil {
		log.Fatal(err)
	}

	//check if local page
	flag := false
	for _, v := range local {
		if strings.Replace(u.Host, "www.", "", 1) == v {
			flag = true
		}
	}

	if ! flag {
		return []string{}
	}
	fmt.Println(urla, path.Base(urla))
	fmt.Println(path.Base(urla))
	//download(urla, path.Base(urla))


	list, err := links.Extract(urla)
	if err != nil {
		log.Print(err)
	}
	return list
}

//download url to filename
func download(url, filename string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(f, resp.Body)
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	// saving only local pages, ignoring eternal links
	for _, v := range os.Args[1:] {
		local = append(local, strings.Replace(path.Base(v), "www.", "", 1))
	}
	fmt.Println(local)
	breadthFirst(crawl, os.Args[1:])
}
