package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {

	urls := []string{
		"https://www.google.com/",
		"https://www.yahoo.co.jp/",
		"https://timeout.testurl/",
	}

	//START GROUP
	eg := errgroup.Group{}
	for _, elm := range urls {
		eg.Go(func() error {
			resp, err := http.Get(elm)
			if err != nil {
				return err
			}
			resp.Body.Close()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
	//END GROUP
}
