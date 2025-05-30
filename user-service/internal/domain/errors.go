package domain

import "errors"

var (
	// Common
	ErrInvalidID       = errors.New("invalid object ID")
	ErrNotFound        = errors.New("resource not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrEmptyPassword   = errors.New("empty password")
	ErrEmptyEmail      = errors.New("empty email")
	ErrEmptyRole       = errors.New("empty role")
	ErrEmptyFullName   = errors.New("empty full name")
	ErrEmptyBarcode    = errors.New("empty barcode")

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
