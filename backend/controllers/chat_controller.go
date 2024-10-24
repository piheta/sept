package controllers

import (
	"fmt"

	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/repos"
)

type ChatController struct {
	chat_repo *repos.ChatRepo
}

func NewChatController(chatRepo *repos.ChatRepo) *ChatController {
	return &ChatController{
		chat_repo: chatRepo,
	}
}

func (cc *ChatController) GetChats() ([]models.Chat, error) {
	chats, err := cc.chat_repo.GetChats()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve chats: %w", err)
	}
	return chats, nil
}
