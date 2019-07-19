package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type NeitherError string

func NewNeitherError(v string) *NeitherError {
	e := NeitherError(v)
	return &e
}

func (n NeitherError) Error() string {
	return string(n)
}

func main() {

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf("this command required argument filename\n")
		os.Exit(1)
	}

	for _, elm := range args {
		is, err := Is(elm)
		if err != nil {
			fmt.Printf("%v\n", err)
			if _, ok := err.(*NeitherError); !ok {
				os.Exit(1)
			}
			continue
		}
		fmt.Printf("%s is %t\n", elm, is)
	}

}

func Is(f string) (bool, error) {

	v, err := getValue(f)
	if err != nil {
		return false, fmt.Errorf("getValue() error: %v")
	}

	if exist(v, "true") {
		return true, nil
	}
	if exist(v, "false") {
		return false, nil
	}
	return false, NewNeitherError(fmt.Sprintf("neither true nor false[%s]", v))
}

func exist(target, v string) bool {
	if strings.Index(target, v) != -1 {
		return true
	}
	return false
}

func getValue(f string) (string, error) {

	fp, err := os.Open(f)
	if err != nil {
		return "", fmt.Errorf("os.Open() error[%s]: %v", f, err)
	}
	defer fp.Close()

	byt, err := ioutil.ReadAll(fp)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll() error[%s]: %v", f, err)
	}
	return string(byt), nil
}
