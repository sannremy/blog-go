package controllers

import (
	"net/http"
	"text/template"

	"github.com/srchea/homepage/internal/pkg/contexts"
	"github.com/srchea/homepage/internal/pkg/models"
)

type viewData struct {
	GlobalViewData *models.GlobalViewData
	PageTitle      string
}

// IndexHandler handles the main page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	staticFiles := r.Context().Value(contexts.StaticFilesKeyContextKey)

	tmpl := template.Must(template.ParseFiles("web/templates/layout.html"))
	data := &viewData{
		PageTitle: "src.onl deployed!",
		GlobalViewData: &models.GlobalViewData{
			StaticFiles: staticFiles,
		},
	}
	tmpl.Execute(w, data)
}
