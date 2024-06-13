package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

// Uses multiple tags for YAML and JSON
type PathURL struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
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
	return MapHandler(mappedData, fallback), nil
}

func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedData, err := parseJSON(jsn)
	if err != nil {
		return nil, err
	}
	mappedData := buildURLMap(parsedData)
	return MapHandler(mappedData, fallback), nil
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

// Helper functino: unmarshals JSON data into a struct slice
func parseJSON(jsn []byte) ([]PathURL, error) {
	var jsonData []PathURL
	err := json.Unmarshal(jsn, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal JSON data: %v", err)
	}
	return jsonData, nil
}

// Helper function: builds a map of paths and URLs from a struct slice
func buildURLMap(parsedData []PathURL) map[string]string {
	pathsToURLS := make(map[string]string)

	// Iterate over the slice of PathURL and populate map with path: url
	for _, pathURL := range parsedData {
		pathsToURLS[pathURL.Path] = pathURL.URL
	}
	return pathsToURLS
}
