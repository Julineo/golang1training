// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode prints html outline in pretty form
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {

/*	if n.Type == html.TextNode {
		fmt.Println("Start:",n.Data,":End")
	}*/

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

var depth int

//modyfied version of startElement
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		last := "/>"
		if n.FirstChild != nil {
			last = ">"
		}
		attrs := make([]string, 0, len(n.Attr))
		for _, a := range n.Attr {
			attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
		}
		attrStr := ""
		if len(n.Attr) != 0 {
			attrStr = " " + strings.Join(attrs, " ")
		}
		fmt.Printf("%*s<%s%s%s\n", depth*2, "", n.Data, attrStr, last)
		if n.FirstChild != nil {
			depth++
		}
}
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if len(text) != 0 {
			fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		}
	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

//modyfied version of endElement
func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
