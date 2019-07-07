package app

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srchea/homepage/internal/pkg/controllers"
	"github.com/srchea/homepage/internal/pkg/libs"
	"github.com/srchea/homepage/internal/pkg/middleware"
)

// Start starts a HTTP server
func Start() {
	// Routes for posts
	libs.InitPosts()

	r := mux.NewRouter()

	// Serve assets
	fs := http.FileServer(http.Dir("dist/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Redirect if www
	r.Use(middleware.RedirectWwwMiddleware)

	// Get static js and css
	r.Use(middleware.StaticFilesMiddleware)

	// Routes
	r.HandleFunc("/", controllers.IndexHandler)
	r.HandleFunc("/about", controllers.AboutHandler)

	// Post routes
	slugsCount := len(libs.PostSlugs)
	for i := 0; i < slugsCount; i++ {
		r.HandleFunc("/"+libs.PostSlugs[i], controllers.PostHandler)
	}

	// Favicon
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/static/assets/favicon.ico")
	})

	// Robots.txt route
	r.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "User-agent: *")
	})

	// Serve
	http.Handle("/", r)
}
