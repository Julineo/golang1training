package reader

import (
	"bytes"
	"testing"
	"fmt"

	"golang.org/x/net/html"
)

func TestNewReader(t *testing.T) {
	s := "hi there"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(NewReader(s))
	if n != int64(len(s)) || err != nil {
		t.Logf("n=%d err=%s", n, err)
		t.Fail()
	}
	if b.String() != s {
		t.Logf(`"%s" != "%s"`, b.String(), s)
	}
}

func TestNewReaderWithHTML(t *testing.T) {
	s := "<html><body><p>Hello, NewReader!</p></body></html>"
	doc, err := html.Parse(NewReader(s))// html.Parse returns doc with nodes in it
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			fmt.Println(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
