package mypkgs

import (
	"fmt"

	"mypkgs/sub"
)

func Hello() {
	fmt.Println("Hello mypkgs!")
	sub.Hello()
}
