package models

type User struct {
	UID      string `json:"id"`
	DNI      int    `json:"dni" `
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    int    `json:"phone"`
}
