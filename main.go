package main

import (
  "html/template"
  "net/http"
  "net/url"
  "strings"
  "github.com/gorilla/mux"
  "google.golang.org/appengine" // Required external App Engine library
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

func RedirectWwwMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if strings.HasPrefix(r.Host, "www.") {
      host := strings.TrimLeft(r.Host, "www.")

      redirectedUrl := url.URL{
        Scheme: "https",
        Host: host,
        Path: r.URL.Path,
        RawQuery: r.URL.RawQuery,
      } 

      http.Redirect(w, r, redirectedUrl.String(), http.StatusMovedPermanently)
    }

    next.ServeHTTP(w, r)
	})
}

func main() {
  r := mux.NewRouter()

  // Serve assets
  fs := http.FileServer(http.Dir("web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

  // Redirect if www
  r.Use(RedirectWwwMiddleware)

  // Index route
  r.HandleFunc("/", IndexHandler)

  http.Handle("/", r)
  appengine.Main() // Starts the server to receive requests
}
