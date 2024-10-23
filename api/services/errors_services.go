package services

import "errors"

var ErrLenghtPassword = errors.New("minimum password length 8")
var ErrLenghPhone = errors.New("Length of phone number 10")
var ErrNameSpecialCharacters = errors.New("the name field must not contain special characters")
var ErrTypeDNI = errors.New("ID type must be CC or NIT")
var ErrHashingPassword = errors.New("Error hashing the password")
var ErrUserNotfound = errors.New("Error not found user")
var ErrInvalidCredentials = errors.New("Invalid email or password")
