package middleware

import (
	"net/http"
	"net/url"
	"strings"
)

// RedirectWwwMiddleware removes www. from the URL
func RedirectWwwMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if URL has www.
		if strings.HasPrefix(r.Host, "www.") {
			// Remove www.
			host := strings.TrimLeft(r.Host, "www.")

			// Build the new URL
			redirectedURL := url.URL{
				Scheme:   "https",
				Host:     host,
				Path:     r.URL.Path,
				RawQuery: r.URL.RawQuery,
			}

			// Redirect
			http.Redirect(w, r, redirectedURL.String(), http.StatusMovedPermanently)
		}

		next.ServeHTTP(w, r)
	})
}
