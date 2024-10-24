package services

import (
	"encoding/json"
	"fmt"

	"github.com/piheta/sept/backend/models"
	"github.com/pion/webrtc/v4"
)

var d *webrtc.DataChannel

type MessagingHandler struct {
	d *webrtc.DataChannel
}

func NewMessagingHandler(dataChannel *webrtc.DataChannel) *MessagingHandler {
	d = dataChannel
	return &MessagingHandler{d: d}
}

func SendP2PMessage(message models.Message) {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Failed to marshall p2p message, ", err)
	}

	if err := d.Send(messageBytes); err != nil {
		fmt.Println("Failed to send p2p message, ", err)
		panic(err)
	}
}
