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

	w.draw()

	/*
		for {
			w.draw()
			w.Repaint()
		}
	*/
}

func NewWindow(w, h int) *Window {
	win := Window{
		Width:  w,
		Height: h,
	}
	driver.Main(win.create)
	return &win
}

type colorData struct {
	x int
	y int

	sum int
	n   int
}

func (w *Window) draw() {

	rgba := w.buffer.RGBA()
	b := rgba.Bounds()
	lox := b.Min.X
	loy := b.Min.Y
	hix := b.Max.X
	hiy := b.Max.Y

	m := make(map[int]map[int]*colorData, 5000)

	mtx := new(sync.Mutex)
	ch := make(chan *colorData)
	for x := lox; x <= hix; x++ {
		m[x] = make(map[int]*colorData)
	}

	fmt.Println(time.Now())

	for y := hiy; y >= loy; y-- {
		for x := lox; x <= hix; x++ {

			go func(x, y int) {

				var data *colorData

				mtx.Lock()
				if y == hiy {
					r := rand.Intn(256)
					data = &colorData{
						x:   x,
						y:   y,
						n:   1,
						sum: r,
					}
				} else {
					data = m[x][y]
				}
				ch <- data

				if y != loy {

					if m[x][y-1] == nil {
						m[x][y-1] = &colorData{
							x:   x,
							y:   y - 1,
							sum: data.sum,
							n:   1,
						}
					} else {
						m[x][y-1].sum += data.sum
						m[x][y-1].n++
					}

					if x != lox {
						if m[x-1][y-1] == nil {
							m[x-1][y-1] = &colorData{
								x:   x - 1,
								y:   y - 1,
								sum: data.sum,
								n:   1,
							}
						} else {
							m[x-1][y-1].sum += data.sum
							m[x-1][y-1].n++
						}
					}

					if x != hix {
						if m[x+1][y-1] == nil {
							m[x+1][y-1] = &colorData{
								x:   x + 1,
								y:   y - 1,
								sum: data.sum,
								n:   1,
							}
						} else {
							m[x+1][y-1].sum += data.sum
							m[x+1][y-1].n++
						}
					}
				}

				if y-2 >= loy {
					if m[x][y-2] == nil {
						m[x][y-2] = &colorData{
							x:   x,
							y:   y - 2,
							sum: data.sum,
							n:   1,
						}
					} else {
						m[x][y-2].sum += data.sum
						m[x][y-2].n++
					}
				}

				mtx.Unlock()

			}(x, y)

		}
	}

	fmt.Println(time.Now())

	idx := 1
	for o := range ch {
		r := o.sum / o.n
		rgba.Set(o.x, o.y, color.RGBA{uint8(r), 0, 0, 0})
		idx++
		if idx == 1024*512 {
			close(ch)
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
