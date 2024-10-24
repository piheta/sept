package controllers

import (
	"fmt"

	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/repos"
	"github.com/piheta/sept/backend/services"
)

type MessageController struct {
	message_repo *repos.MessageRepo
}

func NewMessageController(messageRepo *repos.MessageRepo) *MessageController {
	return &MessageController{
		message_repo: messageRepo,
	}
}

func (mc *MessageController) SendMessage(content string, chat_id string) ([]models.Message, error) {

	msg := models.Message{
		ChatID:  chat_id,
		UserID:  services.AuthedUser.ID,
		Content: content,
	}
	signedMsg, err := services.SignMessage(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to send message %w, ", err)
	}

	services.SendP2PMessage(signedMsg)
	err = mc.message_repo.AddMessage(signedMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}
	return mc.message_repo.GetMessagesByChatID(chat_id)
}

func (mc *MessageController) GetMessagesByChatID(chat_id string) ([]models.Message, error) {
	return mc.message_repo.GetMessagesByChatID(chat_id)
}
