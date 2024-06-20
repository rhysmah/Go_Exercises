package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	htmlString := `<html><body><h1>Hello!</h1><a href="/other-page">A link to another page</a></body></html>`

	// Creates a new reader for strings; this implements the io.Reader
	// interface; it allows us to treat the string as a stream of bytes
	r := strings.NewReader(htmlString)
	tokenizer := html.NewTokenizer(r)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			if tokenizer.Err() != io.EOF {
				fmt.Println("Error parsing HTML: ", tokenizer.Err())
			}
			return
		}

		switch tokenType {
		case html.StartTagToken:
			tagName, _ := tokenizer.TagName()
			if string(tagName) == "a" {
				tagContent, err := readTagContent(tokenizer, string(tagName))
				if err != nil {
					fmt.Println("Error: ", err)
				}
				print(tagContent)
			}
		}
	}
}

func readTagContent(tokenizer *html.Tokenizer, tagName string) (string, error) {
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

		case html.StartTagToken:
			nestedTagName, _ := tokenizer.TagName()
			tagContent.WriteString(fmt.Sprintf("<%s>", string(nestedTagName)))

		case html.SelfClosingTagToken:
			selfClosingTagName, _ := tokenizer.TagName()
			tagContent.WriteString(fmt.Sprintf("<%s>", string(selfClosingTagName)))
		}
	}
	return tagContent.String(), nil
}
