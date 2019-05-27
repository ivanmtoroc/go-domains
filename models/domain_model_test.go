package models

import (
  "time"
  "testing"
)

func TestDomainSaveAndGet(t *testing.T) {
  domain_one := &Domain{0, "truora.com", false, "", "", "", "", false, false, time.Now()}
  domain_one.Save()
  domain_two := &Domain{0, "truora.com", false, "", "", "", "", false, true, time.Now()}
  domain_two.Save()

  database_domain := GetDomainByNameDB(domain_two.Name)
  if database_domain.ID != domain_two.ID {
    t.Error("Domain not save")
  }

  if count := GetDomainsCountDB(); count != 1 {
    t.Error("Domains count failed")
  }

  if domains := GetDomainsDB("", ""); domains[0].ID != domain_two.ID {
    t.Error("Domains order invalid")
  }
}
