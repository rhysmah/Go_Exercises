package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

type PathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

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

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedData, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	mappedData := buildURLMap(parsedData)

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := mappedData[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}, nil
}

// Helper function: unmarshals YAML data into a struct slice
func parseYAML(yml []byte) ([]PathURL, error) {
	var yamlData []PathURL
	err := yaml.Unmarshal(yml, &yamlData)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal YAML data: %v", err)
	}
	return yamlData, nil
}

// Helpre function: builds a map of paths and URLs from a struct slice
func buildURLMap(parsedYAML []PathURL) map[string]string {
	pathsToURLS := make(map[string]string)

	// Iterate over the slice of PathURL and populate map
	// Ignore the index; unpack the struct data
	for _, pathURL := range parsedYAML {
		pathsToURLS[pathURL.Path] = pathURL.URL
	}
	return pathsToURLS
}
