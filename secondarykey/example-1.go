package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf("this command required argument filename\n")
		os.Exit(1)
	}

	for _, elm := range args {
		//USE IS FUNCTION START
		is, err := Is(elm)
		if err != nil {
			fmt.Printf("%+v\n", err)
		} else {
			fmt.Printf("%s is %t\n", elm, is)
		}
		//USE IS FUNCTION END
	}

}

//ISFUNCTION START
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
	return false, nil
}

//ISFUNCTION END
