package models

type user struct {
	UID      string `json:"id" omitempty`
	DNI      int    `json:"dni" `
	Email    string `json:"email"`
	Password string `json:password`
	Address  string `json:"address"`
	Phone    int    `json:"phone"`
}
