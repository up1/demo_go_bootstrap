package group1

import (
	"html/template"
	"net/http"
)

func XxxHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/group1/index.html")
	t.Execute(w, map[string]interface{}{
		"project_name": "For Group 1",
	})
}
