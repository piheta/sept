package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/piheta/sept/models"

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

func AddMessage(chatID int, userID int, content string) {
	db, err := sql.Open("sqlite3", "./infra/db/sept.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
		INSERT INTO messages (chat_id, user_id, content, created_at) 
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
	`

	_, err = db.Exec(query, chatID, userID, content)
	if err != nil {
		log.Fatal(err)
	}
}

// GetMessagesByChatID retrieves all messages for a specific chat
func GetMessagesByChatID(chatID int) []models.Message {
	db, err := sql.Open("sqlite3", "./infra/db/sept.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
		SELECT id, chat_id, user_id, content, created_at
		FROM messages
		WHERE chat_id = ?
	`

	rows, err := db.Query(query, chatID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.ChatID, &msg.UserID, &msg.Content, &msg.CreatedAt); err != nil {
			log.Fatal(err)
		}
		messages = append(messages, msg)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return messages
}

func GetUser(userID int) models.User {
	db, err := sql.Open("sqlite3", "./infra/db/sept.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
        SELECT id, user_id, username, ip, avatar
        FROM users
        WHERE id = ?
    `

	row := db.QueryRow(query, userID)

	var user models.User
	err = row.Scan(&user.ID, &user.UserID, &user.Username, &user.Ip, &user.Avatar)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No user found")
		} else {
			log.Fatal(err)
		}
	}

	return user
}

func GetAllUsers() []models.User {
	db, err := sql.Open("sqlite3", "./infra/db/sept.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
		SELECT id, user_id, username, ip, avatar
		FROM users
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.UserID, &user.Username, &user.Ip, &user.Avatar); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}
