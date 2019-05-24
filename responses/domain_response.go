package responses

import (
  "net/http"
  "go-domains/models"
)

// Struct used to create domain JSON responses
type DomainResponse struct {
  Servers           []*ServerResponse  `json:"servers"`
  ServersChanged    bool               `json:"servers_changed"`
  SslGrade          string             `json:"ssl_grade"`
  PreviousSslGrade  string             `json:"previous_ssl_grade"`
  Logo              string             `json:"logo"`
  Title             string             `json:"title"`
  IsDown            bool               `json:"is_down"`
}

// Create new domain response by domain and your servers
func CreateDomainResponse(domain *models.Domain, servers []*models.Server) *DomainResponse {
  var servers_response []*ServerResponse = make([]*ServerResponse, 0)

  // Create array of servers responses
  for _, server := range servers {
    servers_response = append(servers_response, createServerResponse(server))
  }

  // Create domain respose
  return &DomainResponse{
    servers_response,
    domain.ServersChanged,
    domain.SslGrade,
    domain.PreviousSslGrade,
    domain.Logo,
    domain.Title,
    domain.IsDown,
  }
}

// Method used to Renderer interface of go-chi/render for managing response payloads
func (dr *DomainResponse) Render(w http.ResponseWriter, r *http.Request) error {
  return nil
}
