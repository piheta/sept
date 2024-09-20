package repos

import (
	"database/sql"

	"github.com/piheta/sept/backend/models"
)

type MessageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (mr *MessageRepo) AddMessage(chatID string, userID string, content string) error {
	query := `
		INSERT INTO messages (chat_id, user_id, content, created_at) 
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
	`
	_, err := mr.db.Exec(query, chatID, userID, content)
	return err
}

func (mr *MessageRepo) GetMessagesByChatID(chatID string) ([]models.Message, error) {
	query := `
		SELECT id, chat_id, user_id, content, created_at
		FROM messages
		WHERE chat_id = ?
	`
	rows, err := mr.db.Query(query, chatID)
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
