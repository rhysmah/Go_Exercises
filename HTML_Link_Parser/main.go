package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Link string
}

func main() {

	htmlContent, err := processFile("ex4.html")
	if err != nil {
		log.Fatalln("Error processing file: ", err)
	}

	// Creates a new reader for strings; this implements io.Reader
	// interface, allowing us to treat string as a stream of bytes
	r := strings.NewReader(htmlContent)
	tokenizer := html.NewTokenizer(r)

	documentLinks, err := parseHTML(tokenizer)
	if err != nil {
		log.Fatalln("Error parsing HTML: ", err)
	}

	for _, content := range documentLinks {
		fmt.Printf("HREF: %s\nLink: %s\n", content.Href, content.Link)
	}
}

func processFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %v", filename, err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read html content: %v", err)
	}

	return string(fileContent), nil
}

func parseHTML(tokenizer *html.Tokenizer) ([]Link, error) {
	var documentLinks []Link

	// Loop through HTML tokens
	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			if tokenizer.Err() != io.EOF {
				return nil, tokenizer.Err()
			}
			break // EOF; break loop
		}

		if tokenType == html.StartTagToken {
			tagName, hasAttr := tokenizer.TagName()

			if string(tagName) == "a" && hasAttr {
				link, err := parseLink(tokenizer)
				if err != nil {
					log.Println("Error parsing link: ", err)
					continue
				}
				documentLinks = append(documentLinks, link)
			}
		}
	}
	return documentLinks, nil
}

func parseLink(tokenizer *html.Tokenizer) (Link, error) {

	// (1) Extract href
	hrefContent, err := extractHref(tokenizer)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	// (2) Extract Tag Content
	tagContent, err := extractTagContent(tokenizer, "a")
	if err != nil {
		fmt.Println("Error reading tag content: ", err)
	}

	return Link{
		Href: hrefContent,
		Link: strings.TrimSpace(strings.Join(strings.Fields(tagContent), " ")),
	}, nil
}

func extractHref(tokenizer *html.Tokenizer) (string, error) {
	for {
		attrName, attrContent, moreAttr := tokenizer.TagAttr()
		if string(attrName) == "href" {
			return string(attrContent), nil
		}

		// If no `href`, end infinite loop
		if !moreAttr {
			break
		}
	}

	return "", nil
}

func extractTagContent(tokenizer *html.Tokenizer, tagName string) (string, error) {
	var tagContent strings.Builder

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			if tokenizer.Err() != io.EOF {
				return "", tokenizer.Err()
			}
			break
		}

		switch tokenType {

		case html.TextToken:
			text := tokenizer.Text()
			tagContent.WriteString(string(text))

		case html.EndTagToken:
			endTagName, _ := tokenizer.TagName()
			if string(endTagName) == tagName {
				return tagContent.String(), nil
			}
		}
	}
	return tagContent.String(), nil
}
