package models

import "github.com/pion/webrtc/v4"

type SigMessageType int

const (
	Announce   SigMessageType = 0
	UserSearch SigMessageType = 1
	Connection SigMessageType = 2
	UserAdd    SigMessageType = 3
)

type SigMsg struct {
	Type SigMessageType `json:"type"`
	Data interface{}    `json:"data"`
}

type AnnounceRequest struct {
	Cert string `json:"cert"`
}

type ConnectionRequest struct {
	Type      string               `json:"type"`
	DestIP    string               `json:"destip"`
	SrcIP     string               `json:"srcip"`
	Data      string               `json:"data"`
	Candidate *webrtc.ICECandidate `json:"candidate,omitempty"`
}

type UserSearchRequest struct {
	Username string `json:"username"`
}

type DhtUser struct {
	LoginCert string `json:"cert"`
	IP        string `json:"ip"`
}

type UserAddRequest struct {
	User      DhtUser `json:"user"`
	IP        string  `json:"ip"`
	Signature string  `json:"signature"`
}

type UserAddResponse struct {
	User      DhtUser `json:"user"`
	IP        string  `json:"ip"`
	Signature string  `json:"signature"`
}
