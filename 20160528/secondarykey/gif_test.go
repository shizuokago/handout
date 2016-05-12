package main

import (
	"image"
	"image/gif"
	"os"
	"testing"

	"github.com/lazywei/go-opencv/opencv"
)

var gc *opencv.Capture
var gfm int
var gi image.Image
var gg *gif.GIF

func TestMain(m *testing.M) {

	gc, gfm = getCapture("matrix.mp4")

	c, fm := getCapture("matrix.mp4")
	defer c.Release()

	imgs := make([]image.Image, fm)
	for idx := 0; idx < fm; idx++ {
		img, _ := getImage(c)
		imgs[idx] = img
	}

	gg = createGIF(fm)

	for idx, img := range imgs {
		if idx == 0 {
			gi = img
		}

		in, _ := convertGIF(img)
		gg.Image[idx] = in
		gg.Delay[idx] = 0
	}

	code := m.Run()
	os.Exit(code)
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run("matrix.mp4", "matrix.gif")
	}
}

func BenchmarkGetCapture(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c, _ := getCapture("matrix.mp4")
		c.Release()
	}
}

func BenchmarkCreateGIF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createGIF(gfm)
	}
}

func BenchmarkGetImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i, err := getImage(gc)
		if err != nil {
			break
		}
		if i == nil {
			break
		}
	}
}

func BenchmarkConvertGIF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertGIF(gi)
	}
}

func BenchmarkWriteGIF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		writeGIF("test.gif", gg)
	}
}
