package controllers

import (
	"fmt"

	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/services"
)

type SignalingController struct {
	sn_con_handler *services.SnConnection
}

func NewSignalingController(sn_con_handler *services.SnConnection) *SignalingController {
	return &SignalingController{
		sn_con_handler: sn_con_handler,
	}
}

//
// SIGNALING
//

func (sc *SignalingController) SearchDht(username string) (models.User, error) {
	userChan, err := sc.sn_con_handler.UserSearchRequest(username)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to search DHT: %v", err)
	}

	user := <-userChan
	return user, nil
}

func (sc *SignalingController) SendUserAddRequest(destIp string) {
	sc.sn_con_handler.SendSDPOffer(destIp)
}
func (sc *SignalingController) SendUserAddResponse(destIp string) {
	sc.sn_con_handler.SendSDPOffer(destIp)
}

func (sc *SignalingController) SendOffer(destIp string) {
	sc.sn_con_handler.SendSDPOffer(destIp)
}
