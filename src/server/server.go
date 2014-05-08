package main

import (
	"group1"
	"group2"
	"net/http"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/xxx", makeHandler(group1.XxxHandler))
	http.HandleFunc("/", makeHandler(group2.HomeHandler))
	http.ListenAndServe(":8080", nil)
}
