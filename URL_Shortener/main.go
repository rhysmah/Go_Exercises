package main

import (
	"fmt"
	"net/http"
	"url_shortener/urlshort"
)

func main() {

	fmt.Println("Hello, world")

	pathsToUrls := map[string]string{
		"/dog": "www.samplesite.com/article-on-dogs",
		"/cat": "www.samplesite.com/article-on-cats",
	}

	fallback := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Page not found", http.StatusNotFound)
		})

	urlshort.MapHandler(pathsToUrls, fallback)
}
