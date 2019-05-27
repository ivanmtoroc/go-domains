package models

import (
  "os"
  "testing"
)

func TestMain(m *testing.M) {
  InitDB()
  CreateTables()
  if _, err := getDB().Exec(
    "DELETE domains; DELETE servers;",
  ); err != nil {
    t := new(testing.T)
    t.Error("Cleaning the database error")
  }
  os.Exit(m.Run())
}

func TestInitDBAndGetDB(t *testing.T) {
  db := getDB()

  if db == nil {
    t.Error("Database init failed")
  }
  if err := db.Ping(); err != nil {
    t.Error("Database connection is down")
  }
}

func TestCreateTables(t *testing.T) {
  var count int
  var sql string

  sql = "SELECT COUNT(id) FROM domains;"
  if err := getDB().QueryRow(sql).Scan(&count); err != nil {
    t.Error("Domains table is not created")
  }

  sql = "SELECT COUNT(id) FROM servers;"
  if err := getDB().QueryRow(sql).Scan(&count); err != nil {
    t.Error("Servers table is not created")
  }
}
