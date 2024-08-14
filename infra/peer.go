package peer

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/piheta/sept/models"
	"github.com/pion/webrtc/v4"
)

var chosenIP string
var peerConnection *webrtc.PeerConnection
var ws *websocket.Conn

func handleError(err error) {
	fmt.Println("Unexpected error. Check console.")
	log.Println(err)
}

func Peer() {
	initializePeerConnection()
	createDataChannel()

	var wg sync.WaitGroup
	wg.Add(1)
	go connectToSignalingServer(&wg)
	wg.Wait()
}

func initializePeerConnection() {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	var err error
	peerConnection, err = webrtc.NewPeerConnection(config)
	if err != nil {
		handleError(err)
	}

	peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate == nil {
			return
		}
		sendICECandidate(candidate)
	})

	peerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", state.String())
	})
}

func createDataChannel() {
	sendChannel, err := peerConnection.CreateDataChannel("foo", nil)
	if err != nil {
		handleError(err)
		return
	}
	fmt.Println("Data channel created:", sendChannel.Label())

	sendChannel.OnClose(func() {
		fmt.Println("sendChannel has closed")
	})

	sendChannel.OnOpen(func() {
		fmt.Println("sendChannel has opened")
		candidatePair, err := peerConnection.SCTP().Transport().ICETransport().GetSelectedCandidatePair()
		if err != nil {
			fmt.Println("Error getting candidate pair:", err)
		} else {
			fmt.Println("Selected candidate pair:", candidatePair)
		}
	})

	sendChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("%s: %s\n", sendChannel.Label(), string(msg.Data))
	})

	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

		// Register channel opening handling
		d.OnOpen(func() {
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
		})
	})

	// Add state change handler for the peer connection
	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s\n", s.String())
	})

	// Add state change handler for the ICE connection
	peerConnection.OnICEConnectionStateChange(func(s webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", s.String())
	})
}

func createOffer() string {
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		handleError(err)
	}
	if err := peerConnection.SetLocalDescription(offer); err != nil {
		handleError(err)
	}

	return base64.StdEncoding.EncodeToString([]byte(offer.SDP))
}

func createAnswer(offerBase64 string) string {
	offerBytes, err := base64.StdEncoding.DecodeString(offerBase64)
	if err != nil {
		handleError(err)
	}
	offerSDP := string(offerBytes)

	offer := webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer,
		SDP:  offerSDP,
	}
	if err := peerConnection.SetRemoteDescription(offer); err != nil {
		handleError(err)
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		handleError(err)
	}

	if err := peerConnection.SetLocalDescription(answer); err != nil {
		handleError(err)
	}

	return base64.StdEncoding.EncodeToString([]byte(answer.SDP))
}

func sendICECandidate(candidate *webrtc.ICECandidate) {
	candidateJSON, err := json.Marshal(candidate.ToJSON())
	if err != nil {
		handleError(err)
		return
	}

	cr := models.ConnectionRequest{
		Type:      "candidate",
		Candidate: candidate,
		DestIP:    chosenIP,
		Data:      string(candidateJSON),
	}
	crBytes, err := json.Marshal(cr)
	if err != nil {
		handleError(err)
		return
	}

	if err := ws.WriteMessage(websocket.TextMessage, crBytes); err != nil {
		handleError(err)
	}
}

func connectToSignalingServer(wg *sync.WaitGroup) {
	defer wg.Done()

	var err error
	ws, _, err = websocket.DefaultDialer.Dial("ws://20.100.14.52:8080/ws", nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer ws.Close()

	_, message, err := ws.ReadMessage()
	if err != nil {
		log.Fatalf("Failed to read message: %v", err)
	}
	var ips []string
	if err := json.Unmarshal(message, &ips); err != nil {
		log.Fatalf("Failed to unmarshal IPs: %v", err)
	}
	log.Printf("Nodes connected to signaling server: %v", ips)

	if len(ips) > 0 {
		fmt.Println("Choose IP:")
		fmt.Scanln(&chosenIP)

		if chosenIP != "none" {
			cr := models.ConnectionRequest{Type: "offer", DestIP: chosenIP, Data: createOffer()}
			crBytes, err := json.Marshal(cr)
			if err != nil {
				log.Fatalf("Failed to marshal ConnectionRequest: %v", err)
			}

			err = ws.WriteMessage(websocket.TextMessage, crBytes)
			if err != nil {
				log.Fatalf("Failed to send chosen IP: %v", err)
			}
		}
	}

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		var connectionRequest models.ConnectionRequest
		if err := json.Unmarshal(message, &connectionRequest); err != nil {
			log.Printf("Failed to unmarshal ConnectionRequest: %v", err)
			continue
		}

		switch connectionRequest.Type {
		case "offer":
			fmt.Println("Received offer:", connectionRequest)
			answer := createAnswer(connectionRequest.Data)
			cr := models.ConnectionRequest{Type: "answer", DestIP: *connectionRequest.SrcIP, Data: answer}
			crBytes, err := json.Marshal(cr)
			if err != nil {
				log.Fatalf("Failed to marshal ConnectionRequest: %v", err)
			}

			err = ws.WriteMessage(websocket.TextMessage, crBytes)
			if err != nil {
				log.Fatalf("Failed to send answer: %v", err)
			}
		case "answer":
			fmt.Println("Received answer:", connectionRequest.Data)
			answerBytes, err := base64.StdEncoding.DecodeString(connectionRequest.Data)
			if err != nil {
				handleError(err)
			}
			answerSDP := string(answerBytes)
			answerDesc := webrtc.SessionDescription{
				Type: webrtc.SDPTypeAnswer,
				SDP:  answerSDP,
			}
			if err := peerConnection.SetRemoteDescription(answerDesc); err != nil {
				handleError(err)
			}
		case "candidate":
			fmt.Println("Received ICE candidate")
			candidate := webrtc.ICECandidateInit{}
			if err := json.Unmarshal([]byte(connectionRequest.Data), &candidate); err != nil {
				handleError(err)
				continue
			}
			chosenIP = *connectionRequest.SrcIP //replace "none" with the sender of the offer
			if err := peerConnection.AddICECandidate(candidate); err != nil {
				handleError(err)
			}

		default:
			fmt.Println("Unknown message type:", connectionRequest.Type)
		}
	}
}
