package colorgradient

import (
	"image/color"
	"testing"
)

func TestNewGradient64(t *testing.T) {
	grad, err := NewGradient(
		color.RGBA64{0, 52942, 53713, 65535},
		color.RGBA64{0, 0, 0, 0},
		color.RGBA64{65535, 26985, 46260, 65535},
	)

	if err != nil {
		t.Error(err)
	}

	var want color.RGBA64
	want = color.RGBA64{0, 52942, 53713, 65535}
	if got := grad.At(0); got != want {
		t.Error("At(0) is incorrect, got:", got, "expected:", want)
	}
	want = color.RGBA64{0, 0, 0, 0}
	if got := grad.At(0.5); got != want {
		t.Error("At(0.5) is incorrect, got:", got, "expected:", want)
	}
	want = color.RGBA64{65535, 26985, 46260, 65535}
	if got := grad.At(1); got != want {
		t.Error("At(1) is incorrect, got:", got, "expected:", want)
	}
}

func TestNewGradient32(t *testing.T) {
	grad, err := NewGradient(
		color.RGBA{0, 206, 209, 255},
		color.RGBA{0, 0, 0, 0},
		color.RGBA{255, 105, 180, 255},
	)

	if err != nil {
		t.Error(err)
	}

	var want color.RGBA64
	want = color.RGBA64{0, 52942, 53713, 65535}
	if got := grad.At(0); got != want {
		t.Error("At(0) is incorrect, got:", got, "expected:", want)
	}
	want = color.RGBA64{0, 0, 0, 0}
	if got := grad.At(0.5); got != want {
		t.Error("At(0.5) is incorrect, got:", got, "expected:", want)
	}
	want = color.RGBA64{65535, 26985, 46260, 65535}
	if got := grad.At(1); got != want {
		t.Error("At(1) is incorrect, got:", got, "expected:", want)
	}
}

func TestNewGradientInvalidArgs(t *testing.T) {
	grad, err := NewGradient(
		color.RGBA{0, 206, 209, 255},
	)

	if err == nil {
		t.Error("Expected error, got:", grad)
	}
}
