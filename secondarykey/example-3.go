package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	NeitherErr = errors.New("neither true nor false")
)

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
			if err != NeitherErr {
				os.Exit(1)
			}
			continue
		}
		fmt.Printf("%s is %t\n", elm, is)
	}

}

func Is(f string) (bool, error) {
	fp, err := os.Open(f)
	if err != nil {
		return false, err
	}
	defer fp.Close()

	byt, err := ioutil.ReadAll(fp)
	if err != nil {
		return false, err
	}
	v := string(byt)

	if strings.Index(v, "true") != -1 {
		return true, nil
	}

	if strings.Index(v, "false") != -1 {
		return false, nil
	}

	return false, NeitherErr
}
