package scraper

import (
  "testing"
)

func TestGetDomainAPI(t *testing.T) {
  domain, servers, err := GetDomainByNameAPI("truora.com")
  if err != nil {
    t.Error("Get domain info from API failed")
  }

  if domain.Title != "Truora" {
    t.Error("Title of Truora web page is wrong")
  }
}
