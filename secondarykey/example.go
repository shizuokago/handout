package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Logger struct {
	*http.ServeMux
}

func (h *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	h.ServeMux.ServeHTTP(w, r)
}

type JSONOutput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("menu.tmpl"))
	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html><body><b>Hello,HTML</b></body></html>`)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template.tmpl"))
	str := "Template"
	if err := t.Execute(w, str); err != nil {
		log.Fatal(err)
	}
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {

	o := &JSONOutput{
		Name:        "secondarykey",
		Description: "shizuoka.go",
	}

	res, _ := json.Marshal(o)
	w.Write(res)
}

func registerAssets() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "__session",
		Value: "Hello,cookie",
	}
	http.SetCookie(w, cookie)
	fmt.Fprintf(w, `Write cookie`)
}

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("__session")
	fmt.Fprintf(w, string(cookie.Value))
}

func agreeHandler(w http.ResponseWriter, r *http.Request) {

	url := r.URL.String()
	s := strings.Split(url, "/")

	user := ""
	agree := ""

	if len(s) >= 3 {
		agree = s[2]
	}

	if len(s) >= 4 {
		user = s[3]
	}

	fmt.Fprintf(w, "%s,%s", agree, user)

}

func main() {

	logger := Logger{http.NewServeMux()}

	logger.HandleFunc("/menu", menuHandler)

	logger.HandleFunc("/", indexHandler)
	logger.HandleFunc("/html", htmlHandler)
	logger.HandleFunc("/template", templateHandler)
	logger.HandleFunc("/json", jsonHandler)
	registerAssets()

	logger.HandleFunc("/setCookie", setCookieHandler)
	logger.HandleFunc("/cookie", cookieHandler)

	logger.HandleFunc("/agree/", agreeHandler)

	log.Fatal(http.ListenAndServe(":5550", &logger))
}
