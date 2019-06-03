package models

import (
	"log"
	"time"
)

// Domain structure manipulate data of domains table into database
type Domain struct {
	ID               int64
	Name             string
	ServersChanged   bool
	SslGrade         string
	PreviousSslGrade string
	Logo             string
	Title            string
	IsDown           bool
	IsValid          bool
	CreatedAt        time.Time
}

// Save method save domain into database
func (domain *Domain) Save() error {
	// SQL instruction to insert new domain in domains table
	sql := `
  INSERT INTO domains
  (name, servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down, is_valid, created_at)
  VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  RETURNING id;
  `
	// Execute SQL instruction to insert new domain and get your ID
	if err := GetDB().QueryRow(
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
	).Scan(&domain.ID); err != nil {
		log.Println("domain insertion error")
		log.Println("- error: ", err)
		return err
	}
	return nil
}

// GetDomainByNameDB get most recent domain saved into database by name
func GetDomainByNameDB(name string) (*Domain, error) {
	// SQL query to get last saved domain
	sql := `
  SELECT *
  FROM domains
  WHERE name = $1
  ORDER BY created_at DESC
  LIMIT 1;
	`
	// Execute query
	rows, err := GetDB().Query(sql, name)
	if err != nil {
		log.Println("most recent domain query error")
		log.Println("- error: ", err)
		return nil, err
	}
	// Close rows when function end
	defer rows.Close()

	for rows.Next() {
		domain := &Domain{}
		// Save current row data in new domain object
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
			log.Println("most recent domains scan error")
			log.Println("- error: ", err)
			return nil, err
		}
		return domain, nil
	}
	return nil, nil
}

// GetDomainsDB get part (by offset and limit) of most recent and valid domains into database
func GetDomainsDB(offset, limit string) ([]*Domain, error) {
	var domains []*Domain
	// SQL query to get part of most recent and valid domains into database
	sql := `
  SELECT *
  FROM domains AS d1
  WHERE created_at = (
      SELECT MAX(created_at) FROM domains AS d2 WHERE d1.name = d2.name
  ) AND is_valid = TRUE
  ORDER BY name
  OFFSET $1
  LIMIT $2;
  `
	// Execute query
	rows, err := GetDB().Query(sql, offset, limit)
	if err != nil {
		log.Println("most recent domains query error")
		log.Println("- error: ", err)
		return nil, err
	}
	// Close rows when function end
	defer rows.Close()

	// Iterate rows of query result
	for rows.Next() {
		domain := &Domain{}
		// Save current row data in new domain object
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
			log.Println("most recent domains scan error")
			log.Println("- error: ", err)
			return nil, err
		}
		domains = append(domains, domain)
	}
	return domains, nil
}

// GetDomainsCountDB get the number of all most recent and valid domains into database
func GetDomainsCountDB() (int, error) {
	var count int
	// SQL query to get the number of all most recent and valid domains into database
	sql := `
  SELECT COUNT(id)
  FROM domains AS d1
  WHERE created_at = (
      SELECT MAX(created_at) FROM domains AS d2 WHERE d1.name = d2.name
  ) AND is_valid = TRUE
  `
	// Execute query
	if err := GetDB().QueryRow(sql).Scan(&count); err != nil {
		log.Println("domains table query error")
		log.Println("- error: ", err)
		return count, err
	}
	return count, nil
}

// createDomainsTable create the domains table into database if not exists
func createDomainsTable() {
	// SQL instructions to create new sequence and domains table if not exists
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
	// Execute SQL instructions
	if _, err := GetDB().Exec(sql); err != nil {
		log.Println("domains table creation error")
		log.Fatalln("- error: ", err)
	}
}
