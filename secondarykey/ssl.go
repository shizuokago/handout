package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, HTTPS!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
}
