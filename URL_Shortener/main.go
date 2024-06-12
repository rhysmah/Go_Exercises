package main

import (
	"fmt"
	"net/http"
	"url_shortener/urlshort"
)

func main() {

	yamlData := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	pathsToUrls := map[string]string{
		"/dog": "www.samplesite.com/article-on-dogs",
		"/cat": "www.samplesite.com/article-on-cats",
	}

	// Create multiplexer; this is the default
	mux := defaultMux()

	// Create the MapHandler; if there are no URLs that match
	// what's stored in the map, it will default to `mux`
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Create the yamlHandler; if this causes an error, the
	// program will panic. If there are URLs that match those
	// stored in YAML, it will call those. Else, it will
	// default to the mapHandler
	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlData), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Start server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

// defaultMux returns a ServeMux with a default handler
// that will be used if no custom handler is provided.
// A multiplexer (ServerMux) matches the URL of each incoming
// request against a list of registered patterns and calls the
// handler for the pattern that most closely matches the URL.
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	return mux
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}
