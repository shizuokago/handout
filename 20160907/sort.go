package main

import (
	"fmt"
	"sort"
)

func main() {
	hoge := Hoge{1, 4, 2, 0, 6, 9, 3, 5, 8, 7}
	sort.Sort(hoge)
	fmt.Println(hoge)
}

type Hoge []int

func (h Hoge) Len() int {
	return len(h)
}

func (h Hoge) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Hoge) Swap(i, j int) {
	h[j], h[i] = h[i], h[j]
	return
}
