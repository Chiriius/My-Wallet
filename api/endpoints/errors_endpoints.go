package endpoints

import "errors"

var ErrInvalidCredentials = errors.New("Invalid email or password")
var ErrInterfaceWrong = errors.New("Request interface type wrong")
