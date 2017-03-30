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
		Title:   "Page Title",
		Content: "My Page content",
	}

	templates.ExecuteTemplate(w, "home", data)
}
