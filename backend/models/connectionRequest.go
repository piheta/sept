package models

import "github.com/pion/webrtc/v4"

type ConnectionRequest struct {
	Type      string               `json:"type"`
	DestIP    string               `json:"destip"`
	SrcIP     *string              `json:"srcip"`
	Data      string               `json:"data"`
	Candidate *webrtc.ICECandidate `json:"candidate,omitempty"`
}
