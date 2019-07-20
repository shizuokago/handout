package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"errors"
)

type NeitherError string

func NewNeitherError(v string) error {
	n := NeitherError(v)
	return &n
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

	var nErr *NeitherError
	var pErr *os.PathError

	for _, elm := range args {
		is, err := Is(elm)

		if err != nil {

			fmt.Printf("%v\n", err)

			if errors.As(err, &nErr) {
			} else if errors.As(err, &pErr) {
				err = createFile(elm)
				if err != nil {
					fmt.Printf("createFile(%s)", elm)
					os.Exit(1)
				}
				fmt.Printf("create %s\n", elm)
			} else {
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
		return false, fmt.Errorf("getValue(): %w", err)
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
		return "", fmt.Errorf("os.Open(): %w", err)
	}
	defer fp.Close()

	byt, err := ioutil.ReadAll(fp)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll(): %w", err)
	}
	return string(byt), nil
}

func createFile(f string) error {

	fp, err := os.Create(f)
	if err != nil {
		return fmt.Errorf("os.Create() error: %w", err)
	}
	defer fp.Close()
	return nil
}
