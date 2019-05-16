package scraper

import (
	"time"
	"go-domains/models"
)

func GetDomainByNameAPI(hostName string) (*models.Domain, []*models.Server) {
	domain := &models.Domain{hostName, false, "", "", "", false, "", time.Now()}
	servers := GetServers(domain)

	if domain == nil {
		return nil, nil
	}

	setIconAndTitle(domain)

	return domain, servers
}
