package scraper

import (
  "fmt"
  "log"
  "strings"
  "net/http"
  "golang.org/x/net/html"
  "go-domains/models"
)

// Set icon and title domain from HTML page
func setIconAndTitle(domain *models.Domain) {
  url := "http://" + domain.Name

  fmt.Println("get title and icon from web page")

  // Make http request (Method Get) to page of domain
  response, err := http.Get(url)
  if err != nil {
    log.Println("get response to web page of domain failed")
    return
  }
  // Parse response body to HTML node tree
  element, _ := html.Parse(response.Body)
  // Get title from HTML node tree
  domain.Title, _ = getTitle(element)
  fmt.Printf("- title: %s\n", domain.Title)

  // Make http request (Method Get) to page of domain
  response, err = http.Get(url)
  if err != nil {
    log.Println("get response to web page of domain failed")
    return
  }
  // Create HTML Tokenizer of response body
  document := html.NewTokenizer(response.Body)
  // Get Icon URL from Tokenizer
  domain.Logo = getIcon(document, url)
  fmt.Printf("- logo: %s\n", domain.Logo)
}

// Function that iterate in HTML Tokenaizer to find icon URL of HTML page
func getIcon(document *html.Tokenizer, url string) string {
  // Icon URL
  value := ""
  // Is true if icon is found
  found := false

  for {
    // Get next Token from HTML Tokenaizer
    tokenType := document.Next()
    switch tokenType {
      case html.ErrorToken:
        // If error occurred or is end of document
        return "Icon not found"
      default:
        // Get current token
        token := document.Token()
        // Validate if token is a link tag
        if token.Data == "link" {
          // Search into tag link the icon URL
          value, found = getIconFromToken(token)
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

// Function to search into token icon URL
func getIconFromToken(token html.Token) (string, bool) {
  // Get value of attribute rel
  value, found := getValueOfAttr(token, "rel")
  // Validate if value contain 'icon' word
  if found && strings.Contains(value, "icon") {
    // Get value of attribute href
    value, found = getValueOfAttr(token, "href")
    if found {
      return value, true
    }
  }
  return "", false
}

// Get the value of an attribute of a token
func getValueOfAttr(token html.Token, attribute string) (string, bool) {
  for _, attr := range token.Attr {
    // Validate if attribute exist
    if attr.Key == attribute {
      return string(attr.Val), true
    }
  }
  return "", false
}

// Function that search recursively HTML node tree to find the value of title tag
func getTitle(element *html.Node) (string, bool) {

  // Stop case: if node is a title tag
  if element.Type == html.ElementNode && element.Data == "title" {
    return element.FirstChild.Data, true
  }

  // Recursive case: Iterate by depth in HTML node
  for c := element.FirstChild; c != nil; c = c.NextSibling {
    result, found := getTitle(c)
    if found {
      return result, true
    }
  }

  // If it is a leaf
  return "Title not found", false
}
