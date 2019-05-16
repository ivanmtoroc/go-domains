package responses

import (
  "go-domains/models"
)

type ServerResponse struct {
	Address   string  `json:"address"`
  SslGrade  string  `json:"ssl_grade"`
  Country   string  `json:"country"`
  Owner     string  `json:"owner"`
}

func createServerResponse(server *models.Server) *ServerResponse {
	return &ServerResponse{
      server.Address,
      server.SslGrade,
      server.Country,
      server.Owner,
  }
}
