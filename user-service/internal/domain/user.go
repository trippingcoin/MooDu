package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	FullName          string
	Email             string
	PasswordHash      string
	Role              string // student, instructor, admin
	Barcode           string
	IsDeleted         bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	StudentProfile    *StudentProfile
	InstructorProfile *InstructorProfile
}

func (u User) Validate() error {
	if u.FullName == "" {
		return ErrEmptyFullName
	}
	if u.Email == "" {
		return ErrEmptyEmail
	}
	if u.PasswordHash == "" {
		return ErrEmptyPassword
	}
	if u.Role == "" {
		return ErrEmptyRole
	}
	if u.Barcode == "" {
		return ErrEmptyBarcode
	}
	return nil
}

type StudentProfile struct {
	GPA          float32
	Certificates []string
	BankDetails  string
	Major        string
}

type InstructorProfile struct {
	Department string
}

type UpdateProfileInput struct {
	UserID            string
	FullName          string
	StudentProfile    *StudentProfile
	InstructorProfile *InstructorProfile
}
