package broker

import "cs/course-service/internal/model"

type Publisher interface {
	PublishCourseCreated(course *model.Course) error
	PublishCourseUpdated(course *model.Course) error
	PublishCourseDeleted(courseID string) error
	PublishAssignmentCreated(assignment *model.Assignment) error
	PublishAssignmentUpdated(assignment *model.Assignment) error
	PublishAssignmentDeleted(assignmentID string) error
}
