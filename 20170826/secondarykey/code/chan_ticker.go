package main

import (
	"fmt"
	"time"
)

func main() {

	msgCh := make(chan string)
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case msg := <-msgCh:
			fmt.Println(msg)
		case <-ticker.C:
			fmt.Println("secound!")
		}
	}
	ticker.Stop()
}
