package models

import (
  "log"
  "time"
)

// Struct used to manipulate database domain objects
type Domain struct {
  ID                int64
  Name              string
  ServersChanged    bool
  SslGrade          string
  PreviousSslGrade  string
  Logo              string
  Title             string
  IsDown            bool
  IsValid           bool
  CreatedAt         time.Time
}

// Domain method to save a domain into database
func (domain *Domain) Save() {
  sql := `
  INSERT INTO domains
  (name, servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down, is_valid, created_at)
  VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9);
  `
  // Execute insertion
  if _, err := getDB().Exec(
      sql,
      domain.Name,
      domain.ServersChanged,
      domain.SslGrade,
      domain.PreviousSslGrade,
      domain.Logo,
      domain.Title,
      domain.IsDown,
      domain.IsValid,
      domain.CreatedAt,
    ); err != nil {
    log.Println("Domain insertion error")
    log.Fatalln(err)
  }
}

// Function to create domains table into database
func createDomainsTable() {
  sql := `
  CREATE SEQUENCE IF NOT EXISTS domains_seq;

  CREATE TABLE IF NOT EXISTS domains (
    id INT PRIMARY KEY DEFAULT nextval('domains_seq'),
    name STRING NOT NULL,
    servers_changed STRING NOT NULL,
    ssl_grade STRING NOT NULL,
    previous_ssl_grade STRING NOT NULL,
    logo STRING NOT NULL,
    title STRING NOT NULL,
    is_down BOOLEAN NOT NULL,
    is_valid BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL
  );
  `
  // Execute statement
  if _, err := getDB().Exec(sql); err != nil {
    log.Println("Domains table creation error")
    log.Fatalln(err)
  }
}

// Function to get most recent domain from database by name
func GetDomainByNameDB(name string) *Domain {
  sql := `
  SELECT *
  FROM domains
  WHERE name = $1
  ORDER BY created_at DESC
  LIMIT 1;
  `
  // Execute query
  rows, err := getDB().Query(sql, name)
  if err != nil {
      log.Println("Domain query error")
      log.Fatalln(err)
  }
  // Defer close
  defer rows.Close()

  // Iterate rows of query result
  for rows.Next() {
      domain := &Domain{}
      if err := rows.Scan(
          &domain.ID,
          &domain.Name,
          &domain.ServersChanged,
          &domain.SslGrade,
          &domain.PreviousSslGrade,
          &domain.Logo,
          &domain.Title,
          &domain.IsDown,
          &domain.IsValid,
          &domain.CreatedAt,
        ); err != nil {
          log.Println("Get last domain from database error")
          log.Fatalln(err)
      }
      return domain
  }
  return nil
}

// Function to get all domains
func GetDomainsDB() []*Domain {
  var domains []*Domain
  sql := `
  SELECT *
  FROM domains AS d1
  WHERE created_at = (
      SELECT MAX(created_at) FROM domains AS d2 WHERE d1.name = d2.name
  ) AND is_valid = TRUE
  `
  // Execute query
  rows, err := getDB().Query(sql)
  if err != nil {
      log.Println("Domains query error")
      log.Fatalln(err)
  }
  // Defer close
  defer rows.Close()

  // Iterate rows of query result
  for rows.Next() {
      domain := &Domain{}
      if err := rows.Scan(
          &domain.ID,
          &domain.Name,
          &domain.ServersChanged,
          &domain.SslGrade,
          &domain.PreviousSslGrade,
          &domain.Logo,
          &domain.Title,
          &domain.IsDown,
          &domain.IsValid,
          &domain.CreatedAt,
        ); err != nil {
          log.Println("Get domain from database error")
          log.Fatalln(err)
      }
      domains = append(domains, domain)
  }
  // Validate result
  if len(domains) == 0 {
    return nil
  }
  return domains
}
