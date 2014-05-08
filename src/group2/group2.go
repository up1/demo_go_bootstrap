package group2

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type Topic struct {
		TopicID int
		Name    string
	}

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
