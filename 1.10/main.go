// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"bufio"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	//writing to a file
	f, err := os.OpenFile("/tmp/yourfile", os.O_APPEND|os.O_WRONLY, 0600)//try to open file
	if err != nil {
    		f, err = os.Create("/tmp/yourfile")//create new file if can't open
	}
    	check(err)
    	defer f.Close()
	w := bufio.NewWriter(f)


	for range os.Args[1:] {
		 _, err = fmt.Fprintf(w, "%s\n", <-ch) // receive from channel ch and write to file
    		check(err)
	}
	_, err = fmt.Fprintf(w, "%s\n", "%.2fs elapsed\n", time.Since(start).Seconds())
    	check(err)
    	w.Flush()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

//!-
