package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"golang.org/x/net/html"
)

// CountWordsAndImages does an HTTP GET request for the HTML document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		//fmt.Println(n)
		images++
	}
	if n.Type == html.TextNode {
		//fmt.Println(n.Data)
		arrayWords := strings.Fields(n.Data)
		for _, n := range arrayWords {
			fmt.Printf("%d\n", n)
			words++
		}
	}

	if !(n.Data == "script" || n.Data == "style")  {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			w, i := countWordsAndImages(c)
			//words, images += w, i
			words, images = words+w, images+i
		}
	}
	return words, images
}

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
			continue
		}
		fmt.Println(url)
		fmt.Printf("words: %v, images %v\n", words, images)
	}
}
