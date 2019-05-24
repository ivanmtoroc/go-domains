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
  Title             string
  CreatedAt         time.Time
}

func (domain *Domain) Save() {
  GetDB().Create(domain)
}

func GetDomainByNameDB(name string) *Domain {
  domain := &Domain{}
  GetDB().Table("domains").Where("name = ?", name).Order("created_at desc").First(domain)
  if domain.Name != name {
		return nil
	}
  return domain
}

func GetDomainsDB() []*Domain {
  var domains []*Domain
  GetDB().Raw(`
    SELECT name, is_down, logo, title, created_at
    FROM domains AS d1
    WHERE created_at = (
        SELECT MAX(created_at) FROM domains AS d2 WHERE d1.name = d2.name
    )`,
  ).Scan(&domains)
  if len(domains) == 0 {
		return nil
	}
  return domains
}
