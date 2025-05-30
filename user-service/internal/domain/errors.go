package domain

import "errors"

var (
	// Common
	ErrInvalidID = errors.New("invalid object ID")
	ErrNotFound  = errors.New("resource not found")

	// Auth related
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized access")

	// Validation
	ErrEmailAlreadyExists = errors.New("email already registered")
	ErrInvalidRole        = errors.New("invalid role provided")

	// Session/Token
	ErrSessionExpired = errors.New("session has expired")
	ErrTokenInvalid   = errors.New("invalid token")
)
