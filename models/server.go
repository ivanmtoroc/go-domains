package models

import (
  "time"
)

type Server struct {
  Address     string
  SslGrade    string
  Country     string
  Owner       string
  DomainName  string
  UpdatedAt   time.Time
}

func (server *Server) Create() {
  GetDB().Create(server)
}
