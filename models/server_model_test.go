package models

import (
  "time"
  "testing"
)

func TestServerEqual(t *testing.T) {
  server_one := &Server{0, "34.23.15.1", "A+", "CO", "Truora", 0}
  server_two := &Server{1, "34.23.15.1", "A+", "CO", "Truora", 0}
  server_three := &Server{2, "34.22.15.1", "A+", "CO", "Truora", 0}

  if equal := server_one.Equal(server_two); !equal {
    t.Error("Servers is actually equals")
  }

  if equal := server_one.Equal(server_three); equal {
    t.Error("Servers is actually differents")
  }
}

func TestServerSaveAndGetServers(t *testing.T) {
  domain := &Domain{0, "truora.com", false, "", "", "", "", false, true, time.Now()}
  domain.Save()

  server := &Server{0, "35.11.51.12", "C", "CO", "Azure", 0}
  other_server := &Server{0, "34.12.52.15", "B", "US", "Amazon", 0}
  server.Save(domain)
  other_server.Save(domain)

  servers := GetServersByDomainDB(domain)

  if !server.Equal(servers[1]) {
    t.Error("Server one not saved")
  } else if !other_server.Equal(servers[0]) {
    t.Error("Server two not saved")
  }
}
