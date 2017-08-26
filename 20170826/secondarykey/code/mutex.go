package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := new(sync.Mutex)
	for i := 0; i < 5; i += 1 {
		m.Lock()
		go func(i int) {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(i)
			m.Unlock()
		}(i)
	}

	fmt.Scanln()
}
