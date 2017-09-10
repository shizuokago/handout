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

	for {
		w.draw()
		w.Repaint()
	}
}

type point struct {
	x int
	y int
	v int
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

	//START BUFFER
	ch := make(chan *point, 1024)
	//END BUFFER
	wg := sync.WaitGroup{}

	fmt.Println(time.Now())

	go func() {

		for y := hiy - 1; y >= loy; y-- {

			for x := lox; x < hix; x++ {

				wg.Add(1)

				go func(x int) {

					r := 0
					if y == hiy-1 {
						m[x] = make([]int, hiy)
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

					ch <- &point{
						x: x,
						y: y,
						v: r,
					}

					m[x][y] = r

					wg.Done()
				}(x)

			}

			wg.Wait()
		}
		close(ch)
	}()

	fmt.Println(time.Now())

	//START RANGE
	for p := range ch {
		rgba.Set(p.x, p.y, color.RGBA{uint8(p.v), 0, 0, 0})
	}
	//END RANGE

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
