package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type handler time.Duration

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I can only live for %d seconds.", h)
}

func main() {

	t := time.Duration(10)
	h := handler(t)

	server := &http.Server{Addr: ":8080", Handler: h}
	go func() {
		server.ListenAndServe()
	}()

	c := time.Tick(t * time.Second)
	for range c {
		server.Shutdown(context.TODO())
		os.Exit(1)
	}
}
