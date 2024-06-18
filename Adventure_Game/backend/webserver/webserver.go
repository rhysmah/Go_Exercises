package webserver

import (
	"adventure/jsonhandler"
	"encoding/json"
	"log"
	"net/http"
)

// storyHandler reads the JSON file containing the story data, parses it, and writes it to the response body.
// If there is an error reading the file, the error is written to the response body and the function returns.
func storyHandler(w http.ResponseWriter, r *http.Request) {

	// Attempt to read the file containing the story; if error, write
	// error to response body and return
	storyData, err := jsonhandler.ReadFile("gopher.json")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Attempt to parse the JSON data; if error, write
	// error to response body and return
	storyDataParsed, err := jsonhandler.ParseJSON(storyData)
	if err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusInternalServerError)
		return
	}

	// Set the response header to JSON; encode data as JSON and write to response body.
	// If error, write error to response body and return
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(storyDataParsed)
	if err != nil {
		http.Error(w, "Error encoding JSON data", http.StatusInternalServerError)
		return
	}
}

// StartServer starts the server on port 8080 and listens for requests to the /story endpoint.
func StartServer() {
	http.HandleFunc("/story", storyHandler)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
