package urlshort

import (
	"net/http"
)

// MapHandler is a function that's called whenever the server receives a request.
// The MapHandler is a closure; it has access to request details (http.Request).
// It extracts the URL from this request and checks if it's in the map; if it is,
// it redirects to corresponding URL; if not, it will fallback to fallback handler.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Extract URL from request
		requestedURL := r.URL.Path

		// Check if requestedURL in map
		if url, ok := pathsToUrls[requestedURL]; ok {
			http.Redirect(w, r, url, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}
