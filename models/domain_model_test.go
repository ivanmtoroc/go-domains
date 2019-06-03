package models

import (
	"testing"
	"time"
)

func TestDomainSaveAndGet(t *testing.T) {
	domainOne := &Domain{
		Name:      "truora.com",
		CreatedAt: time.Now(),
	}
	domainOne.Save()

	domainTwo := &Domain{
		Name:      "truora.com",
		IsValid:   true,
		CreatedAt: time.Now(),
	}
	domainTwo.Save()

	databaseDomain, _ := GetDomainByNameDB(domainTwo.Name)
	if databaseDomain.ID != domainTwo.ID {
		t.Error("Domain not save")
	}

	if count, _ := GetDomainsCountDB(); count != 1 {
		t.Error("Domains count failed")
	}

	if domains, _ := GetDomainsDB("0", "10"); domains[0].ID != domainTwo.ID {
		t.Error("Domains order invalid")
	}
}
