package models

import (
  "errors"
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

func GetDomainDB(name string) (*Domain, error) {
  domain := &Domain{}
  GetDB().Table("domains").Where("name = ?", name).First(domain)
  if domain.Name != name {
		return nil, errors.New("Domain not found.")
	}
  return domain, nil
}
