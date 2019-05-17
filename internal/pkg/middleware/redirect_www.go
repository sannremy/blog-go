package middleware

import (
	"net/http"
	"net/url"
	"strings"
)

func RedirectWwwMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.Host, "www.") {
			host := strings.TrimLeft(r.Host, "www.")

			redirectedUrl := url.URL{
				Scheme:   "https",
				Host:     host,
				Path:     r.URL.Path,
				RawQuery: r.URL.RawQuery,
			}

			http.Redirect(w, r, redirectedUrl.String(), http.StatusMovedPermanently)
		}

		next.ServeHTTP(w, r)
	})
}
