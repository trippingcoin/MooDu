package dao

import (
	"time"

	"github.com/aftosmiros/moodu/user-service/internal/domain"
)

type Session struct {
	UserID       string    `bson:"user_id"`
	RefreshToken string    `bson:"refresh_token"`
	ExpiresAt    time.Time `bson:"expires_at"`
	CreatedAt    time.Time `bson:"created_at"`
}

func FromSession(s domain.Session) Session {
	return Session{
		UserID:       s.UserID,
		RefreshToken: s.RefreshToken,
		ExpiresAt:    s.ExpiresAt,
		CreatedAt:    s.CreatedAt,
	}
}

func ToSession(s Session) domain.Session {
	return domain.Session{
		UserID:       s.UserID,
		RefreshToken: s.RefreshToken,
		ExpiresAt:    s.ExpiresAt,
		CreatedAt:    s.CreatedAt,
	}
}
