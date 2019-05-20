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

var js = ""
var css = ""

// StaticFilesMiddleware populate request context with main js and css file names
func StaticFilesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if js == "" || css == "" {
			distDir := "dist/"
			filePrefix := distDir + "main-"

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

		staticFiles := make(map[string]string)
		staticFiles["JS"] = js
		staticFiles["CSS"] = css

		ctx := context.WithValue(r.Context(), contexts.StaticFilesKeyContextKey, staticFiles)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
