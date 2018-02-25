//Write a function to print the contents of all text nodes in an HTML document tree. Do not descend into <script> or <style> elements, since their contents are not visible in a web browser.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "htmltext %v\n", err)
		os.Exit(1)
	}
	htmltext(doc)
}

// htmltext  prints text content from web page
func htmltext(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Printf("n.Data: %v\n", n.Data)
		//fmt.Printf("n: %v\n", n)
		//fmt.Printf("n.Type: %v\n", n.Type)
		//fmt.Printf("n.Data: %v\n", n.Data)
		//fmt.Printf("n.Attr: %v\n", n.Attr)
		//fmt.Printf("n.FirstChild: %v\n", n.FirstChild)
		//fmt.Printf("n.NextSibling: %v\n", n.NextSibling)
	}
	if !(n.Data == "script" || n.Data == "style")  {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			htmltext(c)
		}
	}
}

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
