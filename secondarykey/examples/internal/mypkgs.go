package mypkgs

import (
	"net/http"

	"mypkgs/handler"
	"mypkgs/handler/admin"
	//"mypkgs/other"
)

func Start() error {

	http.HandleFunc("/admin/", admin.IndexHandler)
	//http.HandleFunc("/other", other.IndexHandler)
	http.HandleFunc("/", handler.IndexHandler)

	addr := ":8080"
	fmt.Printf("Listen[%s]\n", addr)
	return http.ListenAndServe(addr, nil)
}
