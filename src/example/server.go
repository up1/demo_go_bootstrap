package example


import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")

	var results []Topic
	results = []Topic{
		Topic{1, "Title1"},
		Topic{2, "Title2"},
	}

	t.Execute(w, map[string]interface{}{
		"project_name": "MY DATA",
		"results":      results,
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

func StartServer() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/xxx", makeHandler(xxxHandler))
	http.HandleFunc("/", makeHandler(homeHandler))
	http.ListenAndServe(":8080", nil)
}