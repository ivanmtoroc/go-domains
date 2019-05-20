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
  GetDB().Table("domains").Select("name, max(is_down) as is_down, max(logo) as logo, max(title) as title, max(created_at) as created_at").Group("name").Scan(&domains)
  if len(domains) == 0 {
		return nil
	}
  return domains
}
