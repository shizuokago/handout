package main

import (
	"net/http"
)

func main() {
	handler := http.FileServer(http.Dir("./"))
	go http.ListenAndServeTLS("localhost:55555", "cert.pem", "key.pem", handler)
	http.ListenAndServe(":5555", handler)
}
