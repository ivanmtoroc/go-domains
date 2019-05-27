package models

import (
  "log"
  "database/sql"
  _ "github.com/lib/pq"
)

// Database connection object
var db *sql.DB

func InitDB() {
  // Address used to init connection
  const DB_ADDRESS = "postgresql://ivanmtoroc@localhost:26257/domains?sslmode=disable"
  var err error

  // Open connection with database
  db, err = sql.Open("postgres", DB_ADDRESS)
  if err != nil {
    log.Println("Database connection error")
    log.Fatalln(err)
  }
}

// Function that execute SQL stament to create tables into database
func CreateTables() {
  // Create 'domains' and 'servers' tables into database
  createDomainsTable()
  createServersTable()
}

// Function that return database connection object
func getDB() *sql.DB {
  return db
}
