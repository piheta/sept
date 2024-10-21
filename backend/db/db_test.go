package db_test

import (
	"testing"

	"github.com/piheta/sept/backend/db"
	"github.com/piheta/sept/backend/models"
)

var testUsr = models.User{
	ID:        "TEST",
	Username:  "TEST",
	Ip:        "TEST",
	Avatar:    "TEST",
	PublicKey: "TEST",
}

func TestInitDb(t *testing.T) {
	err := db.InitDb(testUsr)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
}

func TestDbExists(t *testing.T) {
	err := db.DbExists(testUsr.ID)
	if err != nil {
		t.Errorf("Expected DbExists to return no error for existing database, got: %v", err)
	}

	t.Cleanup(func() {
		db.RemoveDb(testUsr.ID)
	})
}
