package scraper

import (
	"fmt"
	"time"

	"github.com/ivanmtoroc/go-domains/models"
)

// GetDomainByNameAPI return domain and servers information by domain name
func GetDomainByNameAPI(domainName string) (*models.Domain, []*models.Server, error) {
	// Create initial domain object
	domain := &models.Domain{
		Name:      domainName,
		Logo:      "Logo not found",
		Title:     "Title not found",
		IsValid:   true,
		CreatedAt: time.Now(),
	}

	fmt.Println("get information to domain: ", domainName)

	// Get domain and servers information from SSL Labs API
	servers, err := getServers(domain)
	if err != nil {
		return nil, nil, err
	} else if !domain.IsValid {
		return domain, nil, nil
	}

	// Get logo and title to domain using HTML scraping
	if !domain.IsDown {
		getLogoAndTitle(domain)
	}

	return domain, servers, nil
}
