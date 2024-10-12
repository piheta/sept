package sn

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/pion/webrtc/v4"

	"github.com/piheta/sept/backend/models"
)

var (
	offers      []webrtc.SessionDescription
	offersLock  sync.Mutex
	clients     = make(map[*websocket.Conn]string)
	clientsLock sync.Mutex

	userDht     = make(map[string]models.DhtUser)
	userDhtLock sync.RWMutex
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

	//* Listen for messages from clients
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		var clientMessage models.SigMsg
		if err := json.Unmarshal(message, &clientMessage); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		switch clientMessage.Type {
		case models.Announce:
			fmt.Println("Announce, ", clientMessage.Type)
			onAnnounceMsg(clientMessage, c.RemoteAddr().String())
		case models.UserSearch:
			fmt.Println("Usearsearch, ", clientMessage.Type)
			onSearchMsg(clientMessage, c.RemoteAddr().String())
		case models.Connection:
			fmt.Println("Connection, ", clientMessage.Type)
			onConnectionMsg(clientMessage, c.RemoteAddr().String())
		}
	}
}

func onSearchMsg(msg models.SigMsg, senderAddr string) {
	dataBytes, err := json.Marshal(msg.Data)
	if err != nil {
		log.Printf("Failed to marshal Data: %v", err)
		return
	}

	var srchreq models.UserSearchRequest
	if err := json.Unmarshal(dataBytes, &srchreq); err != nil {
		log.Printf("Failed to unmarshal AnnounceRequest: %v", err)
		return
	}

	dhtUser := userDht[srchreq.Username]

	sigMsg := models.SigMsg{
		Type: 1,
		Data: dhtUser,
	}

	dhtUserBytes, err := json.Marshal(sigMsg)
	if err != nil {
		log.Printf("Failed to marshall dhtUser: %v", err)
		return
	}
	send(senderAddr, dhtUserBytes)
}

func onAnnounceMsg(msg models.SigMsg, senderAddr string) {
	dataBytes, err := json.Marshal(msg.Data)
	if err != nil {
		log.Printf("Failed to marshal Data: %v", err)
		return
	}

	var annreq models.AnnounceRequest
	if err := json.Unmarshal(dataBytes, &annreq); err != nil {
		log.Printf("Failed to unmarshal AnnounceRequest: %v", err)
		return
	}

	var username string
	token, _, err := new(jwt.Parser).ParseUnverified(annreq.Cert, jwt.MapClaims{})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username = fmt.Sprint(claims["name"])
	}
	if username == "" {
		return
	}

	userDhtLock.Lock()
	userDht[username] = models.DhtUser{
		LoginCert: annreq.Cert,
		IP:        senderAddr,
	}
	userDhtLock.Unlock()

	fmt.Println(userDht)
}

func onConnectionMsg(msg models.SigMsg, senderAddr string) {
	dataBytes, err := json.Marshal(msg.Data)
	if err != nil {
		log.Printf("Failed to marshal Data: %v", err)
		return
	}

	var conreq models.ConnectionRequest
	if err := json.Unmarshal(dataBytes, &conreq); err != nil {
		log.Printf("Failed to unmarshal ConnectionRequest: %v", err)
		return
	}

	// Set the sender's address
	conreq.SrcIP = &senderAddr

	// Marshal the updated connection request
	message, err := json.Marshal(conreq)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		return
	}

	if err := send(conreq.DestIP, message); err != nil {
		log.Printf("Error sending message: %v", err)
		return
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
			return fmt.Errorf("client not found for address %s", remoteAddr)
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
