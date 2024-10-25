package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/piheta/sept/backend/db"
	"github.com/piheta/sept/backend/models"
	"github.com/piheta/sept/backend/repos"
	"github.com/pion/webrtc/v4"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type SnConnection struct {
	ctx      context.Context
	peerConn *webrtc.PeerConnection
	ws       *websocket.Conn
	userChan chan models.User
	chosenIP string
}

func NewSnConnection() *SnConnection {
	return &SnConnection{
		userChan: make(chan models.User),
	}
}

func (s *SnConnection) SetContext(ctx context.Context) {
	s.ctx = ctx
}

// p1p2, connects to the signaling server
// p1 creates and sends offer to the chosen peer
// p2 creates and sends answer to p1
// p1 sends ICE candidates to p2
// p2 replies with his ICE candidates
// datachannel is made
func (s *SnConnection) SnConnectionHandler() {
	s.initializePeerConnection()
	s.createDataChannel()

	var wg sync.WaitGroup
	wg.Add(1)
	go s.connectToSignalingServer(&wg)
	wg.Wait()
}

func (s *SnConnection) initializePeerConnection() {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	var err error
	s.peerConn, err = webrtc.NewPeerConnection(config)
	if err != nil {
		fmt.Println("sn error, ", err)
	}

	s.peerConn.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate == nil {
			return
		}
		s.sendICECandidate(candidate)
	})

	s.peerConn.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", state.String())
	})
}

func (s *SnConnection) createDataChannel() {
	sendChannel, err := s.peerConn.CreateDataChannel("foo", nil)
	if err != nil {
		fmt.Println("sn error, ", err)
		return
	}
	fmt.Println("Data channel created:", sendChannel.Label())

	sendChannel.OnClose(func() {
		fmt.Println("sendChannel has closed")
	})

	sendChannel.OnOpen(func() {
		fmt.Println("sendChannel has opened")
		candidatePair, err := s.peerConn.SCTP().Transport().ICETransport().GetSelectedCandidatePair()
		if err != nil {
			fmt.Println("Error getting candidate pair:", err)
		} else {
			fmt.Println("Selected candidate pair:", candidatePair)
		}
	})

	sendChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		var message models.Message
		if err := json.Unmarshal(msg.Data, &message); err != nil {
			log.Printf("Failed to unmarshal p2p Message: %v", err)
			return
		}

		message_repo := repos.NewMessageRepo(db.DB)
		message_repo.AddMessage(message)
		runtime.EventsEmit(s.ctx, "message:new")

		fmt.Printf("%s: %s\n", sendChannel.Label(), string(msg.Data)) //* HANDLES RECIEVED P2P MESSAGE
	})

	s.peerConn.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

		// Register channel opening handling
		d.OnOpen(func() {
			NewMessagingHandler(d)
		})
	})

	// Add state change handler for the peer connection
	s.peerConn.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s\n", s.String())
	})

	// Add state change handler for the ICE connection
	s.peerConn.OnICEConnectionStateChange(func(s webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", s.String())
	})
}

func (s *SnConnection) connectToSignalingServer(wg *sync.WaitGroup) {
	defer wg.Done()

	var err error
	s.ws, _, err = websocket.DefaultDialer.Dial("ws://127.0.0.1:8081/ws", nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)

	}
	defer close(s.userChan)
	defer s.ws.Close()

	//! Get and send cert to sig
	announceMessage, err := s.createAnnounceRequest()
	if err != nil {
		log.Fatalf("Failed to create announce request: %v", err)
	}

	err = s.ws.WriteJSON(announceMessage)
	if err != nil {
		log.Fatalf("Failed to send user data: %v", err)
	}

	//* Listen for messages from sig
	for {
		_, message, err := s.ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		var sigMessage models.SigMsg
		if err := json.Unmarshal(message, &sigMessage); err != nil {
			log.Printf("Failed to unmarshal sigMessage: %v", err)
			continue
		}

		switch sigMessage.Type {
		case models.UserSearch:
			// Catch server response
			fmt.Println(sigMessage.Data)
			s.userSearchResponse(sigMessage)
		case models.Connection:

			dataBytes, err := json.Marshal(sigMessage.Data)
			if err != nil {
				log.Printf("Failed to marshal Data: %v", err)
				return
			}

			var connectionRequest models.ConnectionRequest
			if err := json.Unmarshal(dataBytes, &connectionRequest); err != nil {
				log.Printf("Failed to unmarshal ConnectionRequest: %v", err)
				return
			}

			// todo, send these as sigmsg
			switch connectionRequest.Type {
			case "offer":
				s.onSDPOffer(connectionRequest)
			case "answer":
				s.onSDPAnswer(connectionRequest)
			case "candidate":
				s.onICECandidate(connectionRequest)
			}

		default:
			fmt.Println("Unknown message type:", sigMessage.Type)
		}
	}
}

// ! User search request sent to the sig server. The response is captured in the switch above.
func (s *SnConnection) UserSearchRequest(username string) (<-chan models.User, error) {
	req := models.SigMsg{
		Type: models.UserSearch,
		Data: models.UserSearchRequest{
			Username: username,
		},
	}

	err := s.ws.WriteJSON(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send user data: %v", err)
	}

	return s.userChan, nil
}

// ! Ran when data from sig server is recieved and marked as UserSearch
func (s *SnConnection) userSearchResponse(msg models.SigMsg) {
	dataBytes, err := json.Marshal(msg.Data)
	if err != nil {
		log.Printf("Failed to marshal Data: %v", err)
		return
	}

	var dhtuser models.DhtUser
	if err := json.Unmarshal(dataBytes, &dhtuser); err != nil {
		log.Printf("Failed to unmarshal AnnounceRequest: %v", err)
		return
	}

	cert := dhtuser.LoginCert
	if err = VerifyToken(cert); err != nil {
		log.Printf("Token of found user is not valid: %v", err)
		return
	}

	user, err := ExtractUserFromJwt(cert)
	if err != nil {
		log.Printf("Failed to extract found user from jwt, ", err)
		return
	}

	user.Ip = dhtuser.IP

	s.userChan <- user
}

func (s *SnConnection) createAnnounceRequest() (models.SigMsg, error) {
	cert, err := os.ReadFile(db.SEPT_DATA + "/user.jwt")
	if err != nil {
		return models.SigMsg{}, fmt.Errorf("failed to get user cert: %v", err)
	}

	annreq := models.SigMsg{
		Type: models.Announce,
		Data: models.AnnounceRequest{
			Cert: string(cert),
		},
	}

	return annreq, nil
}

// ICE
// Senders
func (s *SnConnection) SendSDPOffer(destIp string) {
	sigMsg := models.SigMsg{
		Type: models.Connection,
		Data: models.ConnectionRequest{
			Type:   "offer",
			DestIP: destIp,
			Data:   s.createSDPOffer(),
		},
	}

	s.chosenIP = destIp
	if err := s.ws.WriteJSON(sigMsg); err != nil {
		fmt.Println("sn error, ", err)
	}
}

func (s *SnConnection) sendSDPAnswer(destIP, answer string) {
	sigMsg := models.SigMsg{
		Type: models.Connection,
		Data: models.ConnectionRequest{
			Type:   "answer",
			DestIP: destIP,
			Data:   answer,
		},
	}

	if err := s.ws.WriteJSON(sigMsg); err != nil {
		fmt.Println("sn error, ", err)
	}
}

func (s *SnConnection) sendICECandidate(candidate *webrtc.ICECandidate) {
	candidateJSON, err := json.Marshal(candidate.ToJSON())
	if err != nil {
		fmt.Println("sn error, ", err)
		return
	}

	sigMsg := models.SigMsg{
		Type: models.Connection,
		Data: models.ConnectionRequest{
			Type:      "candidate",
			Candidate: candidate,
			DestIP:    s.chosenIP,
			Data:      string(candidateJSON),
		},
	}

	if err := s.ws.WriteJSON(sigMsg); err != nil {
		fmt.Println("sn error, ", err)
	}
}

func (s *SnConnection) onSDPOffer(connectionRequest models.ConnectionRequest) {
	runtime.EventsEmit(s.ctx, "offer:new", connectionRequest.SrcIP)
	fmt.Println("Received offer:", connectionRequest)
	answer := s.createSDPAnswer(connectionRequest.Data)
	s.sendSDPAnswer(connectionRequest.SrcIP, answer)
}

func (s *SnConnection) onSDPAnswer(connectionRequest models.ConnectionRequest) {
	fmt.Println("Received answer:", connectionRequest.Data)
	answerBytes, err := base64.StdEncoding.DecodeString(connectionRequest.Data)
	if err != nil {
		fmt.Println("sn error, ", err)
	}
	answerSDP := string(answerBytes)
	answerDesc := webrtc.SessionDescription{
		Type: webrtc.SDPTypeAnswer,
		SDP:  answerSDP,
	}
	if err := s.peerConn.SetRemoteDescription(answerDesc); err != nil {
		fmt.Println("sn error, ", err)
	}
}

func (s *SnConnection) onICECandidate(connectionRequest models.ConnectionRequest) {
	fmt.Println("Received ICE candidate")
	candidate := webrtc.ICECandidateInit{}
	if err := json.Unmarshal([]byte(connectionRequest.Data), &candidate); err != nil {
		fmt.Println("sn error, ", err)
		return
	}
	s.chosenIP = connectionRequest.SrcIP // Replace "none" with the sender of the offer
	if err := s.peerConn.AddICECandidate(candidate); err != nil {
		fmt.Println("sn error, ", err)
	}
}

// Helpers

func (s *SnConnection) createSDPOffer() string {
	offer, err := s.peerConn.CreateOffer(nil)
	if err != nil {
		fmt.Println("sn error, ", err)
	}
	if err := s.peerConn.SetLocalDescription(offer); err != nil {
		fmt.Println("sn error, ", err)
	}

	return base64.StdEncoding.EncodeToString([]byte(offer.SDP))
}

func (s *SnConnection) createSDPAnswer(offerBase64 string) string {
	offerBytes, err := base64.StdEncoding.DecodeString(offerBase64)
	if err != nil {
		fmt.Println("sn error, ", err)
	}
	offerSDP := string(offerBytes)

	offer := webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer,
		SDP:  offerSDP,
	}
	if err := s.peerConn.SetRemoteDescription(offer); err != nil {
		fmt.Println("sn error, ", err)
	}

	answer, err := s.peerConn.CreateAnswer(nil)
	if err != nil {
		fmt.Println("sn error, ", err)
	}

	if err := s.peerConn.SetLocalDescription(answer); err != nil {
		fmt.Println("sn error, ", err)
	}

	return base64.StdEncoding.EncodeToString([]byte(answer.SDP))
}
