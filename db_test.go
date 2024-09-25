package main

import (
	"testing"

	"github.com/piheta/sept/backend/db"
)

var testDb = "DB-TEST"

// TestInitDbAndDbExists tests the InitDb and DbExists functions.
func TestInitDbAndDbExists(t *testing.T) {
	err := db.InitDb(testDb)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

}

// Test DbExists with the same test ID
func TestDbExists(t *testing.T) {
	err := db.DbExists(testDb)
	if err != nil {
		t.Errorf("Expected DbExists to return no error for existing database, got: %v", err)
	}

	t.Cleanup(func() {
		db.RemoveDb(testDb)
	})
}
