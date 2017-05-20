package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.FileServer(http.Dir("./"))
	log.Fatal(http.ListenAndServeTLS("localhost:55555", "cert.pem", "key.pem", handler))

}
