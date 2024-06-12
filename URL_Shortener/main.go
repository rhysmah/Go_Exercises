package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"url_shortener/urlshort"
)

func main() {

	yamlFileName := flag.String("yaml", "data.yaml", "YAML file with path: url data")
	flag.Parse()

	// Map of paths to URLs; this can be converted to a JSON file,
	// which will require additional functions in `handler.go`
	pathsToUrls := map[string]string{
		"/dog": "http://www.samplesite.com/article-on-dogs",
		"/cat": "http://www.samplesite.com/article-on-cats",
	}

	yamlData, err := loadYAMLData(*yamlFileName)
	if err != nil {
		log.Fatal(err)
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
	yamlHandler, err := urlshort.YAMLHandler(yamlData, mapHandler)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func loadYAMLData(filename string) ([]byte, error) {
	data, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	log.Print("Successfully opened file: ", filename)
	defer data.Close()

	yamlData, dataErr := io.ReadAll(data)
	if dataErr != nil {
		return nil, err
	}
	log.Print("Successfully read data from file: ", filename)

	return yamlData, nil
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
