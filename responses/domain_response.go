package responses

import (
	"net/http"

	"github.com/ivanmtoroc/go-domains/models"
)

// DomainResponse structure to create domains JSON responses
type DomainResponse struct {
	Servers          []*ServerResponse `json:"servers"`
	ServersChanged   bool              `json:"serversChanged"`
	SslGrade         string            `json:"sslGrade"`
	PreviousSslGrade string            `json:"previousSslGrade"`
	Logo             string            `json:"logo"`
	Title            string            `json:"title"`
	IsDown           bool              `json:"isDown"`
}

// CreateDomainResponse create new domain response by domain and your servers
func CreateDomainResponse(domain *models.Domain, servers []*models.Server) *DomainResponse {
	var serversResponse = make([]*ServerResponse, 0)
	// Create array of servers responses
	for _, server := range servers {
		serversResponse = append(serversResponse, createServerResponse(server))
	}
	// Create domain respose
	return &DomainResponse{
		Servers:          serversResponse,
		ServersChanged:   domain.ServersChanged,
		SslGrade:         domain.SslGrade,
		PreviousSslGrade: domain.PreviousSslGrade,
		Logo:             domain.Logo,
		Title:            domain.Title,
		IsDown:           domain.IsDown,
	}
}

// Render is a method of Renderer interface of go-chi/render for managing response payloads
func (dr *DomainResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
