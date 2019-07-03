package controllers

import (
	"net/http"
	"text/template"

	"github.com/srchea/homepage/internal/pkg/contexts"
	"github.com/srchea/homepage/internal/pkg/models"
)

// Private view data for index
type viewData struct {
	GlobalViewData *models.GlobalViewData
	PageTitle      string
}

// IndexHandler handles the main page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Get static files from context
	staticFiles := r.Context().Value(contexts.StaticFilesKeyContextKey)

	// Get template
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/partials/icons.html",
		"web/templates/partials/navbar.html",
	))

	// View data
	data := &viewData{
		PageTitle: "src.onl deployed!",
		GlobalViewData: &models.GlobalViewData{
			StaticFiles: staticFiles,
		},
	}

	// Execute view data + template
	tmpl.Execute(w, data)
}
