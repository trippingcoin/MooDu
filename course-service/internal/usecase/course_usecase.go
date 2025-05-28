package usecase

import "cs/course-service/internal/model"

type CourseRepository interface {
	Create(course *model.Course) error
	GetByID(id string) (*model.Course, error)
	List() ([]*model.Course, error)
}

type CourseUsecase struct {
	repo CourseRepository
}

func New(repo CourseRepository) *CourseUsecase {
	return &CourseUsecase{repo: repo}
}

func (uc *CourseUsecase) Create(course *model.Course) error {
	return uc.repo.Create(course)
}

func (uc *CourseUsecase) GetByID(id string) (*model.Course, error) {
	return uc.repo.GetByID(id)
}

func (uc *CourseUsecase) List() ([]*model.Course, error) {
	return uc.repo.List()
}
