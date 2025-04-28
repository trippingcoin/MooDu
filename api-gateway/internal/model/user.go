package model

import "time"

type User struct {
	ID              uint64
	Name            string
	Phone           string
	Email           string
	CurrentPassword string
	NewPassword     string
	CreatedAt       time.Time
	UpdatedAt       time.Time

	IsDeleted bool
}
