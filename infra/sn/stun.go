package sn

import (
	"log"
	"net"

	"github.com/pion/stun"
)

func Stun() {
	// Listen on UDP port 3478
	addr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3478,
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on UDP port: %v", err)
	}
	defer conn.Close()

	log.Println("STUN server started on port 3478")

	buf := make([]byte, 1500)
	for {
		n, src, err := conn.ReadFrom(buf)
		if err != nil {
			log.Printf("Error reading from UDP: %v", err)
			continue
		}

		// Create a new STUN message
		msg := new(stun.Message)
		// Set the raw data to the message
		msg.Raw = buf[:n]
		// Decode the raw data into the STUN message
		if err := msg.Decode(); err != nil {
			log.Printf("Failed to decode STUN message: %v", err)
			continue
		}

		// Log STUN message
		log.Printf("Received STUN message from %s: %s", src, msg)

		// Build response message
		response := stun.MustBuild(
			stun.TransactionID,
			stun.BindingSuccess,
			&stun.XORMappedAddress{
				IP:   src.(*net.UDPAddr).IP,
				Port: src.(*net.UDPAddr).Port,
			},
		)

		// Send response
		if _, err := conn.WriteTo(response.Raw, src); err != nil {
			log.Printf("Failed to send STUN response: %v", err)
		}
	}
}
