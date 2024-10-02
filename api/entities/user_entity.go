package entities

type User struct {
	ID       string `json:"id,omitempty" bson:"id,omitempty"`
	DNI      int    `validate:"required"`
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Address  string `validate:"required"`
	Phone    int    `validate:"required"`
	State    bool   `validate:"required"`
}
