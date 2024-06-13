package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"url_shortener/urlshort"
)

func main() {

	filename := flag.String("file", "data.yaml", "YAML or JSON file with `path: url` data")
	flag.Parse()

	data, fileExt, err := loadData(*filename)
	if err != nil {
		log.Fatal("Failed to load data:", err)
	}

	// Instantiate the default multiplexer
	mux := defaultMux()

	// Create the map handler with no initial mappings
	mapHandler := urlshort.MapHandler(nil, mux)

	// Determine the handler based on file extenstion type
	var handler http.Handler
	switch fileExt {
	case "yaml":
		handler, err = urlshort.YAMLHandler(data, mapHandler)
	case "json":
		handler, err = urlshort.JSONHandler(data, mapHandler)
	default:
		log.Fatal("unsupported file extension")
	}
	if err != nil {
		log.Fatal("Failed to create handler: ", err)
	}

	fmt.Println("Start server on :8080")
	http.ListenAndServe(":8080", handler)
}

func loadData(filename string) ([]byte, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, "", fmt.Errorf("failed to open file: %w", err)
	}
	log.Print("Successfully opened file: ", filename)
	defer file.Close()

	// Read and store the file extension
	filePathExt := filepath.Ext(filename)
	var fileExt string
	if filePathExt == ".yaml" || filePathExt == ".yml" {
		fileExt = "yaml"
	} else if filePathExt == ".json" {
		fileExt = "json"
	} else {
		return nil, "", fmt.Errorf("unsupported file extension: %s", filePathExt)
	}

	// Read content of file
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, "", err
	}
	log.Print("Successfully read data from file: ", filename)

	return content, fileExt, nil
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
