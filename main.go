package main

import (
	"github.com/srchea/homepage/internal/app"
	"google.golang.org/appengine" // Required external App Engine library
)

func main() {
	app.Start()
	appengine.Main() // Starts the server to receive requests
}
