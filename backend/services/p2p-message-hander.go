package services

import (
	"fmt"

	"github.com/pion/webrtc/v4"
)

func MessagingHandler(d *webrtc.DataChannel) {
	for {
		// Send the message as text
		var messageSend string
		_, err := fmt.Scanln(&messageSend)
		if err != nil {
			panic(err)
		}

		if err := d.SendText(messageSend); err != nil {
			panic(err)
		}
	}
}
