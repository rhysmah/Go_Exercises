package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

const aLinkToken = "a"
const hrefToken = "href"

type Link struct {
	Href string
	Link string
}

func main() {

	htmlString := `<html><body><h1>Hello!</h1><a href="/other-page">A link to another page</a></body></html>`

	// Creates a new reader for strings; this implements io.Reader
	// interface, allowing us to treat string as a stream of bytes
	r := strings.NewReader(htmlString)
	tokenizer := html.NewTokenizer(r)

	// Loop through HTML tokens
	for {
		tokenType := tokenizer.Next()

		// ErrorToken can be a parsing error or EOF
		if tokenType == html.ErrorToken {
			if tokenizer.Err() != io.EOF {
				fmt.Println("Error parsing HTML: ", tokenizer.Err())
			}
			break
		}

		// No error, no EOF; start parsing tokens

		if tokenType == html.StartTagToken {
			linkData := Link{}
			tagName, hasAttribute := tokenizer.TagName()

			if string(tagName) == aLinkToken {

				// Extract attribute
				if hasAttribute {
					attrName, attrContent, _ := tokenizer.TagAttr()
					if string(attrName) == hrefToken {
						linkData.Href = string(attrContent)
					}
				}

				// Extract tag content
				tagContent, err := readTagContent(tokenizer, "a")
				if err != nil {
					fmt.Println("Error reading tag content: ", err)
				} else {
					linkData.Link = tagContent
				}
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
