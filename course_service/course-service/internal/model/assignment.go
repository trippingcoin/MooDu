package model

type Assignment struct {
	ID          string   `bson:"_id,omitempty"`
	Title       string   `bson:"title"`
	Grade       int      `bson:"grade"`
	Description string   `bson:"description"`
	CourseID    string   `bson:"course_id"`
	DueDate     string   `bson:"due_date"`
	Submissions []string `bson:"submissions,omitempty"`
	CreatedAt   string   `bson:"created_at,omitempty"`
	UpdatedAt   string   `bson:"updated_at,omitempty"`
}
