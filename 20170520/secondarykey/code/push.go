package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	http.Handle("/index_files/", http.StripPrefix("/index_files/", http.FileServer(http.Dir("index_files"))))

	//START OMIT
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if pusher, ok := w.(http.Pusher); ok {
			if err := pusher.Push("/index_files/bg_2048.jpg", nil); err != nil {
				log.Printf("Error Push: %v", err)
			}
		}

		fp, _ := os.OpenFile("index.html", os.O_RDONLY, 0644)
		data, _ := ioutil.ReadAll(fp)
		fmt.Fprintf(w, string(data))
	})

	log.Fatal(http.ListenAndServeTLS("localhost:55555", "cert.pem", "key.pem", nil))
	//END OMIT
}
