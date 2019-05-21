package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/srchea/homepage/internal/pkg/contexts"
)

// Buffer for JS file path
var js = ""

// Buffer for CSS file path
var css = ""

// StaticFilesMiddleware populate request context with main js and css file names
func StaticFilesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Check if js and css files have been already fetched
		if js == "" || css == "" {
			distDir := "dist/"
			filePrefix := distDir + "main-"

			if os.Getenv("ENVIRONMENT") == "development" {
				filePrefix += "dev"
			}

			// Find main js and css files
			err := filepath.Walk(distDir, func(path string, info os.FileInfo, err error) error {
				if strings.HasPrefix(path, filePrefix) {
					if strings.HasSuffix(path, ".js") {
						js = strings.TrimLeft(path, distDir)
					}

					if strings.HasSuffix(path, ".css") {
						css = strings.TrimLeft(path, distDir)
					}
				}

				return nil
			})

			if err != nil {
				log.Fatal(err)
			}
		}

		// Add files in the static files map
		staticFiles := make(map[string]string)
		staticFiles["JS"] = js
		staticFiles["CSS"] = css

		// Add to context
		ctx := context.WithValue(r.Context(), contexts.StaticFilesKeyContextKey, staticFiles)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
