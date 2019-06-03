package models

import (
	"log"
)

// Server structure manipulate data of servers table into database
type Server struct {
	ID       int64
	Address  string
	SslGrade string
	Country  string
	Owner    string
	DomainID int64
}

// Save method save server into database
func (server *Server) Save(domain *Domain) error {
	// SQL instruction to insert new server in servers table
	sql := `
  INSERT INTO servers
  (address, ssl_grade, country, owner, domain_id)
  VALUES
  ($1, $2, $3, $4, $5)
  RETURNING id;
  `
	// Execute SQL instruction to insert new server and get your ID
	if err := GetDB().QueryRow(
		sql,
		server.Address,
		server.SslGrade,
		server.Country,
		server.Owner,
		domain.ID,
	).Scan(&server.ID); err != nil {
		log.Println("server insertion error")
		log.Println("- error: ", err)
		return err
	}
	return nil
}

// Equal determine if two server are equal
func (server *Server) Equal(otherServer *Server) bool {
	if server.Address != otherServer.Address {
		return false
	} else if server.SslGrade != otherServer.SslGrade {
		return false
	} else if server.Country != otherServer.Country {
		return false
	} else if server.Owner != otherServer.Owner {
		return false
	}
	return true
}

// GetServersByDomainDB get all servers of a domain saved into database
func GetServersByDomainDB(domain *Domain) ([]*Server, error) {
	var servers []*Server
	// SQL query to get all servers of a domain into database
	sql := `
  SELECT *
  FROM servers
  WHERE domain_id = $1
  ORDER BY address;
  `
	// Execute query
	rows, err := GetDB().Query(sql, domain.ID)
	if err != nil {
		log.Println("servers query error")
		log.Println("- error: ", err)
		return nil, err
	}
	// Close rows when function end
	defer rows.Close()

	// Iterate rows of query result
	for rows.Next() {
		server := &Server{}
		// Save current row data in new server object
		if err := rows.Scan(
			&server.ID,
			&server.Address,
			&server.SslGrade,
			&server.Country,
			&server.Owner,
			&server.DomainID,
		); err != nil {
			log.Println("server scan error")
			log.Println("- error: ", err)
			return nil, err
		}
		servers = append(servers, server)
	}
	return servers, nil
}

// ServersChanges determine if two collection of servers are equal
func ServersChanges(servers, previousServers []*Server) bool {
	// If the number of servers is different then servers changed
	if len(servers) != len(previousServers) {
		return true
	}
	// Compare differences server by server
	for i, server := range servers {
		if !server.Equal(previousServers[i]) {
			return true
		}
	}
	return false
}

// createServersTable create the servers table into database if not exists
func createServersTable() {
	// SQL instructions to create new sequence and servers table if not exists
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
	// Execute SQL instructions
	if _, err := GetDB().Exec(sql); err != nil {
		log.Println("servers table creation error")
		log.Fatalln("- error: ", err)
	}
}
