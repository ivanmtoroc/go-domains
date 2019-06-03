package handlers

import (
	"os"
	"testing"

	"github.com/ivanmtoroc/go-domains/models"
)

func TestMain(m *testing.M) {
	models.InitDB()
	models.CreateTables()

	sql := "DELETE domains; DELETE servers;"

	if _, err := models.GetDB().Exec(sql); err != nil {
		t := new(testing.T)
		t.Error("database clean error")
	}

	os.Exit(m.Run())
}
