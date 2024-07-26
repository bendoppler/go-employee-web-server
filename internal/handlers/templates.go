package handlers

import (
	"html/template"
	"net/http"
)

var templates = template.Must(
	template.ParseFiles(
		"web/template/employees.html", "web/template/view.html", "web/template/edit.html",
	),
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
