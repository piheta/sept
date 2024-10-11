package sn

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/pion/webrtc/v4"

	"github.com/piheta/sept/backend/models"
)

var (
	offers      []webrtc.SessionDescription
	offersLock  sync.Mutex
	clients     = make(map[*websocket.Conn]string)
	clientsLock sync.Mutex
)

func Signaling(c *websocket.Conn) {
	defer func() {
		clientsLock.Lock()
		delete(clients, c)
		clientsLock.Unlock()
		c.Close()
	}()

	clientsLock.Lock()
	clients[c] = c.RemoteAddr().String()
	clientsLock.Unlock()

	// Send the list of connected IPs (excluding the client's own IP) to the client
	clientsLock.Lock()
	var ips []string
	// clientIP := c.RemoteAddr().String() // Get the current client's IP
	for _, ip := range clients {
		// if ip != clientIP { // Exclude the client's own IP:port
		ips = append(ips, ip)
		// }
	}
	clientsLock.Unlock()

	if err := c.WriteJSON(ips); err != nil {
		log.Printf("Failed to write message: %v", err)
		return
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		log.Printf("Received message: %s from %s", message, c.RemoteAddr())

		var cr models.ConnectionRequest
		if err := json.Unmarshal(message, &cr); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		// Set the source IP
		cr.SrcIP = func(s string) *string { return &s }(c.RemoteAddr().String())

		// Marshal the updated connection request
		message, err = json.Marshal(cr)
		if err != nil {
			log.Printf("Failed to marshal message: %v", err)
			continue
		}

		if err := send(cr.DestIP, message); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}

func send(remoteAddr string, message []byte) error {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	for client, addr := range clients {
		if addr == remoteAddr {
			if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func SuperNode(sigport int) {
	// Start STUN server in a separate goroutine
	go Stun()

	// Start Fiber app
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(Signaling))

	log.Println("Signaling server started on port ", sigport)
	if err := app.Listen(":" + strconv.Itoa(sigport)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
