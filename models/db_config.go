package models

import (
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB() {
  const DB_ADDRESS = "postgresql://ivanmtoroc@localhost:26257/domains?sslmode=disable"
  var err error

  // Init connection with database
  db, err = gorm.Open("postgres", DB_ADDRESS)
  if err != nil {
    log.Println("Database initialization error:")
    log.Fatalln(err)
  }

  // Migrate models
  db.AutoMigrate(&Server{}, &Domain{})
}

func GetDB() *gorm.DB {
  return db
}
