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
  var err error

  // Add CORS to response
  utilities.SetCORS(w)

  // Get 'domain_name' URL parameter
  if domain_name := chi.URLParam(r, "domain_name"); domain_name != "" {
    // Use scraper to get domain and servers info
    domain, servers, err = scraper.GetDomainByNameAPI(domain_name)
    // Validate if domain is not valid
    if err != nil {
      render.Render(w, r, responses.ERROR_500)
      return
    } else if !domain.IsValid {
      render.Render(w, r, responses.ERROR_INVALID_DOMAIN)
      return
    }
  } else {
    // If URL parameter is not set
    render.Render(w, r, responses.ERROR_400)
    return
  }

  // Get previous domain from database
  previous_domain := models.GetDomainByNameDB(domain.Name)

  // Save domain and servers into database
  domain.Save()
  for _, server := range servers {
    server.Save()
  }

  // If exist previous domain verify servers changes
  if previous_domain != nil {
    verifyServersChanges(domain, previous_domain)
  }

  // Render domain response
  if err := render.Render(w, r, responses.CreateDomainResponse(domain, servers)); err != nil {
    render.Render(w, r, responses.ERROR_500)
    return
  }
}

// Function to verify if the servers of two domains are different
func verifyServersChanges(domain, previous_domain *models.Domain)  {
  // One hour duration
  one_hour, _ := time.ParseDuration("1h")

  domain_time := domain.CreatedAt
  // Add one hour to previous domain created at time
  previous_domain_time := previous_domain.CreatedAt.Add(one_hour)

  // Compare info only if previous domain time is <= to current domain time
  // Only if previous server is of one hour or more before
  // => if previous_domain_time + one_hour <= domain_time
  if previous_domain_time.Before(domain_time) ||
  previous_domain_time.Equal(domain_time) {
    // Set previous SSL grade to current domain
    domain.PreviousSslGrade = previous_domain.SslGrade
    // If changed status or SSL grade is enough to affirm that domain changed
    if domain.SslGrade != previous_domain.SslGrade ||
    domain.IsDown != previous_domain.IsDown {
      domain.ServersChanged = true
    } else {
      domain.ServersChanged = serversChanges(domain, previous_domain)
    }
  }
}

// Function to compare the servers of two domains
func serversChanges(domain, previous_domain *models.Domain) bool {
  // Get servers
  servers := models.GetServersByDomainDB(domain)
  previous_servers := models.GetServersByDomainDB(previous_domain)

  // If number of servers is different then servers changed
  if len(servers) != len(previous_servers) {
    return true
  }

  // Compare all servers
  for i, server := range servers {
    if !server.Equal(previous_servers[i]) {
      return true
    }
  }

  return false
}
