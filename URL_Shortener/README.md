## URL Shortener

This exercise is based on Exercise #2: URL Shortener on the [gophercises GitHub](https://github.com/gophercises/urlshort).

### Overview

URL shorteners are typically used to turn long URLs into shorter, more human-readable URLs (for promotional purposes, as an example). When a user clicks on the shortened URL or types it in, the shortener service checks if the URL is valid and redirects the user to the original, "real" URL.

This exercise simulates a URL shortener service that reads an incoming URL request, check that exists within a map, then serves the corresponding URL. If the URL is not found, the service returns a "fallback" that redirects the user to a 404 error page.