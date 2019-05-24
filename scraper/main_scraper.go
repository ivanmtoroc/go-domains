package scraper

import (
  "time"
  "go-domains/models"
)

// Funtion that return domain information and servers by host name
func GetDomainByNameAPI(hostName string) (*models.Domain, []*models.Server, error) {
  // Create initial domain object
  domain := &models.Domain{0, hostName, false, "", "", "", "", false, true, time.Now()}

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
