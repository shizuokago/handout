package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"os"
	"runtime"

	"github.com/lazywei/go-opencv/opencv"
)

func init() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
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

func convertGIF(i image.Image, idx int) palettedIdx {

	p := palettedIdx{
		pal: nil,
		idx: idx,
		err: nil,
	}

	r := new(bytes.Buffer)
	err := gif.Encode(r, i, nil)
	if err != nil {
		p.err = fmt.Errorf("GIF encoding error:%v", err)
		return p
	}

	in, err := gif.Decode(r)
	if err != nil {
		p.err = fmt.Errorf("GIF decoding error:%v", err)
		return p
	}

	p.pal = in.(*image.Paletted)
	return p
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

type bufferIdx struct {
	buf *bytes.Buffer
	idx int
	err error
}

type palettedIdx struct {
	pal *image.Paletted
	idx int
	err error
}

func createBuffer(img image.Image, idx int) *bufferIdx {

	b := bufferIdx{
		buf: new(bytes.Buffer),
		idx: idx,
		err: nil,
	}

	err := gif.Encode(b.buf, img, nil)
	if err != nil {
		b.err = fmt.Errorf("GIF encoding error:%v", err)
	}
	return &b
}

func createPaletted(b *bufferIdx) *palettedIdx {

	p := palettedIdx{
		pal: nil,
		idx: b.idx,
		err: nil,
	}

	pal, err := gif.Decode(b.buf)
	if err != nil {
		p.err = fmt.Errorf("GIF decoding error:%v", err)
	}
	p.pal = pal.(*image.Paletted)
	return &p
}

func run(i, o string) error {

	c, fm := getCapture(i)
	if c == nil {
		return fmt.Errorf("can not open video")
	}
	defer c.Release()

	fmt.Printf("フレーム数:%d\n", fm)

	imgs := make([]image.Image, fm)
	for idx := 0; idx < fm; idx++ {
		img, err := getImage(c)
		if err != nil {
			return fmt.Errorf("Error of getImage:%v", err)
		}
		imgs[idx] = img
	}

	g := createGIF(fm)

	bufCh := make(chan *bufferIdx)
	palCh := make(chan *palettedIdx)

	for idx, img := range imgs {
		go func(ig image.Image, i int) {
			bufCh <- createBuffer(ig, i)
		}(img, idx)
	}

	cnt := 0
L:
	for {
		select {
		case buf := <-bufCh:
			go func(b *bufferIdx) {
				palCh <- createPaletted(b)
			}(buf)
		case pal := <-palCh:
			g.Image[pal.idx] = pal.pal
			g.Delay[pal.idx] = 0
			cnt++
			if cnt == fm {
				break L
			}
		}
	}

	err := writeGIF(o, g)
	if err != nil {
		return fmt.Errorf("Write Error:%v", err)
	}
	return nil
}
