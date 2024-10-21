package services

import (
	"github.com/pion/webrtc/v4"
)

func MessagingHandler(d *webrtc.DataChannel) {
	for {
		// Send the message as text
		var messageSend string

		if err := d.SendText(messageSend); err != nil {
			panic(err)
		}
	}
}
