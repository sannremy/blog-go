package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/srchea/homepage/internal/pkg/contexts"
	"github.com/srchea/homepage/internal/pkg/libs"
	"github.com/srchea/homepage/internal/pkg/models"
	"golang.org/x/text/message"
)

// Get template
var tmpl = template.Must(template.ParseFiles(
	"web/templates/layout.html",

	"web/templates/partials/icons.html",
	"web/templates/partials/navbar.html",

	"web/templates/pages/posts.html",
	"web/templates/pages/post.html",
	"web/templates/pages/about.html",

	"web/templates/pages/error.html",
))

// Private view data
type indexViewData struct {
	GlobalViewData *models.GlobalViewData
	PageView       string
	PageTitle      string
	Posts          *postsTemplateViewData
}

// Private view data
type aboutViewData struct {
	GlobalViewData *models.GlobalViewData
	PageView       string
	PageTitle      string
}

// Private view data
type postViewData struct {
	GlobalViewData *models.GlobalViewData
	PageView       string
	PageTitle      string
	Post           *postTemplateViewData
}

// View for page-posts template
type postsTemplateViewData struct {
	PostTitles map[string]string
	PostDates  map[string]time.Time
	PostSlugs  []string
}

// View for page-post template
type postTemplateViewData struct {
	PostTitle     string
	PostDate      time.Time
	PostSlug      string
	PostHTML      string
	PostViewCount string
}

// Private view data
type errorViewData struct {
	GlobalViewData *models.GlobalViewData
	PageView       string
	PageTitle      string
}

// IndexHandler handles the / page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Get static files from context
	staticFiles := r.Context().Value(contexts.StaticFilesKeyContextKey)

	// View data
	data := &indexViewData{
		Posts: &postsTemplateViewData{
			PostTitles: libs.PostTitles,
			PostDates:  libs.PostDates,
			PostSlugs:  libs.PostSlugs,
		},
		PageView:  "posts",
		PageTitle: "Blog posts",
		GlobalViewData: &models.GlobalViewData{
			StaticFiles: staticFiles,
		},
	}

	// Execute view data + template
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("Template execution failed: %s", err)
	}
}

// AboutHandler handles the /about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Get static files from context
	staticFiles := r.Context().Value(contexts.StaticFilesKeyContextKey)

	// View data
	data := &aboutViewData{
		PageView:  "about",
		PageTitle: "About",
		GlobalViewData: &models.GlobalViewData{
			StaticFiles: staticFiles,
		},
	}

	// Execute view data + template
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("Template execution failed: %s", err)
	}
}

// PostHandler handles post pages
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// l10n
	printer := message.NewPrinter(message.MatchLanguage("en"))

	// Get static files from context
	staticFiles := r.Context().Value(contexts.StaticFilesKeyContextKey)

	slug := strings.TrimPrefix(r.URL.Path, "/")
	title := libs.PostTitles[slug]
	date := libs.PostDates[slug]
	html := libs.PostHTMLs[slug]
	viewCount := printer.Sprintf("%d", libs.PostViewCounts[slug])

	// View data
	data := &postViewData{
		PageView:  "post",
		PageTitle: title,
		Post: &postTemplateViewData{
			PostTitle:     title,
			PostDate:      date,
			PostSlug:      slug,
			PostHTML:      html,
			PostViewCount: viewCount,
		},
		GlobalViewData: &models.GlobalViewData{
			StaticFiles: staticFiles,
		},
	}

	// Execute view data + template
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("Template execution failed: %s", err)
	}

	libs.IncrementPostViewCount(slug)
}

// RedirectNotFoundHandler redirects to 404
func RedirectNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("/404"), http.StatusPermanentRedirect)
}

// ErrorNotFoundHandler handles 404 pages
func ErrorNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Get static files from context
	staticFiles := r.Context().Value(contexts.StaticFilesKeyContextKey)

	// View data
	data := &errorViewData{
		PageView:  "error",
		PageTitle: "Page not found",
		GlobalViewData: &models.GlobalViewData{
			StaticFiles: staticFiles,
		},
	}

	// 404 status
	w.WriteHeader(http.StatusNotFound)

	// Execute view data + template
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("Template execution failed: %s", err)
	}
}
