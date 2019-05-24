package models

import (
  "log"
  "time"
)

// Struct used to manipulate database server objects
type Server struct {
  ID          int64
  Address     string
  SslGrade    string
  Country     string
  Owner       string
  DomainName  string
  CreatedAt   time.Time
}

// Function to save a server into database
func (server *Server) Save() {
  sql := `
  INSERT INTO servers
  (address, ssl_grade, country, owner, domain_name, created_at)
  VALUES
  ($1, $2, $3, $4, $5, $6);
  `
  // Execute insertion
  if _, err := getDB().Exec(
      sql,
      server.Address,
      server.SslGrade,
      server.Country,
      server.Owner,
      server.DomainName,
      server.CreatedAt,
    ); err != nil {
    log.Println("Server insertion error")
    log.Fatalln(err)
  }
}

// Function to determine if two server are equal
func (server *Server) Equal(other_server *Server) bool {
  if server.Address != other_server.Address {
    return false
  } else if server.SslGrade != other_server.SslGrade {
    return false
  } else if server.Country != other_server.Country {
    return false
  } else if server.Owner != other_server.Owner {
    return false
  } else if server.DomainName != other_server.DomainName {
    return false
  }
  return true
}

// Function to create domains table into database
func createServersTable() {
  sql := `
  CREATE SEQUENCE IF NOT EXISTS servers_seq;

  CREATE TABLE IF NOT EXISTS servers (
    id INT PRIMARY KEY DEFAULT nextval('servers_seq'),
    address STRING NOT NULL,
    ssl_grade STRING NOT NULL,
    country STRING NOT NULL,
    owner STRING NOT NULL,
    domain_name STRING NOT NULL,
    created_at TIMESTAMP NOT NULL
  );
  `
  // Execute statement
  if _, err := getDB().Exec(sql); err != nil {
    log.Println("Servers table creation error")
    log.Fatalln(err)
  }
}

// Function to get servers of domain
func GetServersByDomainDB(domain *Domain) []*Server {
  var servers []*Server
  sql := `
  SELECT *
  FROM servers
  WHERE domain_name = $1 AND created_at = $2
  ORDER BY address;
  `
  // Execute query
  rows, err := getDB().Query(sql, domain.Name, domain.CreatedAt)
  if err != nil {
      log.Println("Servers query error")
      log.Fatalln(err)
  }
  // Defer close
  defer rows.Close()

  // Iterate rows of query result
  for rows.Next() {
      server := &Server{}
      if err := rows.Scan(
          &server.ID,
          &server.Address,
          &server.SslGrade,
          &server.Country,
          &server.Owner,
          &server.DomainName,
          &server.CreatedAt,
        ); err != nil {
          log.Println("Get server from database error")
          log.Fatalln(err)
      }
      servers = append(servers, server)
  }
  // Validate result
  if len(servers) == 0 {
    return nil
  }
  return servers
}
