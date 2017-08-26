package main

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"

	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
)

type Window struct {
	Width  int
	Height int

	window screen.Window
	buffer screen.Buffer
}

func main() {

	rand.Seed(time.Now().UnixNano())

	w := NewWindow(1024, 512)
	defer w.Release()

	//START GOROUTINE
	go func() {
		for {
			switch e := w.window.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case paint.Event:
				w.window.Upload(image.Point{}, w.buffer, w.buffer.Bounds())
				w.window.Publish()
			}
		}
	}()
	//END GOROUTINE

	for {
		w.draw()
		w.Repaint()
	}
}

func NewWindow(w, h int) *Window {
	win := Window{
		Width:  w,
		Height: h,
	}
	driver.Main(win.create)
	return &win
}

func (w *Window) draw() {

	rgba := w.buffer.RGBA()
	b := rgba.Bounds()
	lox := b.Min.X
	loy := b.Min.Y
	hix := b.Max.X
	hiy := b.Max.Y

	for x := lox; x <= hix; x++ {
		for y := loy; y <= hiy; y++ {
			r := rand.Intn(256)
			g := rand.Intn(256)
			b := rand.Intn(256)
			rgba.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}
	//END DRAW
}

func (w *Window) create(s screen.Screen) {

	opt := &screen.NewWindowOptions{
		Title:  "Display",
		Width:  w.Width,
		Height: w.Height,
	}

	win, err := s.NewWindow(opt)
	if err != nil {
		log.Fatal("Error:Create Window")
	}
	w.window = win

	winSize := image.Point{w.Width, w.Height}

	b, err := s.NewBuffer(winSize)
	if err != nil {
		log.Fatal("Error:Create Window")
	}
	w.buffer = b

	return
}

func (w *Window) Repaint() {
	//START REPAINT
	w.window.Send(paint.Event{})
	//END REPAINT
}

func (w *Window) Release() {
	w.window.Release()
	w.buffer.Release()
}
