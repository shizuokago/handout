package main

import (
	"fmt"
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

	m := make([][]int, hix)

	fmt.Println(time.Now())

	for y := hiy - 1; y >= loy; y-- {

		for x := lox; x < hix; x++ {

			if y == hiy-1 {
				m[x] = make([]int, hiy)
			}

			r := 0
			if y == hiy-1 {
				r = rand.Intn(256)
			} else {
				sum := 0
				idx := 1

				sum += m[x][y+1]
				if x-1 >= lox {
					sum += m[x-1][y+1]
					idx++
				}

				if x+1 < hix {
					sum += m[x+1][y+1]
					idx++
				}

				if y+2 < hiy {
					sum += m[x][y+2]
					idx++
				}
				r = sum / idx
			}

			m[x][y] = r
		}
	}

	fmt.Println(time.Now())

	for x := lox; x < hix; x++ {
		for y := loy; y < hiy; y++ {
			ran := m[x][y]
			rgba.Set(x, y, color.RGBA{uint8(ran), 0, 0, 0})
		}
	}

	fmt.Println(time.Now())
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
	w.window.Send(paint.Event{})
}

func (w *Window) Release() {
	w.window.Release()
	w.buffer.Release()
}
