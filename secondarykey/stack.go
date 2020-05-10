package main

import (
	"fmt"
)

func main() {
	var s Stack
	s.Push("test1")
	s.Push("test2")

	s.Push([]byte("test3"))

	obj := s.Peek()
	fmt.Println(obj)
}

type Stack []interface{}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Peek() interface{} {
	rtn := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return rtn
}
