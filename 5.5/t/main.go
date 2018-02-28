package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// visit appends to links each link found in n, and returns the result.
func visit(n *html.Node) (images, words int){
	if n.Type == html.ElementNode && n.Data == "img" {
		//fmt.Println(n)
		images++
	}
	if n.Type == html.TextNode {
		words++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		i, w := visit(c)
		images, words = images + i, words + w
	}
	return images, words
}

func main() {
	for _, url := range os.Args[1:] {
		images, words, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		fmt.Printf("Images: %v\n", images)
		fmt.Printf("Words: %v\n", words)
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) (int, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return 0, 0, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return 0, 0, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	t1, t2 := visit(doc)
	return t1, t2, nil
}
