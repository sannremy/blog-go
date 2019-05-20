package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srchea/homepage/internal/pkg/controllers"
	"github.com/srchea/homepage/internal/pkg/middleware"
)

func Start() {
	r := mux.NewRouter()

	// Serve assets
	fs := http.FileServer(http.Dir("dist/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Redirect if www
	r.Use(middleware.RedirectWwwMiddleware)

	// Get static js and css
	r.Use(middleware.StaticFilesMiddleware)

	// Index route
	r.HandleFunc("/", controllers.IndexHandler)

	http.Handle("/", r)
}
