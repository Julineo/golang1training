// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"math"
	//"net/http"
	"log"
	"os"
	"math/big"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	
	zoom := big.NewInt(1)
	widthT := big.NewInt(1024) 
	heightT := big.NewInt(1024)
	width := new(big.Int)
	height := new(big.Int)
	width.Mul(widthT, zoom)
	height.Mul(heightT, zoom)
	widthP := new(big.Int)
	heightP := new(big.Int)
	two := big.NewInt(2)
	widthP.Mul(width, two)
	heightP.Mul(height, two)

	//creating multidimensional structure
	superSamples := make([][]color.Color, heightP.Int64())
    for i := range superSamples {
        superSamples[i] = make([]color.Color, widthP.Int64())
    }

	y := new(big.Float)
	x := new(big.Float)
	for py := big.NewInt(0); py.Int64() < heightP.Int64(); py.Add(py, py) {
		//pyB := new(big.Float).SetInt(py)
		//heightPB := new(big.Float).SetInt(heightP)
		y.Div(pyB, heightPB)// * (ymax - ymin) + ymin
		for px := 0; px < widthP; px++ {
			//pxB := new(big.Float).SetInt(px)
			//widthP := new(big.Float).SetInt(widthP)
		    x := pxB / widthP * (xmax - xmin) + xmin
		    z := complex(x, y)

		    superSamples[px][py] = mandelbrot(z)
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < width; py++  {
		for px := 0; px < height; px++ {
			sj, si := py * 2, px * 2
			// Averaging colors
			r1,g1,b1,_ := superSamples[si][sj].RGBA()
			r2,g2,b2,_ := superSamples[si+1][sj].RGBA()
			r3,g3,b3,_ := superSamples[si][sj+1].RGBA()
			r4,g4,b4,_ := superSamples[si+1][sj+1].RGBA()
			r := (r1+r2+r3+r4)/4
			g := (g1+g2+g3+g4)/4
			b := (b1+b2+b3+b4)/4
			ru := r >> 8
			gu := g >> 8
			bu := b >> 8
			c := color.RGBA{uint8(ru), uint8(gu), uint8(bu), 255}
			img.Set(px, py, c)
				
		}
	}
	
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	
	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}
	
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func mandelbrot(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			//return acosRGBA(v)
			//return acos(v)
			//return sqrt(v)
			//return newton(v)
			return newtonColor(v)
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:
func acosRGBA(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.RGBA{red, uint8(math.Abs(float64(red-blue))), blue, 255}
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

func newtonColor(z complex64) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(complex128(z*z*z*z-1)) < 1e-6 {
			return color.RGBA{uint8(real(z)*128) + 127,uint8(imag(z)*128) + 127,uint8(math.Abs(float64(real(z)+imag(z))))*128,255}
		}
	}
	return color.Black
}