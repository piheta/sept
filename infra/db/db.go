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
	db, err := sql.Open("sqlite3", "./sept.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Read schema from file
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Execute schema
	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database schema created successfully")
}
