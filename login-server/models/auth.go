package models

import "github.com/google/uuid"

// LoginRequest model info
// @Description Login request with email and password
type LoginRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
} //@name LoginRequest

// JWT claims
// @Description Created when client calls a endpoint, used for auth and rbac
type JWT struct {
	ID   uuid.UUID `validate:"required" json:"id"`
	Exp  int       `validate:"required" json:"exp"`
	Name string    `validate:"required" json:"name"`
	Sub  string    `validate:"required" json:"sub"`
} //@name JWT
