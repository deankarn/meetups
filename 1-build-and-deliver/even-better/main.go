package main

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func main() {

	templates, _ = template.ParseGlob("*.tmpl")

	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Title   string
		Content string
	}{
		Title:   "Page New Title",
		Content: "My Page content Even Better!",
	}

	templates.ExecuteTemplate(w, "home", data)
}
