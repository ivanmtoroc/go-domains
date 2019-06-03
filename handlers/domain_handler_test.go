package handlers

import (
	"testing"
	"time"

	"github.com/ivanmtoroc/go-domains/models"
)

func TestVerifyServersChanges(t *testing.T) {
	now := time.Now()

	previousDomain := &models.Domain{
		Name:      "truora.com",
		SslGrade:  "A",
		IsValid:   true,
		CreatedAt: now,
	}
	previousDomain.Save()

	previousServerOne := &models.Server{
		Address:  "32.32.32.32",
		SslGrade: "A",
		Country:  "CO",
		Owner:    "Amazon",
	}
	previousServerOne.Save(previousDomain)

	previousServerTwo := &models.Server{
		Address:  "33.33.33.33",
		SslGrade: "B",
		Country:  "US",
		Owner:    "Google",
	}
	previousServerTwo.Save(previousDomain)

	domain := &models.Domain{
		Name:      "truora.com",
		SslGrade:  "A",
		IsValid:   true,
		CreatedAt: now.Add(time.Hour),
	}

	var servers []*models.Server
	serverOne := &models.Server{
		Address:  "32.32.32.32",
		SslGrade: "A",
		Country:  "CO",
		Owner:    "Amazon",
	}
	servers = append(servers, serverOne)

	serverTwo := &models.Server{
		Address:  "33.33.33.32",
		SslGrade: "B",
		Country:  "US",
		Owner:    "Google",
	}
	servers = append(servers, serverTwo)

	verifyServersChanges(domain, servers)

	if !domain.ServersChanged {
		t.Error("servers actually changed")
	}
}
