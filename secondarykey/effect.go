package main

import (
	"fmt"
)

func main() {
	if trigger() != nil {
		fmt.Println("not nil")
	}
}

type OriginalError struct{}

func (e *OriginalError) Error() string {
	return "Error"
}

func newError() *OriginalError {
	return nil
}

func trigger() error {
	return newError()
}
