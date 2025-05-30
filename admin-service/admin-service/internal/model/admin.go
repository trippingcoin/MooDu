package model

import "time"

type Transcript struct {
	StudentID string    `bson:"student_id"`
	Purpose   string    `bson:"purpose"`
	CreatedAt time.Time `bson:"created_at"`
}

type QueueEntry struct {
	StudentID string `bson:"student_id"`
	Reason    string `bson:"reason"`
	JoinedAt  string `bson:"joined_at"`
}

type Retake struct {
	StudentID string    `bson:"student_id"`
	CourseID  string    `bson:"course_id"`
	Reason    string    `bson:"reason"`
	CreatedAt time.Time `bson:"created_at"`
}

type ScheduleEntry struct {
	StudentID string `bson:"student_id"`
	CourseID  string `bson:"course_id"`
	Day       string `bson:"day"`
	Time      string `bson:"time"`
	Room      string `bson:"room"`
}

type CertificateRequest struct {
	StudentID       string    `bson:"student_id"`
	CertificateType string    `bson:"certificate_type"`
	Details         string    `bson:"details"`
	CreatedAt       time.Time `bson:"created_at"`
}
