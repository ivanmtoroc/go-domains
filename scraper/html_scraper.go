package scraper

import (
	"strings"
	"net/http"
	"golang.org/x/net/html"
	"go-domains/models"
)

func setIconAndTitle(domain *models.Domain) {
	url := "http://" + domain.Name

	response, err := http.Get(url)
	if err != nil {
		domain = nil
		return
	}

	element, _ := html.Parse(response.Body)
	domain.Title, _ = getTitle(element)

	response, err = http.Get(url)
	if err != nil {
		domain = nil
		return
	}

	document := html.NewTokenizer(response.Body)
	domain.Logo = getIcon(document, url)
}

func getIcon(document *html.Tokenizer, url string) string {
	value := ""
	found := false

	for {
		if found {
			break
		}

		tokenType := document.Next()
		switch tokenType {
			case html.SelfClosingTagToken:
				token := document.Token()
				if token.Data == "link" {
					value, found = getIconFromToken(token)
				}
			case html.ErrorToken:
				return "Icon not found"
			default:
				token := document.Token()
				if token.Data == "link" {
					value, found = getIconFromToken(token)
				}
		}
	}
	if string(value[0]) == "/" {
		value = url + value
	}
	return value
}

func getIconFromToken(token html.Token) (string, bool) {
	value, found := getValueOfAttr(token, "rel")
	if found && strings.Contains(value, "icon") {
		value, found = getValueOfAttr(token, "href")
		if found {
			return value, true
		}
	}
	return "", false
}

func getValueOfAttr(token html.Token, attribute string) (string, bool) {
	for _, attr := range token.Attr {
		if attr.Key == attribute {
			return string(attr.Val), true
		}
	}
	return "", false
}

func getTitle(element *html.Node) (string, bool) {

  if element.Type == html.ElementNode && element.Data == "title" {
    return element.FirstChild.Data, true
  }

  for c := element.FirstChild; c != nil; c = c.NextSibling {
    result, found := getTitle(c)
    if found {
      return result, true
    }
  }

  return "Title not found", false
}
