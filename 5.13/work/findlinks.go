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
	"path/filepath"

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
	fmt.Println(urla)
	fmt.Println(path.Base(urla))

	download(urla)


	list, err := links.Extract(urla)
	if err != nil {
		log.Print(err)
	}
	return list
}

var origHost string

// Download url to filename
func download(rawurl string) error {
	url, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Errorf("bad url: %s", err)
	}
	if origHost == "" {
		origHost = url.Host
	}
	if origHost != url.Host {
		return nil
	}
	dir := url.Host
	var filename string
	if filepath.Ext(filename) == "" {
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(url.Path))
		filename = url.Path
	}
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}
	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	// Check for delayed write errors section 5.8.
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
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
