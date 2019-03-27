package main

import (
	"fmt"
	"github.com/shizuokago/golin"
)

func main() {
	err := golin.Print()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
