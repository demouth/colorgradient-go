package colorgradient

import (
	"errors"
	"image/color"
)

type Gradient struct {
	colors []color.Color
	pos    []float64
}

func NewGradient(colors ...color.Color) (Gradient, error) {
	if len(colors) < 2 {
		return Gradient{}, errors.New("Invalid Args")
	}
	pos := linspace(0, 1, uint(len(colors)))
	return Gradient{
		colors: colors,
		pos:    pos,
	}, nil
}

// copy from https://github.com/mazznoer/colorgrad/blob/cd89eb0d684d7015ca39dc48e2ad5f4a951f0f60/util.go#L11-L22
func linspace(min, max float64, n uint) []float64 {
	if n == 1 {
		return []float64{min}
	}
	d := max - min
	l := float64(n) - 1
	res := make([]float64, n)
	for i := range res {
		res[i] = (min + (float64(i)*d)/l)
	}
	return res
}

func (lg Gradient) At(t float64) color.Color {
	if t <= 0 {
		t = 0
	} else if t >= 1 {
		t = 1
	}

	low := 0
	high := len(lg.pos)

	for {
		if low >= high {
			break
		}
		mid := (low + high) / 2
		if lg.pos[mid] < t {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low == 0 {
		low = 1
	}

	p1 := lg.pos[low-1]
	p2 := lg.pos[low]
	t = (t - p1) / (p2 - p1)
	x1, y1, z1, a1 := lg.colors[low-1].RGBA()
	x2, y2, z2, a2 := lg.colors[low].RGBA()
	r := linearInterpolate(float64(x1), float64(x2), t)
	g := linearInterpolate(float64(y1), float64(y2), t)
	b := linearInterpolate(float64(z1), float64(z2), t)
	a := linearInterpolate(float64(a1), float64(a2), t)
	r *= a / 65535
	g *= a / 65535
	b *= a / 65535
	r16 := uint16(r)
	g16 := uint16(g)
	b16 := uint16(b)
	a16 := uint16(a)

	return color.RGBA64{R: r16, G: g16, B: b16, A: a16}
}

func linearInterpolate(a, b float64, t float64) float64 {
	f := a + t*(b-a)
	return f
}
