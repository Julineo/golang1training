// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"net/http"
	"log"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.1      // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "image/svg+xml")
		_, err := fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
		check(err)
		//calculating z max and z min
		var zmax, zmin, z float64 = 0, -0, 0
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				_, _, z = corner(+i, j)
				if math.IsNaN(z) ||  math.IsInf(z, 0) {continue}
				if z > zmax {
					zmax = z
				}
				if z < zmin {
					zmin = z
				}
			}
		}
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, z := corner(i+1, j)
				if math.IsNaN(ax) ||  math.IsInf(ax, 0) {continue} 
				if math.IsNaN(ay) ||  math.IsInf(ay, 0) {continue} //skip polygon if it has non-finite value
				bx, by, _ := corner(i, j)
				if math.IsNaN(bx) ||  math.IsInf(bx, 0) {continue} 
				if math.IsNaN(by) ||  math.IsInf(by, 0) {continue} //skip polygon if it has non-finite value
				cx, cy, _ := corner(i, j+1)
				if math.IsNaN(cx) ||  math.IsInf(cx, 0) {continue} 
				if math.IsNaN(cy) ||  math.IsInf(cy, 0) {continue} //skip polygon if it has non-finite value
				dx, dy, _ := corner(i+1, j+1)
				if math.IsNaN(dx) ||  math.IsInf(dx, 0) {continue} 
				if math.IsNaN(dy) ||  math.IsInf(dy, 0) {continue} //skip polygon if it has non-finite value

				_, err = fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color(z, zmax, zmin))
				check(err)
			}
		}
		_, err = fmt.Fprintf(w, "</svg>")
	    	check(err)

	}//handler := func
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
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

//extrapolation function
func ys(x0, y0, x1, y1, x float64) float64 {
	return y0 + (x - x0)/(x1 - x0) * (y1 - y0)
}

//function to convert z coordinate into color from #ff0000 to #0000ff
func color(z, zmax, zmin float64) string {
	blue := 255
	red := 255
	if z > 0 {
		red = int(ys(0,255,zmax,0, z))
		//if red < 0 {red = 0}
		//if red > 255 {red = 255}
	} else {
		blue = int(ys(zmin,0,0,255, z))
		//if blue < 0 {blue = 0}
	}
	return "#" + string(fmt.Sprintf("%02x", blue)) + "00" + string(fmt.Sprintf("%02x", red))
}
