package main

import (
	"fmt"
	"image"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

type Window struct {
	Width  int
	Height int

	window screen.Window
	buffer screen.Buffer
}

func main() {

	w := NewWindow(1024, 512)
	defer w.Release()

	fmt.Scanln()

}

func NewWindow(w, h int) *Window {
	win := Window{
		Width:  w,
		Height: h,
	}
	driver.Main(win.create)
	return &win
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

	//START BUFFER
	winSize := image.Point{w.Width, w.Height}
	b, err := s.NewBuffer(winSize)
	if err != nil {
		log.Fatal("Error:Create Window")
	}
	w.buffer = b
	//END BUFFER

	return
}

func (w *Window) Release() {
	w.window.Release()
	w.buffer.Release()
}
