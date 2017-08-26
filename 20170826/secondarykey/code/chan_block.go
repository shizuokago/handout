package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan string)

	select {
	case msg := <-msgCh:
		fmt.Println(msg)
	default:
		fmt.Println("non blocking")
	}
}
