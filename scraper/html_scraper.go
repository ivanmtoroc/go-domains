package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ivanmtoroc/go-domains/models"
	"golang.org/x/net/html"
)

// getLogoAndTitle get logo and title to domain from web page
func getLogoAndTitle(domain *models.Domain) {
	url := "http://" + domain.Name
	fmt.Println("get logo and title from web page")

	// Make http request to web page of domain
	response, err := http.Get(url)
	if err != nil {
		log.Println("get response to web page of domain failed")
		return
	}
	// Create HTML Tokenizer of response body
	tonkenizer := html.NewTokenizer(response.Body)
	// Get Logo URL from HTML Tokenizer
	domain.Logo = getLogo(tonkenizer, url)
	fmt.Printf("- logo: %s\n", domain.Logo)

	// Make http request to web page of domain
	response, err = http.Get(url)
	if err != nil {
		log.Println("get response to web page of domain failed")
		return
	}
	// Parse response body to HTML node tree
	rootNode, err := html.Parse(response.Body)
	if err != nil {
		log.Println("response body parse error")
		return
	}
	// Get title from HTML node tree
	domain.Title, _ = getTitle(rootNode)
	fmt.Printf("- title: %s\n", domain.Title)
}

// getLogo iterate in HTML Tokenaizer to find icon URL
func getLogo(document *html.Tokenizer, url string) string {
	// Icon URL value
	value := ""
	// Is true if icon is found
	found := false

	for {
		// Get next Token from HTML Tokenaizer
		tokenType := document.Next()
		switch tokenType {
		case html.ErrorToken:
			// If error occurred or is end of document
			return "Logo not found"
		default:
			// Get current token
			token := document.Token()
			// Validate if token is a link tag
			if token.Data == "link" {
				// Search into tag link the icon URL
				value, found = getLogoFromToken(token)
			}
		}
		// Validate if icon URL is found
		if found {
			break
		}
	}
	// If icon URL is relative append domain URL
	if string(value[0]) == "/" {
		value = url + value
	}
	return value
}

// getLogoFromToken search into HTML token to icon URL
func getLogoFromToken(token html.Token) (string, bool) {
	// Get value of attribute rel if exist
	value, found := getValueOfAttr(token, "rel")
	// Validate if attribute value contain 'icon' word
	if found && strings.Contains(value, "icon") {
		// Get value of attribute href
		value, found = getValueOfAttr(token, "href")
		if found {
			return value, true
		}
	}
	return "", false
}

// getValueOfAttr get the value of an attribute of a token
func getValueOfAttr(token html.Token, attribute string) (string, bool) {
	for _, attr := range token.Attr {
		// Validate if attribute exist
		if attr.Key == attribute {
			// Get value of attribute
			return string(attr.Val), true
		}
	}
	return "", false
}

// getTitle search recursively in HTML node tree to find the value of title tag
func getTitle(node *html.Node) (string, bool) {
	// Stop case: if node is a title tag
	if node.Type == html.ElementNode && node.Data == "title" {
		return node.FirstChild.Data, true
	}
	// Recursive case: Iterate by depth in HTML nodes
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		title, found := getTitle(c)
		// Stop case: If title tag is found
		if found {
			return title, true
		}
	}
	// If it is a leaf
	return "Title not found", false
}
