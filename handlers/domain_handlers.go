package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ivanmtoroc/go-domains/models"
	"github.com/ivanmtoroc/go-domains/responses"
	"github.com/ivanmtoroc/go-domains/scraper"
	"github.com/ivanmtoroc/go-domains/utilities"
)

// GetDomain is http handler to '/domains/{domainName}' endpoint
func GetDomain(w http.ResponseWriter, r *http.Request) {
	var domain *models.Domain
	var servers []*models.Server
	var err error
	var isValid bool

	// Add CORS to response
	utilities.SetCORS(w)

	// Get 'domainName' URL parameter
	if domainName := chi.URLParam(r, "domainName"); domainName != "" {
		// Validate domainName
		if isValid, err = utilities.ValidateDomainName(domainName); err != nil {
			render.Render(w, r, responses.Error500)
			return
		} else if !isValid {
			render.Render(w, r, responses.Error404)
			return
		}
		// Use scraper to get domain and servers information
		domain, servers, err = scraper.GetDomainByNameAPI(domainName)
		if err != nil {
			render.Render(w, r, responses.Error500)
			return
		} else if !domain.IsValid {
			render.Render(w, r, responses.Error404)
			return
		}
	} else {
		// If URL parameter is not set
		render.Render(w, r, responses.Error400)
		return
	}

	// Verify if servers information changed in one hour or more before
	verifyServersChanges(domain, servers)

	// Save domain and servers into database
	domain.Save()
	for _, server := range servers {
		server.Save(domain)
	}

	// Create and render domain response
	if err := render.Render(w, r, responses.CreateDomainResponse(domain, servers)); err != nil {
		render.Render(w, r, responses.Error500)
	}
}

// verifyServersChanges function verify if the servers information of a domain changed
// one hour or more before
func verifyServersChanges(domain *models.Domain, servers []*models.Server) {
	// Get previous domain from database if exist
	previousDomain, err := models.GetDomainByNameDB(domain.Name)

	if err != nil {
		log.Println("get previous domain error")
		return
	} else if previousDomain == nil {
		return
	}

	// Get and parse time without nanoseconds and location
	previousDomainTime := time.Date(
		previousDomain.CreatedAt.Year(),
		previousDomain.CreatedAt.Month(),
		previousDomain.CreatedAt.Day(),
		previousDomain.CreatedAt.Hour(),
		previousDomain.CreatedAt.Minute(),
		previousDomain.CreatedAt.Second(),
		0, time.UTC,
	)
	// Get and parse time without nanoseconds and location
	domainTime := time.Date(
		domain.CreatedAt.Year(),
		domain.CreatedAt.Month(),
		domain.CreatedAt.Day(),
		domain.CreatedAt.Hour(),
		domain.CreatedAt.Minute(),
		domain.CreatedAt.Second(),
		0, time.UTC,
	)

	// Compare info only if duration between previous domain and domian is mayor or
	// equal that 1 hour
	if domainTime.Sub(previousDomainTime) >= (time.Hour) {
		fmt.Println("verifying changes in servers")
		// Set previous SSL grade to current domain
		domain.PreviousSslGrade = previousDomain.SslGrade
		fmt.Println("- previous_sslGrade: ", domain.PreviousSslGrade)
		// If changed status or SSL grade is enough to affirm that domain changed
		if domain.SslGrade != previousDomain.SslGrade ||
			domain.IsDown != previousDomain.IsDown {
			domain.ServersChanged = true
		} else {
			// Get servers of previous domain
			previousServers, err := models.GetServersByDomainDB(previousDomain)
			if err != nil {
				log.Println("get previous servers error")
				return
			}
			// Compare changes in servers information
			domain.ServersChanged = models.ServersChanges(servers, previousServers)
		}
		fmt.Println("- servers_changed: ", domain.ServersChanged)
	}
}
