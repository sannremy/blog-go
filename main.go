package main

import (
  "io"
  "net/http"
  "net/url"
  "strings"
  "github.com/gorilla/mux"
  "google.golang.org/appengine" // Required external App Engine library
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  io.WriteString(w, "src.onl auto deployed!\n")
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

  // Redirect if www
  r.Use(RedirectWwwMiddleware)

  // Index route
  r.HandleFunc("/", IndexHandler)

  http.Handle("/", r)
  appengine.Main() // Starts the server to receive requests
}
