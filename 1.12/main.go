// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
package main

import (
	"log"
	"net/http"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"fmt"
	"net/url"
	//"strconv"
)
//https://golang.org/pkg/net/url/#Values
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.RawQuery)
    			m, _ := url.ParseQuery(r.URL.RawQuery)
    			fmt.Println(m)
			/*cycles, err := strconv.ParseFloat(m["cycles"][0], 64)
    			check(err)
			res, err := strconv.ParseFloat(m["res"][0], 64)
    			check(err)
			size, err := strconv.Atoi(m["size"][0])
    			check(err)
			nframes, err := strconv.Atoi(m["nframes"][0])
    			check(err)
			delay, err := strconv.Atoi(m["delay"][0])
    			check(err)*/
			lissajous(w/*, cycles, res, size, nframes, delay*/)
	}
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer/*, cycles int, res float64, size int, nframes int, delay int*/) {
	/*const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)*/
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}
//!-
