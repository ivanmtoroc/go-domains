package models

import (
	"testing"
	"time"
)

func TestServerEqual(t *testing.T) {
	serverOne := &Server{
		ID:       0,
		Address:  "34.23.15.1",
		SslGrade: "A+",
		Country:  "CO",
		Owner:    "Truora",
		DomainID: 0,
	}

	serverTwo := &Server{
		ID:       1,
		Address:  "34.23.15.1",
		SslGrade: "A+",
		Country:  "CO",
		Owner:    "Truora",
		DomainID: 0,
	}

	serverThree := &Server{
		ID:       2,
		Address:  "34.22.15.1",
		SslGrade: "A+",
		Country:  "CO",
		Owner:    "Truora",
		DomainID: 0,
	}

	if equal := serverOne.Equal(serverTwo); !equal {
		t.Error("servers is actually equals")
	}

	if equal := serverOne.Equal(serverThree); equal {
		t.Error("servers is actually differents")
	}
}

func TestServerSaveAndGetServers(t *testing.T) {
	domain := &Domain{
		Name:      "truora.com",
		IsValid:   true,
		CreatedAt: time.Now(),
	}

	server := &Server{
		Address:  "35.11.51.12",
		SslGrade: "C",
		Country:  "CO",
		Owner:    "Azure",
	}

	otherServer := &Server{
		Address:  "34.12.52.15",
		SslGrade: "B",
		Country:  "US",
		Owner:    "Amazon",
	}

	domain.Save()
	server.Save(domain)
	otherServer.Save(domain)

	servers, _ := GetServersByDomainDB(domain)

	if !server.Equal(servers[1]) {
		t.Error("server one not saved")
	} else if !otherServer.Equal(servers[0]) {
		t.Error("server two not saved")
	}
}
