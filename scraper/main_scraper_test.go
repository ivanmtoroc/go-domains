package scraper

import (
	"testing"
)

func TestGetDomainAPI(t *testing.T) {
	domainName := "truora.com"
	domain, _, err := GetDomainByNameAPI(domainName)

	if err != nil {
		t.Error("get domain info from API failed")
	}
	if domain.Title != "Truora" {
		t.Error("get title to Truora web page failed")
	}
}
