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
  CreatedAt   time.Time
}

func (server *Server) Equal(other_server *Server) bool {
  if server.Address != other_server.Address {
    return false
  } else if server.SslGrade != other_server.SslGrade {
    return false
  } else if server.Country != other_server.Country {
    return false
  } else if server.Owner != other_server.Owner {
    return false
  } else if server.DomainName != other_server.DomainName {
    return false
  }

  return true
}

func (server *Server) Save() {
  GetDB().Create(server)
}

func GetServersDB(domain *Domain) []*Server {
  var servers []*Server
  GetDB().Table("servers").Where(
      "domain_name = ? AND created_at >= ?",
      domain.Name,
      domain.CreatedAt,
  ).Find(&servers)
  if len(servers) == 0 {
		return nil
	}
  return servers
}
