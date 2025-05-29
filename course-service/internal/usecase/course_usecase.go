package usecase

import (
	"context"
	"cs/course-service/internal/cache"
	"cs/course-service/internal/model"
	"cs/course-service/pkg/broker"
	"cs/course-service/pkg/transactor"

	"log"
	"time"
)

type CourseRepository interface {
	Create(course *model.Course) error
	Update(course *model.Course) error
	GetByID(id string) (*model.Course, error)
	List() ([]*model.Course, error)
	Delete(id string) error
}

type CourseUsecase struct {
	repo      CourseRepository
	publisher broker.Publisher
	callTx    transactor.WithinTransactionFunc
	cache     cache.CourseCache
}

func New(repo CourseRepository, publisher broker.Publisher, callTx transactor.WithinTransactionFunc, c cache.CourseCache) *CourseUsecase {
	return &CourseUsecase{
		repo:      repo,
		publisher: publisher,
		callTx:    callTx,
		cache:     c,
	}
}

func (uc *CourseUsecase) Create(ctx context.Context, course *model.Course) error {
	txFn := func(ctx context.Context) error {
		course.CreatedAt = time.Now().UTC()
		course.UpdatedAt = time.Now().UTC()
		if err := uc.repo.Create(course); err != nil {
			return err
		}
		uc.cache.Set(course.ID, course)
		uc.cache.ClearList()
		err := uc.publisher.PublishCourseCreated(course)
		if err != nil {
			log.Println("uc.publisher.PublishCourseCreated: %w", err)
		}
		return nil
	}

	err := uc.callTx(ctx, txFn)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CourseUsecase) UpdateCourse(course *model.Course) error {
	err := uc.repo.Update(course)
	if err != nil {
		return err
	}
	uc.cache.Set(course.ID, course)
	uc.cache.ClearList()
	return uc.publisher.PublishCourseUpdated(course)
}

func (uc *CourseUsecase) GetByID(id string) (*model.Course, error) {
	if course, ok := uc.cache.Get(id); ok {
		return course, nil
	}

	course, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	uc.cache.Set(id, course)
	return course, nil
}

func (uc *CourseUsecase) List() ([]*model.Course, error) {
	if courses, ok := uc.cache.List(); ok {
		return courses, nil
	}

	courses, err := uc.repo.List()
	if err != nil {
		return nil, err
	}

	uc.cache.SetList(courses)
	return courses, nil
}

func (uc *CourseUsecase) Delete(id string) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}
	// uc.cache.Delete(id)
	return uc.publisher.PublishCourseDeleted(id)
}
