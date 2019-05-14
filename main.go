package main

import (
  "fmt"
  "strings"
  "net/http"
  "github.com/gorilla/mux"
  "google.golang.org/appengine" // Required external App Engine library
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "src.onl auto deployed!")
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", IndexHandler)
  http.Handle("/", r)

  if appengine.IsAppEngine() {
    go http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      host := r.Host
      if strings.Contains(host, "www.") {
        host = strings.TrimPrefix(host, "www.")
      }

      http.Redirect(w, r, "https://" + host + r.URL.String(), http.StatusMovedPermanently)
    }))
  }

  appengine.Main() // Starts the server to receive requests
}
