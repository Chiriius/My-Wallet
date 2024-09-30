package models

type User struct {
	DNI      int    `json:"dni" `
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    int    `json:"phone"`
}
