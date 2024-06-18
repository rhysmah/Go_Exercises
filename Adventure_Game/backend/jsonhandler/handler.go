package jsonhandler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type StoryData struct {
	Intro StoryArc `json:"intro"`
}

type StoryArc struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
}

type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func ReadFile(fileName string) ([]byte, error) {

	// Attempt to open file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully opened file:", fileName)
	defer file.Close()

	// Read file
	contentAsByteSlice, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully read data from file:", fileName)
	return contentAsByteSlice, nil
}

func ParseJSON(jsn []byte) (StoryData, error) {
	var jsonData StoryData
	err := json.Unmarshal(jsn, &jsonData)

	// If error, return empty struct
	if err != nil {
		return StoryData{}, fmt.Errorf("cannot unmarshal JSON: %v", err)
	}

	log.Println("Successfully parsed JSON data")
	return jsonData, nil
}
