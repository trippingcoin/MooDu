package dao

import (
	"time"

	"github.com/aftosmiros/moodu/user-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	FullName          string             `bson:"full_name"`
	Email             string             `bson:"email"`
	PasswordHash      string             `bson:"password_hash"`
	Role              string             `bson:"role"`
	Barcode           string             `bson:"barcode"`
	IsDeleted         bool               `bson:"is_deleted"`
	CreatedAt         time.Time          `bson:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at"`
	StudentProfile    *StudentProfile    `bson:"student_profile,omitempty"`
	InstructorProfile *InstructorProfile `bson:"instructor_profile,omitempty"`
}

type StudentProfile struct {
	GPA          float32  `bson:"gpa"`
	Certificates []string `bson:"certificates"`
	BankDetails  string   `bson:"bank_details"`
	Major        string   `bson:"major"`
}

type InstructorProfile struct {
	Department string `bson:"department"`
}

func FromUser(u *domain.User) *User {
	return &User{
		ID:           u.ID,
		FullName:     u.FullName,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Role:         u.Role,
		Barcode:      u.Barcode,
		IsDeleted:    u.IsDeleted,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		StudentProfile: func() *StudentProfile {
			if u.StudentProfile == nil {
				return nil
			}
			return &StudentProfile{
				GPA:          u.StudentProfile.GPA,
				Certificates: u.StudentProfile.Certificates,
				BankDetails:  u.StudentProfile.BankDetails,
				Major:        u.StudentProfile.Major,
			}
		}(),
		InstructorProfile: func() *InstructorProfile {
			if u.InstructorProfile == nil {
				return nil
			}
			return &InstructorProfile{
				Department: u.InstructorProfile.Department,
			}
		}(),
	}
}

func ToUser(d *User) *domain.User {
	return &domain.User{
		ID:           d.ID,
		FullName:     d.FullName,
		Email:        d.Email,
		PasswordHash: d.PasswordHash,
		Role:         d.Role,
		Barcode:      d.Barcode,
		IsDeleted:    d.IsDeleted,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
		StudentProfile: func() *domain.StudentProfile {
			if d.StudentProfile == nil {
				return nil
			}
			return &domain.StudentProfile{
				GPA:          d.StudentProfile.GPA,
				Certificates: d.StudentProfile.Certificates,
				BankDetails:  d.StudentProfile.BankDetails,
				Major:        d.StudentProfile.Major,
			}
		}(),
		InstructorProfile: func() *domain.InstructorProfile {
			if d.InstructorProfile == nil {
				return nil
			}
			return &domain.InstructorProfile{
				Department: d.InstructorProfile.Department,
			}
		}(),
	}
}
