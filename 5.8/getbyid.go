// getByID prints the node by its ID.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: getbyid id url url url...")
	}


	id := os.Args[1]

	for _, url := range os.Args[2:] {
		getNodeByID(url, id)
	}
}

func getNodeByID(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Result: %v\n", ElementByID(doc, id))
	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {

	var relmnt *html.Node

	startTrav := func(n *html.Node) bool {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				relmnt = n
				return false
			}
		}
		return true
	}

	forEachNode(doc, startTrav, nil)

	return relmnt
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {

	if pre != nil {
		if !pre(n)	{
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	//we dont need this actually
	if post != nil {
		fmt.Println("post !=nil")
	}

}
