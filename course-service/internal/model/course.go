package model

import "time"

type Course struct {
	ID          string
	Title       string
	Description string
	TeacherID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
