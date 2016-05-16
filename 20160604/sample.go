package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	println("Hello console!")
	js.Global.Get("document").Call("write", "Hello world!")
}
