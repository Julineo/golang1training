// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"os"
	"bufio"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	//writing to a file
	f, err := os.OpenFile("D:/Projects/GoProjects/src/golang1training/3.1/yourfile", os.O_APPEND|os.O_WRONLY, 0600)//try to open file
	if err != nil {
		f, err = os.Create("D:/Projects/GoProjects/src/golang1training/3.1/yourfile")//create new file if can't open
	}
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	_, err = fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	check(err)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			if math.IsNaN(ax) ||  math.IsInf(ax, 0) {continue} 
			if math.IsNaN(ay) ||  math.IsInf(ay, 0) {continue} //skip polygon if it has non-finite value
			bx, by := corner(i, j)
			if math.IsNaN(bx) ||  math.IsInf(bx, 0) {continue} 
			if math.IsNaN(by) ||  math.IsInf(by, 0) {continue} //skip polygon if it has non-finite value
			cx, cy := corner(i, j+1)
			if math.IsNaN(cx) ||  math.IsInf(cx, 0) {continue} 
			if math.IsNaN(cy) ||  math.IsInf(cy, 0) {continue} //skip polygon if it has non-finite value
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(dx) ||  math.IsInf(dx, 0) {continue} 
			if math.IsNaN(dy) ||  math.IsInf(dy, 0) {continue} //skip polygon if it has non-finite value

			_, err = fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			check(err)
		}
	}
	_, err = fmt.Fprintf(w, "</svg>")
    	check(err)
    	w.Flush()
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}
//!-
