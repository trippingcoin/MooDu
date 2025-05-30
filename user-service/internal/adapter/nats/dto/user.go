package dto

import (
	"time"

	"github.com/aftosmiros/moodu/user-service/internal/domain"
)

type UserCreatedEvent struct {
	ID         string    `json:"id"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Barcode    string    `json:"barcode"`
	Major      string    `json:"major,omitempty"`
	Department string    `json:"department,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserUpdatedEvent struct {
	ID         string    `json:"id"`
	FullName   string    `json:"full_name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Role       string    `json:"role,omitempty"`
	Major      string    `json:"major,omitempty"`
	Department string    `json:"department,omitempty"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserDeletedEvent struct {
	ID        string    `json:"id"`
	DeletedAt time.Time `json:"deleted_at"`
}

type LoginAttemptEvent struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip,omitempty"`
}

func FromDomainUser(user *domain.User) *UserCreatedEvent {
	event := &UserCreatedEvent{
		ID:       user.ID.Hex(),
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		Barcode:  user.Barcode,
	}

	if user.StudentProfile != nil {
		event.Major = user.StudentProfile.Major
	}

	if user.InstructorProfile != nil {
		event.Department = user.InstructorProfile.Department
	}

	return event
}
