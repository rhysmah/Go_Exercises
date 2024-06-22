package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

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

	documentLinks, err := parseHTML(tokenizer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(documentLinks)
}

func parseHTML(tokenizer *html.Tokenizer) ([]Link, error) {
	var documentLinks []Link

	// Loop through HTML tokens
	for {
		tokenType := tokenizer.Next()

		// ErrorToken can be parsing error or EOF
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
					fmt.Println(err)
					continue
				}
				documentLinks = append(documentLinks, link)
			}
		}
	}
	return documentLinks, nil
}

func parseLink(tokenizer *html.Tokenizer) (Link, error) {
	link := Link{}

	// (1) Extract href
	hrefContent, err := extractHref(tokenizer)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	link.Href = hrefContent

	// (2) Extract Tag Content
	tagContent, err := extractTagContent(tokenizer, "a")
	if err != nil {
		fmt.Println("Error reading tag content: ", err)
	}
	link.Link = tagContent

	return link, nil
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
