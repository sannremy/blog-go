package controllers

import (
	"net/http"
	"text/template"

	"github.com/srchea/homepage/internal/pkg/contexts"
	"github.com/srchea/homepage/internal/pkg/models"
)

// Get template
var tmpl = template.Must(template.ParseFiles(
	"web/templates/layout.html",

	"web/templates/partials/icons.html",
	"web/templates/partials/navbar.html",

	"web/templates/pages/posts.html",
	"web/templates/pages/about.html",
))

// Private view data
type viewData struct {
	GlobalViewData *models.GlobalViewData
	PageView       string
}

// IndexHandler handles the / page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Get static files from context
	staticFiles := r.Context().Value(contexts.StaticFilesKeyContextKey)

	// View data
	data := &viewData{
		PageView: "posts",
		GlobalViewData: &models.GlobalViewData{
			StaticFiles: staticFiles,
		},
	}

	// Execute view data + template
	tmpl.Execute(w, data)
}

// AboutHandler handles the /about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Get static files from context
	staticFiles := r.Context().Value(contexts.StaticFilesKeyContextKey)

	// View data
	data := &viewData{
		PageView: "about",
		GlobalViewData: &models.GlobalViewData{
			StaticFiles: staticFiles,
		},
	}

	// Execute view data + template
	tmpl.Execute(w, data)
}
