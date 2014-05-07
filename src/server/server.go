package main

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, map[string]interface{}{
		"project_name": "MY DATA",
	})
}

func xxxHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, map[string]interface{}{
		"project_name": "XXXX",
	})
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/xxx", makeHandler(xxxHandler))
	http.HandleFunc("/", makeHandler(homeHandler))
	http.ListenAndServe(":8080", nil)
}
