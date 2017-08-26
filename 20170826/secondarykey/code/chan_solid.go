package main

import (
	"fmt"
)

func main() {
	msgCh := make(chan string)
	go func() {
		msgCh <- "Shizuoka.go"
	}()

	msg := <-msgCh
	fmt.Println(msg)
}
