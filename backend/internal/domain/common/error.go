package common

import "errors"

var ErrNotFound = errors.New("resource not found")

var ErrAlreadyExists = errors.New("resource already exists")

var ErrPasswordMismatch = errors.New("password mismatch")
