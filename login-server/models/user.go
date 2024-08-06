package models

import "github.com/google/uuid"

// User model info
// @Description User account information
type User struct {
	ID       uuid.UUID `gorm:"primary_key; unique; type:uuid;" swaggerignore:"true" json:"id"`
	Name     string    `validate:"required" json:"name"`
	Email    string    `validate:"required" json:"email"`
	Password string    `validate:"required" json:"password"`
} //@name User

// UpdateUser model info
// @Description UpdateUser account information
type UpdateUser struct {
	Name     *string
	Email    *string
	Password *string
} //@name UpdateUser
