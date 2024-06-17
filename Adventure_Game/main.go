package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type StoryData struct {
	Story StoryArc `json:"intro"`
}

type StoryArc struct {
	Title   string         `json:"title"`
	Story   []string       `json:"story"`
	Options []StoryOptions `json:"options"`
}

type StoryOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func readFile(fileName string) ([]byte, error) {

	// Attempt to open file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully opened file:", fileName)
	defer file.Close()

	// Read file
	contentAsBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully read data from file:", fileName)
	return contentAsBytes, nil
}

func parseJSON(jsn []byte) (StoryData, error) {
	var jsonData StoryData
	err := json.Unmarshal(jsn, &jsonData)

	// If error, return empty struct
	if err != nil {
		return StoryData{}, fmt.Errorf("cannot unmarshal JSON: %v", err)
	}

	log.Println("Successfully parsed JSON data")
	return jsonData, nil
}

func main() {
	content, err := readFile("gopher.json")
	if err != nil {
		log.Fatal(err)
	}

	storyData, err := parseJSON(content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(storyData)
}
