package controllers

import (
	"html/template"
	"net/http"
)

type IndexPageData struct {
	PageTitle string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html"))
	data := IndexPageData{
		PageTitle: "src.onl deployed!",
	}
	tmpl.Execute(w, data)
}
