package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/piheta/sept/backend/models"
)

var (
	db      *sql.DB
	dbMutex sync.Mutex
	dbName  string
)

func InitDb(id string) error {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	dbName = "./sept_data/" + id + ".db"

	var err error
	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	err = db.Ping()
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
		row := db.QueryRow(query)
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
	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}

func AddMessage(chatID string, userID string, content string) error {
	query := `
		INSERT INTO messages (chat_id, user_id, content, created_at) 
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
	`
	_, err := db.Exec(query, chatID, userID, content)
	return err
}

func GetMessagesByChatID(chatID string) ([]models.Message, error) {
	query := `
		SELECT id, chat_id, user_id, content, created_at
		FROM messages
		WHERE chat_id = ?
	`
	rows, err := db.Query(query, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.ChatID, &msg.UserID, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}

func AddChat(name string) error {
	chat_id := uuid.New()
	query := `INSERT OR IGNORE INTO chats (id, name) VALUES (?, ?)`
	chat, err := db.Exec(query, chat_id, name)
	fmt.Println(chat)
	return err
}

func AddUser(user models.User) error {
	query := `INSERT OR IGNORE INTO users (id, username, ip, avatar) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, user.ID, user.Username, user.Ip, user.Avatar)
	return err
}

func AddUserToChat(user_id string, chat_id string) error {
	query := `INSERT OR IGNORE INTO user_chats (user_id, chat_id) VALUES (?, ?)`
	_, err := db.Exec(query, user_id, chat_id)
	return err
}

func GetUser(userID string) (models.User, error) {
	query := `
        SELECT id, user_id, username, ip, avatar
        FROM users
        WHERE user_id = ?
    `
	var user models.User
	err := db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Ip, &user.Avatar)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	query := `
		SELECT id, username, ip, avatar, public_key
		FROM users
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Ip, &user.Avatar, &user.PublicKey); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func GetChatByName(chatName string) (models.Chat, error) {
	query := `
		SELECT id, name
		FROM chats
		WHERE name = ?
	`
	var chat models.Chat
	err := db.QueryRow(query, chatName).Scan(&chat.ID, &chat.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Chat{}, fmt.Errorf("chat with name %s not found", chatName)
		}
		return models.Chat{}, err
	}
	return chat, nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if db != nil {
		return db.Close()
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
