package handlers

import (
	"time"
	"net/http"
	"go-domains/responses"
	"go-domains/models"
	"go-domains/scraper"
	"go-domains/utilities"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Function to handle requests to '/domains/{domain_name}' endpoint
func GetDomain(w http.ResponseWriter, r *http.Request) {
	var domain *models.Domain
	var servers []*models.Server

	// Add CORS to response
	utilities.SetCORS(w)

	if domain_name := chi.URLParam(r, "domain_name"); domain_name != "" {
		domain, servers = scraper.GetDomainByNameAPI(domain_name)
		if domain == nil {
			render.Render(w, r, responses.ERROR_INVALID_DOMAIN)
			return
		}
	} else {
		render.Render(w, r, responses.ERROR_400)
		return
	}

	verifyChanges(domain, servers)

	domain.Save()
	for _, server := range servers {
		server.Save()
	}

	if err := render.Render(w, r, responses.CreateDomainResponse(domain, servers)); err != nil {
		render.Render(w, r, responses.ERROR_500)
		return
	}
}

func verifyChanges(domain *models.Domain, servers []*models.Server)  {
	previous_domain := models.GetDomainByNameDB(domain.Name)
	if previous_domain == nil {
		return
	}

	one_hour, _ := time.ParseDuration("1h")
	previous_domain_time := previous_domain.CreatedAt.Add(one_hour)
	domain_time := domain.CreatedAt

	if previous_domain_time.Before(domain_time) || previous_domain_time.Equal(domain_time) {
		domain.PreviousSslGrade = previous_domain.SslGrade
		if domain.SslGrade != previous_domain.SslGrade {
			domain.ServersChanged = true
		} else {
			domain.ServersChanged = serversChanges(servers, previous_domain)
		}
	}
}

func serversChanges(servers []*models.Server, previous_domain *models.Domain) bool {
	previous_servers := models.GetServersByDomainDB(previous_domain)
	if len(servers) != len(previous_servers) {
		return true
	}

	byAddress := func (server_one, server_two models.Server) bool {
		return server_one.Address < server_two.Address
	}
	utilities.OrderServersBy(byAddress).Sort(servers)
	utilities.OrderServersBy(byAddress).Sort(previous_servers)

	for i, server := range servers {
		if !server.Equal(previous_servers[i]) {
			return true
		}
	}

	return false
}
