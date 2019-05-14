package models

import (
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB() {
  const ADDRESS = "postgresql://ivanmtoroc@localhost:26257/domains?sslmode=disable"

  db_connection, err := gorm.Open("postgres", ADDRESS)
  if err != nil {
    log.Fatal(err)
  }

  db = db_connection
  db.AutoMigrate(&Server{})
  db.AutoMigrate(&Domain{})
}

func GetDB() *gorm.DB {
  return db
}
