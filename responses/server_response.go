package responses

import (
  "go-domains/models"
)

// Struct used to create server JSON responses
type ServerResponse struct {
  Address   string  `json:"address"`
  SslGrade  string  `json:"ssl_grade"`
  Country   string  `json:"country"`
  Owner     string  `json:"owner"`
}

// Create new server response by server
func createServerResponse(server *models.Server) *ServerResponse {
  return &ServerResponse{
      server.Address,
      server.SslGrade,
      server.Country,
      server.Owner,
  }
}
