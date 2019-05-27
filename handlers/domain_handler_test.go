package handlers

import (
  "time"
  "testing"
  "go-domains/models"
)

func TestVerifyServersChanges(t *testing.T) {
  now := time.Now()

  previous_domain := &models.Domain{0, "truora.com", false, "A", "", "", "", false, true, now}
  previous_domain.Save()
  previous_server_one := &models.Server{0, "32.32.32.32", "A", "CO", "Amazon", 0}
  previous_server_one.Save(previous_domain)
  previous_server_two := &models.Server{0, "33.33.33.33", "B", "US", "Google", 0}
  previous_server_two.Save(previous_domain)

  domain := &models.Domain{1, "truora.com", false, "A", "", "", "", false, true, now.Add(time.Hour)}
  var servers []*models.Server
  server_one := &models.Server{0, "32.32.32.32", "A", "CO", "Amazon", 1}
  servers = append(servers, server_one)
  server_two := &models.Server{0, "33.33.33.32", "B", "US", "Google", 1}
  servers = append(servers, server_two)

  verifyServersChanges(domain, servers)

  if !domain.ServersChanged {
    t.Error("Servers actually changed")
  }
}
