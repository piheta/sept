package repos

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/piheta/sept/backend/models"
)

type ChatRepo struct {
	db *sql.DB
}

func NewChatRepo(db *sql.DB) *ChatRepo {
	return &ChatRepo{db: db}
}

func (cr *ChatRepo) AddChat(name string) error {
	chat_id := uuid.New()
	query := `INSERT OR IGNORE INTO chats (id, name) VALUES (?, ?)`
	_, err := cr.db.Exec(query, chat_id, name)
	return err
}

func (cr *ChatRepo) GetChatByName(chatName string) (models.Chat, error) {
	query := `
		SELECT id, name
		FROM chats
		WHERE name = ?
	`
	var chat models.Chat
	err := cr.db.QueryRow(query, chatName).Scan(&chat.ID, &chat.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Chat{}, fmt.Errorf("chat with name %s not found", chatName)
		}
		return models.Chat{}, err
	}
	return chat, nil
}
