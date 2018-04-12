// getByID prints the node by its ID.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		getByID(url)
	}
}

func getByID(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", ElementByID(doc, "playgroundButton"))

	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, startTrav, endTrav, id)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, id string) bool, id string) *html.Node {

/*	if n.Type == html.TextNode {
		fmt.Println("Start:",n.Data,":End")
	}*/

	if pre != nil {
		if !pre(n, id)	{
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, id)
	}

	if post != nil {
		post(n, id)
	}
	return nil
}

var depth int

func startTrav(n *html.Node, id string) bool {
	for _, a := range n.Attr {
		if a.Key == "id" {
			return false
		}
	}
	return true
}

func endTrav(n *html.Node, id string) bool {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return true
}
