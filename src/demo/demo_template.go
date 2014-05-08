package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "Up1")
	if err != nil {

	}
}
