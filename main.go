package main

import (
  "fmt"
  "strings"
  "net/http"
  "net/url"

  "google.golang.org/appengine" // Required external App Engine library
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  isUrlChanged := false
  targetUrl := url.URL{
    Scheme: r.URL.Scheme,
    Host: r.Host,
    Path: r.URL.Path,
    RawQuery: r.URL.RawQuery,
  } 

  // Force HTTPS
  if r.URL.Scheme == "http" {
    targetUrl.Scheme = "https"
    isUrlChanged = true
  }

  // Remove www.
  if strings.Contains(r.URL.Host, "www.") {
    targetUrl.Host = strings.TrimPrefix(r.URL.Host, "www.")
    isUrlChanged = true
  }

  if isUrlChanged {
    http.Redirect(w, r, targetUrl.String(), http.StatusFound)
    return
  }

  // if statement redirects all invalid URLs to the root homepage.
  if r.URL.Path != "/" {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }

  fmt.Fprintln(w, "src.onl auto deployed!")
}

func main() {
  http.HandleFunc("/", IndexHandler)
  appengine.Main() // Starts the server to receive requests
}
