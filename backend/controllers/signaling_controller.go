package controllers

import (
	"fmt"

	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/services"
)

type SignalingController struct {
}

func NewSignalingController() *SignalingController {
	return &SignalingController{}
}

//
// SIGNALING
//

func (sc *SignalingController) SearchDht(username string) (models.User, error) {
	userChan, err := services.UserSearchRequest(username)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to search DHT: %v", err)
	}

	user := <-userChan
	return user, nil
}

func (sc *SignalingController) SendOffer(destIp string) {
	services.SendSDPOffer(destIp)
}
