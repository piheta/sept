package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/piheta/sept/backend/models"
)

var (
	DB      *sql.DB
	dbMutex sync.Mutex
	dbName  string
)

func InitDb(id string) error {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	dbName = "./sept_data/" + id + ".db"

	var err error
	DB, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	err = createTablesIfNotExist()
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	err = os.Chmod(dbName, 0600)
	if err != nil {
		return fmt.Errorf("failed to set database file permissions: %w", err)
	}

	fmt.Println("Database initialized successfully")
	return nil
}

func createTablesIfNotExist() error {
	// Check if any of the tables in the schema already exist
	tables := []string{"users", "chats", "user_chats", "messages"}
	for _, table := range tables {
		query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s';", table)
		row := DB.QueryRow(query)
		var tableName string
		err := row.Scan(&tableName)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if tableName == table {
			return nil // Tables already exist
		}
	}

	// Read schema from file
	schema, err := os.ReadFile("./backend/db/schema.sql")
	if err != nil {
		return err
	}

	// Execute schema to create tables
	_, err = DB.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}

func DbExists(userID string) error {
	dbPath := "./sept_data/" + userID + ".db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		return err // File does not exist
	} else if err != nil {
		return fmt.Errorf("error checking database existence: %w", err)
	}

	return nil // File exists
}
