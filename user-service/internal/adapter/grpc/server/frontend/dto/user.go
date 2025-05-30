package dto

import (
	"github.com/aftosmiros/moodu/user-service/internal/domain"
	"github.com/aftosmiros/moodu/user-service/proto/userpb"
)

func ToProfileResponse(u *domain.User) *userpb.GetProfileResponse {
	res := &userpb.GetProfileResponse{
		FullName: u.FullName,
		Email:    u.Email,
		Role:     u.Role,
		Barcode:  u.Barcode,
	}

	if u.StudentProfile != nil {
		res.Gpa = u.StudentProfile.GPA
		res.Certificates = u.StudentProfile.Certificates
		res.BankDetails = u.StudentProfile.BankDetails
		res.Major = u.StudentProfile.Major
	}

	if u.InstructorProfile != nil {
		res.Department = u.InstructorProfile.Department
	}

	return res
}

func FromRegisterRequest(req *userpb.RegisterRequest) *domain.User {
	user := &domain.User{
		FullName:          req.FullName,
		Email:             req.Email,
		Role:              req.Role,
		Barcode:           req.Barcode,
		StudentProfile:    nil,
		InstructorProfile: nil,
	}

	if req.Role == "student" {
		user.StudentProfile = &domain.StudentProfile{
			GPA:          req.Gpa,
			Certificates: req.Certificates,
			BankDetails:  req.BankDetails,
			Major:        req.Major,
		}
	}

	if req.Role == "instructor" {
		user.InstructorProfile = &domain.InstructorProfile{
			Department: req.Department,
		}
	}

	return user
}

func FromUpdateProfileRequest(req *userpb.UpdateProfileRequest) *domain.UpdateProfileInput {
	var student *domain.StudentProfile
	if req.Gpa != 0 || len(req.Certificates) > 0 || req.BankDetails != "" || req.Major != "" {
		student = &domain.StudentProfile{
			GPA:          req.Gpa,
			Certificates: req.Certificates,
			BankDetails:  req.BankDetails,
			Major:        req.Major,
		}
	}

	var instructor *domain.InstructorProfile
	if req.Department != "" {
		instructor = &domain.InstructorProfile{
			Department: req.Department,
		}
	}

	return &domain.UpdateProfileInput{
		UserID:            req.UserId,
		FullName:          req.FullName,
		StudentProfile:    student,
		InstructorProfile: instructor,
	}
}
