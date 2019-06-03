package responses

import (
	"github.com/ivanmtoroc/go-domains/models"
)

// ServerResponse structure to create servers JSON responses
type ServerResponse struct {
	Address  string `json:"address"`
	SslGrade string `json:"ssl_grade"`
	Country  string `json:"country"`
	Owner    string `json:"owner"`
}

// Create new server response by server
func createServerResponse(server *models.Server) *ServerResponse {
	// Create new server response
	return &ServerResponse{
		Address:  server.Address,
		SslGrade: server.SslGrade,
		Country:  server.Country,
		Owner:    server.Owner,
	}
}
