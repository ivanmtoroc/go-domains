package api

import (
  "net/http"
  "go-domains/models"
)

type DomainResponse struct {
  Servers          []*ServerResponse `json:"servers"`
  ServersChanged   bool		           `json:"servers_changed"`
  SslGrade         string	           `json:"ssl_grade"`
  PreviousSslGrade string	           `json:"previous_ssl_grade"`
  Logo             string	           `json:"logo"`
  IsDown           bool		           `json:"is_down"`
}

func CreateDomainResponse(domain *models.Domain, servers []*models.Server) *DomainResponse {
  var servers_response []*ServerResponse

  for _, server := range servers {
    servers_response = append(servers_response, createServerResponse(server))
  }

	return &DomainResponse{
    servers_response,
    domain.ServersChanged,
    domain.SslGrade,
    domain.PreviousSslGrade,
    domain.Logo,
    domain.IsDown,
  }
}

func (dr *DomainResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
