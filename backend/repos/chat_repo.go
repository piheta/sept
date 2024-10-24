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

func (cr *ChatRepo) GetChats() ([]models.Chat, error) {
	query := `
		SELECT id, name, avatar
		FROM chats
	`
	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []models.Chat
	for rows.Next() {
		var chat models.Chat
		err := rows.Scan(&chat.ID, &chat.Name, &chat.Avatar)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return chats, nil
}

func (cr *ChatRepo) AddChat(name string, avatar string) error {
	chat_id := uuid.New()
	query := `INSERT OR IGNORE INTO chats (id, name, avatar) VALUES (?, ?, ?)`
	_, err := cr.db.Exec(query, chat_id, name, avatar)
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

func (cr *ChatRepo) SetDB(db *sql.DB) {
	cr.db = db
}
