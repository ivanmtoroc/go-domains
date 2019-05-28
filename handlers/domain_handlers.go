package handlers

import (
  "fmt"
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
    // Validate domain_name
    if is_valid := utilities.ValidateDomainName(domain_name); !is_valid {
      render.Render(w, r, responses.ERROR_INVALID_DOMAIN)
      return
    }
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

  verifyServersChanges(domain, servers)

  // Save domain and servers into database
  domain.Save()
  for _, server := range servers {
    server.Save(domain)
  }

  // Render domain response
  if err := render.Render(w, r, responses.CreateDomainResponse(domain, servers)); err != nil {
    render.Render(w, r, responses.ERROR_500)
    return
  }
}

// Function to verify if the servers of two domains are different
func verifyServersChanges(domain *models.Domain, servers []*models.Server)  {
  // Get previous domain from database
  previous_domain := models.GetDomainByNameDB(domain.Name)
  if previous_domain == nil {
    return
  }

  // Get and parse time without ns and location
  previous_domain_time := time.Date(
    previous_domain.CreatedAt.Year(),
    previous_domain.CreatedAt.Month(),
    previous_domain.CreatedAt.Day(),
    previous_domain.CreatedAt.Hour(),
    previous_domain.CreatedAt.Minute(),
    previous_domain.CreatedAt.Second(),
    0, time.UTC,
  )
  // Get and parse time without ns and location
  domain_time := time.Date(
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
  if domain_time.Sub(previous_domain_time) >= (1 * time.Hour) {
    fmt.Println("verifying changes in servers")

    // Set previous SSL grade to current domain
    domain.PreviousSslGrade = previous_domain.SslGrade
    fmt.Println("- previous_sslGrade: ", domain.PreviousSslGrade)
    // If changed status or SSL grade is enough to affirm that domain changed
    if domain.SslGrade != previous_domain.SslGrade ||
    domain.IsDown != previous_domain.IsDown {
      domain.ServersChanged = true
    } else {
      // Get servers of previous domain
      previous_servers := models.GetServersByDomainDB(previous_domain)
      // Compare changes in servers
      domain.ServersChanged = models.ServersChanges(servers, previous_servers)
    }
    fmt.Println("- servers_changed: ", domain.ServersChanged)
  }
}
