package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	templ := template.Must(template.ParseFiles("html/index.html"))
	templ.Execute(w, nil)
}

func faviconHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/favicon.ico")
}

func htmxHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/htmx.min.js")
}

var count = 0

func clickHandler(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, fmt.Sprintf("%d", count))
	count += 1
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/htmx.min.js", htmxHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/click", clickHandler)

	log.Printf("About to listen on 8000. Go to http://127.0.0.1:8000/")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
