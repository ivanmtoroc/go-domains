package models

import (
	"database/sql"
	"log"

	// Dependency to start connection to database
	_ "github.com/lib/pq"
)

// db is the database connection object
var db *sql.DB

// InitDB function init the connection to database
func InitDB() {
	// DbAddress used to init connection
	const DbAddress = "postgresql://ivanmtoroc@localhost:26257/domains?sslmode=disable"
	var err error

	// Open connection
	db, err = sql.Open("postgres", DbAddress)
	if err != nil {
		log.Println("database connection error")
		log.Fatalln("- error: ", err)
	}
}

// CreateTables execute SQL staments to create tables into database
func CreateTables() {
	// Create 'domains' and 'servers' tables into database
	createDomainsTable()
	createServersTable()
}

// GetDB return the database connection object
func GetDB() *sql.DB {
	return db
}
