package handlers

import (
  "os"
  "testing"
  "go-domains/models"
)

func TestMain(m *testing.M) {
  models.InitDB()
  models.CreateTables()
  if _, err := models.GetDB().Exec(
    "DELETE domains; DELETE servers;",
  ); err != nil {
    t := new(testing.T)
    t.Error("Cleaning the database error")
  }
  os.Exit(m.Run())
}
