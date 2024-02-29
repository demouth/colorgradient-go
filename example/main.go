package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/demouth/colorgradient-go"
)

func main() {
	linearGradient()
	radialGradient()
}

func linearGradient() {
	grad, err := colorgradient.
		NewGradient(
			color.RGBA{0, 206, 209, 255},
			color.RGBA{0, 0, 0, 0},
			color.RGBA{255, 105, 180, 255},
		)
	if err != nil {
		fmt.Println(err)
		return
	}

	const (
		width  = 300
		height = 50
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, grad.At(float64(x)/float64(width)))
		}
	}
	f, _ := os.Create("./linear-gradient.png")
	defer f.Close()
	png.Encode(f, img)
}

func radialGradient() {
	grad, err := colorgradient.
		NewGradient(
			color.RGBA{249, 206, 249, 255},
			color.RGBA{100, 105, 180, 255},
			color.RGBA{100, 105, 180, 0},
		)
	if err != nil {
		fmt.Println(err)
		return
	}

	const (
		width  = 300
		height = 300
	)
	fw := float64(width)
	fh := float64(height)
	cw := fw / 2
	ch := fh / 2
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			dx := float64(x) - cw
			dy := float64(y) - ch
			delta := math.Sqrt(dx*dx+dy*dy) / math.Min(cw, ch)
			col := grad.At(delta)
			img.Set(x, y, col)
		}
	}
	f, _ := os.Create("./radial-gradient.png")
	defer f.Close()
	png.Encode(f, img)
}
