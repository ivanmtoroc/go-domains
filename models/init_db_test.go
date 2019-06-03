package models

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitDB()
	CreateTables()

	sql := "DELETE domains; DELETE servers;"

	if _, err := GetDB().Exec(sql); err != nil {
		t := new(testing.T)
		t.Error("database clean error")
	}

	os.Exit(m.Run())
}

func TestInitDBAndGetDB(t *testing.T) {
	if GetDB() == nil {
		t.Error("database init failed")
	}
	if err := GetDB().Ping(); err != nil {
		t.Error("database connection is down")
	}
}

func TestCreateTables(t *testing.T) {
	var count int
	var sql string

	sql = "SELECT COUNT(id) FROM domains;"
	if err := GetDB().QueryRow(sql).Scan(&count); err != nil {
		t.Error("domains table is not created")
	}

	sql = "SELECT COUNT(id) FROM servers;"
	if err := GetDB().QueryRow(sql).Scan(&count); err != nil {
		t.Error("servers table is not created")
	}
}
