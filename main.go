package main

import (
  "io"
  "net/http"
  "github.com/gorilla/mux"
  "google.golang.org/appengine" // Required external App Engine library
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  io.WriteString(w, "src.onl auto deployed!\n")
}

func main() {
  r := mux.NewRouter()

  // Index route
  r.HandleFunc("/", IndexHandler)

  http.Handle("/", r)
  appengine.Main() // Starts the server to receive requests
}
