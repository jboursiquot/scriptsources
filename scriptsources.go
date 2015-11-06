// Package scriptsources searches through HTML for script tags and returns the URLs found in the src attributes.
package scriptsources

import (
	"golang.org/x/net/html"
	"io"
)

// ExtractResources takes an io.Reader (such as an http response body), parses and returns the resource references in the script tags.
func ExtractResources(body io.Reader) (urls []string) {
	tokenizer := html.NewTokenizer(body)

	for {
		tokenType := tokenizer.Next()
		switch {
		case tokenType == html.ErrorToken:
			return
		case tokenType == html.StartTagToken:
			token := tokenizer.Token()
			if token.Data != "script" {
				continue
			}
			url := extractResource(&token)
			if len(url) > 0 {
				urls = append(urls, url)
			}
		}
	}
}

func extractResource(token *html.Token) (url string) {
	for _, attr := range token.Attr {
		if attr.Key == "src" {
			url = attr.Val
		}
	}
	return
}
