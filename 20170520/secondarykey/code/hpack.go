package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/http2/hpack"
)

func main() {
//OMIT START
	s := "www.example.com"

	fmt.Println(len(s))
	fmt.Println(hpack.HuffmanEncodeLength(s))

	b := hpack.AppendHuffmanString(nil, s)
	fmt.Printf("%x\n", b)

	var buf bytes.Buffer
	hpack.HuffmanDecode(&buf, b)

	fmt.Printf("%s\n", buf.String())
//OMIT END
}
