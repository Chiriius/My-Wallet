package entities

import "time"

type User struct {
	ID           string    `json:"id,omitempty" bson:"id,omitempty"`
	DNI          int       `validate:"required"`
	Name         string    `validate:"required"`
	Email        string    `validate:"required,email"`
	Password     string    `validate:"required"`
	Address      string    `validate:"required"`
	Phone        int       `validate:"required"`
	StateActive  bool      `validate:"required"`
	Token        string    `json:"token"`
	Created_at   time.Time `json:"created_at"`
	RefreshToken string    `json:"refresh_token"`
	Update_at    time.Time `json:"updated_at"`
}
