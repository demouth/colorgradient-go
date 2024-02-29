# colorgradient-go

colorgradient-go is a library for generating gradients in the Go language. It provides a simple API to create smooth gradients that transition between multiple colors.

## Features
- Create smooth gradients that transition between multiple colors
- Get the color at any position
- Simple API

## Usage

```go
package main

import (
	"image/color"
	"testing"

	"github.com/demouth/colorgradient-go"
)

func main() {
    grad, _ := colorgradient.NewGradient(
        color.RGBA64{0, 52942, 53713, 65535},
        color.RGBA64{0, 0, 0, 0},
        color.RGBA64{65535, 26985, 46260, 65535},
    )

    // Get colors at 0%, 50%, and 100%
    grad.At(0)   // color.RGBA64{0, 52942, 53713, 65535}
    grad.At(0.5) // color.RGBA64{0, 0, 0, 0}
    grad.At(1)   // color.RGBA64{65535, 26985, 46260, 65535}
}
```
