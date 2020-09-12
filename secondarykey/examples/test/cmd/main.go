package main

import (
	"fmt"
	"os"

	"mypkgs"
)

func main() {

	if !mypkgs.Deny() {
		fmt.Println("Access Deny!")
		os.Exit(1)
	}
	fmt.Println("Success")
}
