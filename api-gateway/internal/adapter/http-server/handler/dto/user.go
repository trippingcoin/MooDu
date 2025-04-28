package dto

type User struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`

	IsDeleted bool `json:"isDeleted"`
}

type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateResponse struct {
	ID uint64 `json:"id"`
}

type UserUpdateRequest struct {
	ID          uint64  `json:"id"`
	Email       *string `json:"email,omitempty"`
	Password    *string `json:"currentPassword,omitempty"`
	NewPassword *string `json:"newPassword,omitempty"`
	Name        *string `json:"fullName,omitempty"`
	Phone       *string `json:"phone,omitempty"`
}

type UserUpdateResponse struct {
	Email string `json:"email"`
	Name  string `json:"fullName"`
	Phone string `json:"phone"`
}
