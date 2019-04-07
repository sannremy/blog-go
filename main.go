// import "$PROJECT_ROOT"

package main

import (
  "fmt"
  "net/http"

  "google.golang.org/appengine" // Required external App Engine library
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  // if statement redirects all invalid URLs to the root homepage.
  if r.URL.Path != "/" {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }

  fmt.Fprintln(w, "srchea-com auto deployed!")
}

func main() {
  http.HandleFunc("/", IndexHandler)
  appengine.Main() // Starts the server to receive requests
}
