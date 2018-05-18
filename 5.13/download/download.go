package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

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
	pUrl := flag.String("url", "http://www.alexfast.org", "URL to be processed")
	flag.Parse()
	url := *pUrl
	if url == "" {
		fmt.Fprintf(os.Stderr, "Error: empty URL!\n")
		return
	}

	filename :=path.Base(url)
	fmt.Println("test filename: ", filename)
	fmt.Println("Checking if " + filename + " exists ...")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		download(url, filename)
		fmt.Println(filename, " saved!")
	} else {
		fmt.Println(filename, " error ", err)
	}
}
