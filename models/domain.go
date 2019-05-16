package models

import (
  "time"
)

type Domain struct {
  Name              string
  ServersChanged    bool
  SslGrade          string
  PreviousSslGrade  string
  Logo              string
  IsDown            bool
  UpdatedAt         time.Time
}

func (domain *Domain) Create() {
  GetDB().Create(domain)
}

func GetDomainDB(name string) *Domain {
  domain := &Domain{}
  GetDB().Table("domains").Where("name = ?", name).Order("updated_at desc").First(domain)
  if domain.Name != name {
		return nil
	}
  return domain
}
