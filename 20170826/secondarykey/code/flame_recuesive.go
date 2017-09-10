package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"sync"
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

	//for {
	w.draw()
	w.Repaint()
	//}
}

func NewWindow(w, h int) *Window {
	win := Window{
		Width:  w,
		Height: h,
	}
	driver.Main(win.create)
	return &win
}

type point struct {
	x int
	y int
	v int
}

func (w *Window) draw() {

	rgba := w.buffer.RGBA()

	m := make([][]int, w.Width)
	ch := make(chan *point, 1024)

	fmt.Println(time.Now())

	for x := 0; x < w.Width; x++ {
		m[x] = make([]int, w.Height)
	}
	go w.sendColor(w.Height-1, m, ch)

	fmt.Println(time.Now())

	for elm := range ch {
		rgba.Set(elm.x, elm.y, color.RGBA{uint8(elm.v), 0, 0, 0})
	}

	fmt.Println(time.Now())
}

func (w *Window) sendColor(y int, m [][]int, ch chan *point) {

	wg := sync.WaitGroup{}

	for x := 0; x < w.Width; x++ {

		wg.Add(1)

		go func(x int) {

			r := 0
			if y == w.Height-1 {
				r = rand.Intn(256)
			} else {
				sum := 0
				num := 0

				sum += m[x][y+1]
				num++

				if x-1 >= 0 {
					sum += m[x-1][y+1]
					num++
				}

				if x+1 < w.Width {
					sum += m[x+1][y+1]
					num++
				}

				if y+2 < w.Height {
					sum += m[x][y+2]
					num++
				}
				r = sum / num
			}

			m[x][y] = r
			ch <- &point{
				x: x,
				y: y,
				v: r,
			}
			wg.Done()
		}(x)
	}

	wg.Wait()

	if y == 0 {
		close(ch)
		return
	}

	go w.sendColor(y-1, m, ch)

	return
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
