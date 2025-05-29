package cache

import "cs/course-service/internal/model"

type CourseCache interface {
	Get(id string) (*model.Course, bool)
	Set(id string, course *model.Course)
	Delete(id string)
	List() ([]*model.Course, bool)
	SetList(courses []*model.Course)
	ClearList()
}
