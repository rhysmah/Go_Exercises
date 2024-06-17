package main

import (
	"adventure/jsonhandler"
	"fmt"
	"log"
)

func main() {
	content, err := jsonhandler.ReadFile("gopher.json")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	storyData, err := jsonhandler.ParseJSON(content)
	if err != nil {
		log.Fatal("Error parsing JSON data: ", err)
	}

	fmt.Println(storyData)
}
