package scraper

import (
  "fmt"
  "time"
  "go-domains/models"
)

// Funtion that return domain information and servers by domain name
func GetDomainByNameAPI(domain_name string) (*models.Domain, []*models.Server, error) {
  // Create initial domain object
  domain := &models.Domain{0, domain_name, false, "", "", "Icon not found", "Title not found", false, true, time.Now()}

  fmt.Println("get information to domain: ", domain_name)

  // Get servers and complete information of domain object from SSL Labs API
  servers, err := getServers(domain)
  if err != nil {
    return nil, nil, err
  } else if !domain.IsValid {
    return domain, nil, nil
  }

  // Set Icon and Title to domain by html scraper
  if !domain.IsDown {
    setIconAndTitle(domain)
  }

  return domain, servers, nil
}
