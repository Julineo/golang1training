//5.17 Write a variadic function ElementsByTagName that, given an HTML node tree and zero or more names, returns all the elements that match one of those names.

package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://www.golang.com"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check Content-Type is HTML (e.g., "text/html; charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	defer resp.Body.Close()

	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		fmt.Errorf("%s has type %s, not text/html", url, ct)
		return
}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}

	titles := ElementsByTagName(doc, "title")
	others := ElementsByTagName(doc, "a", "body", "span")

	//showing results
	for _, n := range titles {
		fmt.Println(n.Data)
	}
	for _, n := range others {
		fmt.Println(n.Data)
	}
}

// Returns html elements with matching names
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	fmt.Println(name)
	res := []*html.Node{}

	visitNode := func(n *html.Node) {
		for _, v := range name {
			if n.Type == html.ElementNode && n.FirstChild != nil && v == n.Data {
				res = append(res, n)
			}
		}
	}

	forEachNode(doc, visitNode, nil)

	return res
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
