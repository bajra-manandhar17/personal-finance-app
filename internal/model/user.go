package model

import "time"

type User struct {
	UserId         string    `json:"user_id"`
	FirstName      string    `json:"first_name"`
	MiddleName     string    `json:"middle_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	PhoneNumber    string    `json:"phone_number"`
	Currency       string    `json:"currency"`
	ProfilePicture string    `json:"profile_picture"`
	IsVerified     bool      `json:"is_verified"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
