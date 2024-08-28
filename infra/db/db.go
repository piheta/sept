package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./infra/db/sept.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if any of the tables in the schema already exist
	tables := []string{"users", "chats", "user_chats", "messages"}
	for _, table := range tables {
		query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s';", table)
		row := db.QueryRow(query)
		var tableName string
		err = row.Scan(&tableName)

		if err != nil && err != sql.ErrNoRows {
			log.Fatal(err)
		}

		if tableName == table {
			fmt.Println("Database is already initialized")
			return
		}
	}

	// Read schema from file
	schema, err := os.ReadFile("./infra/db/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Execute schema to create tables
	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database schema created successfully")
}
