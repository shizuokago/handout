package main

import (
	"fmt"
	"time"
)

func main() {

	msgCh := make(chan string)

	go worker(msgCh)

	select {
	case msg := <-msgCh:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
	}
}

func worker(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "woker done!"
}
