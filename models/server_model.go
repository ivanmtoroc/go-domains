package models

import (
  "log"
)

// Struct used to manipulate database server objects
type Server struct {
  ID          int64
  Address     string
  SslGrade    string
  Country     string
  Owner       string
  DomainID    int64
}

// Function to save a server into database
func (server *Server) Save(domain *Domain) {
  sql := `
  INSERT INTO servers
  (address, ssl_grade, country, owner, domain_id)
  VALUES
  ($1, $2, $3, $4, $5)
  RETURNING id;
  `
  // Execute insertion
  if err := GetDB().QueryRow(
      sql,
      server.Address,
      server.SslGrade,
      server.Country,
      server.Owner,
      domain.ID,
    ).Scan(&server.ID); err != nil {
    log.Println("server insertion error into database")
    log.Fatalln("- error: ", err)
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
    domain_id INT NOT NULL,
    FOREIGN KEY (domain_id) REFERENCES domains (id)
  );
  `
  // Execute statement
  if _, err := GetDB().Exec(sql); err != nil {
    log.Println("servers table creation error into database")
    log.Fatalln("- error: ", err)
  }
}

// Function to get servers of domain
func GetServersByDomainDB(domain *Domain) []*Server {
  var servers []*Server
  sql := `
  SELECT *
  FROM servers
  WHERE domain_id = $1
  ORDER BY address;
  `
  // Execute query
  rows, err := GetDB().Query(sql, domain.ID)
  if err != nil {
    log.Println("servers table query error")
    log.Fatalln("- error: ", err)
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
          &server.DomainID,
        ); err != nil {
          log.Println("error while get a server from database")
          log.Fatalln("- error: ", err)
      }
      servers = append(servers, server)
  }
  // Validate result
  if len(servers) == 0 {
    return nil
  }
  return servers
}

// Function to compare two [] of servers
func ServersChanges(servers, previous_servers []*Server) bool {
  // If number of servers is different then servers changed
  if len(servers) != len(previous_servers) {
    return true
  }

  // Compare all servers
  for i, server := range servers {
    if !server.Equal(previous_servers[i]) {
      return true
    }
  }

  return false
}
