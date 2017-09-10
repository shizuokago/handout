package main

import (
	"fmt"
)

func main() {

	msgCh := make(chan string, 3)

	msgCh <- "GoCon"
	msgCh <- "GCPUG Shonan"
	msgCh <- "Shizuoka.go"

	fmt.Println(<-msgCh)

	msgCh <- "Release Party Shizuoka.go"

	close(msgCh)

	for msg := range msgCh {
		fmt.Println(msg)
	}
}
