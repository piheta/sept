package db

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/repos"
)

//go:embed schema.sql
var schemaSQL []byte

var (
	DB      *sql.DB
	dbMutex sync.Mutex
	dbName  string

	SEPT_DATA string
)

func InitDb(user models.User) error {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	dbName = SEPT_DATA + "/" + user.ID + ".db"

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	DB = db

	err = createTablesIfNotExist(user)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	err = os.Chmod(dbName, 0600)
	if err != nil {
		return fmt.Errorf("failed to set database file permissions: %w", err)
	}

	return nil
}

func createTablesIfNotExist(user models.User) error {
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

	// Execute schema to create tables
	_, err := DB.Exec(string(schemaSQL))
	if err != nil {
		return err
	}

	err = populateTables(user)
	if err != nil {
		return err
	}

	return nil
}

func populateTables(user models.User) error {
	// create repos, use them to fill database
	user_repo := repos.NewUserRepo(DB)
	chat_repo := repos.NewChatRepo(DB)
	userchat_repo := repos.NewUserchatRepo(DB)

	// Add the user to the database
	if err := user_repo.AddUser(user); err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}

	// Add a new chat for the user
	if err := chat_repo.AddChat(user.Username, user.Avatar); err != nil {
		return fmt.Errorf("failed to add chat: %w", err)
	}

	// Retrieve the created chat by username
	chat, err := chat_repo.GetChatByName(user.Username)
	if err != nil {
		return fmt.Errorf("failed to get chat: %w", err)
	}
	fmt.Println("chat:", chat)
	// Add the user to the chat
	if err := userchat_repo.AddUserToChat(user.ID, chat.ID); err != nil {
		return fmt.Errorf("failed to add user to chat: %w", err)
	}

	return nil
}

func DbExists(userID string) error {
	dbPath := SEPT_DATA + "/" + userID + ".db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		return err // File does not exist
	} else if err != nil {
		return fmt.Errorf("error checking database existence: %w", err)
	}

	return nil // File exists
}

func RemoveDb(userID string) error {
	// currently only used for testing
	// Todo also check password before deletion
	dbPath := SEPT_DATA + "/" + userID + ".db"
	err := os.Remove(dbPath)
	if err != nil {
		return fmt.Errorf("error removing database: %w", err)
	}

	return nil
}

func Search(searchString string) ([]string, error) {
	var results []string

	// Search in chats table (search by chat name)
	chatsQuery := `SELECT name FROM chats WHERE name LIKE ?`
	rows, err := DB.Query(chatsQuery, "%"+searchString+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to search chats: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var chatName string
		if err := rows.Scan(&chatName); err != nil {
			return nil, fmt.Errorf("failed to scan chat: %w", err)
		}
		results = append(results, fmt.Sprintf("Chat: %s", chatName))
	}

	if len(results) > 0 {
		return results, nil
	}

	// Search in messages table (search by message content)
	messagesQuery := `SELECT content FROM messages WHERE content LIKE ?`
	rows, err = DB.Query(messagesQuery, "%"+searchString+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to search messages: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var messageContent string
		if err := rows.Scan(&messageContent); err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		results = append(results, messageContent)
	}

	return results, nil
}
