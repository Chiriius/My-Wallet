package repository_user

import "errors"

var ErrDisbledUser = errors.New("Disabled user")
var ErrUserNotfound = errors.New("Error not found user")
var ErrNotasks = errors.New("No tasks were deleted")
