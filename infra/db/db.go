package db

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/piheta/sept/models"
	"golang.org/x/crypto/argon2"
)

const (
	saltSize = 16
	keySize  = 32
	time     = 1
	memory   = 64 * 1024
	threads  = 4
)

var (
	db         *sql.DB
	dbMutex    sync.Mutex
	dbName     string
	saltName   string
	encryptKey []byte
)

func InitDb(id, password string) error {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	dbName = "./infra/db/" + id + ".db"
	saltName = "./infra/db/" + id + ".salt"

	// Create or load salt
	salt, err := loadOrCreateSalt()
	if err != nil {
		return fmt.Errorf("failed to load or create salt: %w", err)
	}

	// Derive encryption key
	encryptKey = deriveKey(password, salt)
	keyHex := hex.EncodeToString(encryptKey)

	db, err = sql.Open("sqlite3", fmt.Sprintf("%s?_pragma_key=x'%s'&_pragma_cipher_page_size=4096", dbName, keyHex))
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	// Test the connection
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Check if tables exist, create if not
	err = createTablesIfNotExist()
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	// Set permissions for the database file
	err = os.Chmod(dbName, 0600)
	if err != nil {
		return fmt.Errorf("failed to set database file permissions: %w", err)
	}

	fmt.Println("Database initialized successfully")
	return nil
}

func loadOrCreateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := os.Stat(saltName)
	if os.IsNotExist(err) {
		// Generate new salt
		_, err := rand.Read(salt)
		if err != nil {
			return nil, fmt.Errorf("error generating salt: %w", err)
		}
		// Save salt to file
		err = os.WriteFile(saltName, salt, 0600)
		if err != nil {
			return nil, fmt.Errorf("error saving salt: %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("error checking salt file: %w", err)
	} else {
		// Load existing salt
		salt, err = os.ReadFile(saltName)
		if err != nil {
			return nil, fmt.Errorf("error reading salt: %w", err)
		}
	}
	return salt, nil
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
	schema, err := os.ReadFile("./infra/db/schema.sql")
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

func deriveKey(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, time, memory, threads, keySize)
}

func AddMessage(chatID int, userID int, content string) error {
	query := `
		INSERT INTO messages (chat_id, user_id, content, created_at) 
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
	`
	_, err := db.Exec(query, chatID, userID, content)
	return err
}

func GetMessagesByChatID(chatID int) ([]models.Message, error) {
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
	query := `INSERT OR IGNORE INTO chats (name) VALUES (?)`
	chat, err := db.Exec(query, name)
	fmt.Println(chat)
	return err
}

func AddUser(user models.User) error {
	query := `INSERT OR IGNORE INTO users (user_id, username, ip, avatar) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, user.UserID, user.Username, user.Ip, user.Avatar)
	return err
}

func AddUserToChat(user_id string, chat_id int) error {
	query := `INSERT OR IGNORE INTO user_chats (user_id, chat_id) VALUES (?, ?)`
	_, err := db.Exec(query, user_id, chat_id)
	return err
}

func GetUser(userID int) (models.User, error) {
	query := `
        SELECT id, user_id, username, ip, avatar
        FROM users
        WHERE id = ?
    `
	var user models.User
	err := db.QueryRow(query, userID).Scan(&user.ID, &user.UserID, &user.Username, &user.Ip, &user.Avatar)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	query := `
		SELECT id, user_id, username, ip, avatar
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
		if err := rows.Scan(&user.ID, &user.UserID, &user.Username, &user.Ip, &user.Avatar); err != nil {
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
