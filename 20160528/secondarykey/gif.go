package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"os"

	"github.com/lazywei/go-opencv/opencv"
)

func init() {
}

func getCapture(f string) (*opencv.Capture, int) {
	c := opencv.NewFileCapture(f)
	if c == nil {
		return nil, 0
	}
	frames := int(c.GetProperty(opencv.CV_CAP_PROP_FRAME_COUNT))
	return c, frames
}

func createGIF(f int) *gif.GIF {
	g := &gif.GIF{}
	g.Image = make([]*image.Paletted, f)
	g.Delay = make([]int, f)
	return g
}

func getImage(c *opencv.Capture) (image.Image, error) {
	i := c.QueryFrame()
	if i == nil {
		return nil, fmt.Errorf("QueryFrame() is nil")
	}
	return i.ToImage(), nil
}

func convertGIF(i image.Image) (*image.Paletted, error) {
	b := new(bytes.Buffer)
	if err := gif.Encode(b, i, nil); err != nil {
		return nil, fmt.Errorf("GIF encoding error:%v", err)
	}

	p, err := gif.Decode(b)
	if err != nil {
		return nil, fmt.Errorf("GIF decoding error:%v", err)
	}
	return p.(*image.Paletted), nil
}

func writeGIF(o string, g *gif.GIF) error {
	f, err := os.OpenFile(o, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("Open File error:%v", err)
	}

	defer f.Close()
	err = gif.EncodeAll(f, g)
	if err != nil {
		return fmt.Errorf("Create Animation GIF error:%v", err)
	}
	return nil
}

func main() {

	f := "matrix.mp4"
	o := "matrix.gif"

	err := run(f, o)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
}

func run(i, o string) error {

	c, fm := getCapture(i)
	if c == nil {
		return fmt.Errorf("can not open video")
	}
	defer c.Release()

	fmt.Printf("フレ ーム数:%d\n", fm)

	imgs := make([]image.Image, fm)
	for idx := 0; idx < fm; idx++ {
		img, err := getImage(c)
		if err != nil {
			return fmt.Errorf("Error of getImage:%v", err)
		}
		imgs[idx] = img
	}

	g := createGIF(fm)

	for idx, img := range imgs {
		in, err := convertGIF(img)
		if err != nil {
			return fmt.Errorf("convertGIF error:%v", err)
		}
		g.Image[idx] = in
		g.Delay[idx] = 0
	}

	err := writeGIF(o, g)
	if err != nil {
		return fmt.Errorf("Write Error:%v", err)
	}
	return nil
}
