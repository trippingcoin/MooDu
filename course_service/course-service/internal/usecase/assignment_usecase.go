package usecase

import (
	"context"
	cache "cs/course-service/internal/cache/redis"
	"cs/course-service/internal/model"
	"cs/course-service/pkg/broker"
	"cs/course-service/pkg/transactor"
	"errors"
	"time"
)

type AssignmentRepository interface {
	Create(*model.Assignment) error
	Update(*model.Assignment) error
	Delete(id string) error
	GetByID(id string) (*model.Assignment, error)
	List() ([]*model.Assignment, error)

	AddSubmission(assignmentID string, submissionID string) error
	AddSubmissions(assignmentID string, submissionIDs []string) error
}

type AssignmentUsecase struct {
	repo      AssignmentRepository
	publisher broker.Publisher
	callTx    transactor.WithinTransactionFunc
	cache     cache.AssignmentCache
}

func NewAssignemntUc(repo AssignmentRepository, publisher broker.Publisher, callTx transactor.WithinTransactionFunc, c cache.AssignmentCache) *AssignmentUsecase {
	return &AssignmentUsecase{
		repo:      repo,
		publisher: publisher,
		callTx:    callTx,
		cache:     c,
	}
}

func (u *AssignmentUsecase) Create(ctx context.Context, assignment *model.Assignment) error {
	txFn := func(ctx context.Context) error {
		if assignment == nil {
			return errors.New("assignment cannot be nil")
		}
		if assignment.CourseID == "" {
			return errors.New("course ID is required")
		}
		if assignment.Title == "" {
			return errors.New("assignment title is required")
		}

		if err := u.repo.Create(assignment); err != nil {
			return err
		}

		err := u.publisher.PublishAssignmentCreated(assignment)
		if err != nil {
			return err
		}
		return nil
	}
	err := u.callTx(ctx, txFn)
	if err != nil {
		return err
	}

	return nil
}

func (u *AssignmentUsecase) Update(assignment *model.Assignment) error {
	if assignment.ID == "" {
		return errors.New("assignment ID is required")
	}

	if err := u.repo.Update(assignment); err != nil {
		return err
	}

	u.publisher.PublishAssignmentUpdated(assignment)
	_ = u.cache.Set(assignment)

	return nil
}

func (u *AssignmentUsecase) Delete(id string) error {
	if id == "" {
		return errors.New("assignment ID is required")
	}

	if err := u.repo.Delete(id); err != nil {
		return err
	}

	u.publisher.PublishAssignmentDeleted(id)
	_ = u.cache.Delete(id)

	return nil
}

func (u *AssignmentUsecase) GetByID(id string) (*model.Assignment, error) {
	if id == "" {
		return nil, errors.New("assignment ID is required")
	}

	cached, err := u.cache.Get(id)
	if err == nil && cached != nil {
		return cached, nil
	}

	assignment, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	_ = u.cache.Set(assignment)

	return assignment, nil
}

func (u *AssignmentUsecase) List() ([]*model.Assignment, error) {
	cachedList, err := u.cache.List()
	if err == nil && cachedList != nil {
		return cachedList, nil
	}

	assignments, err := u.repo.List()
	if err != nil {
		return nil, err
	}

	_ = u.cache.SetList(assignments)

	return assignments, nil
}

func (u *AssignmentUsecase) AddSubmission(assignmentID string, submissionID string) error {
	if assignmentID == "" || submissionID == "" {
		return errors.New("assignment ID and submission ID are required")
	}

	err := u.repo.AddSubmission(assignmentID, submissionID)
	if err != nil {
		return err
	}

	assignment, err := u.repo.GetByID(assignmentID)
	if err == nil {
		_ = u.cache.Set(assignment)
	}

	return nil
}

func (u *AssignmentUsecase) AddSubmissions(assignmentID string, submissionIDs []string) error {
	if assignmentID == "" || len(submissionIDs) == 0 {
		return errors.New("assignment ID and submission IDs are required")
	}

	assignment, err := u.repo.GetByID(assignmentID)
	if err != nil {
		return err
	}

	assignment.Submissions = append(assignment.Submissions, submissionIDs...)
	assignment.UpdatedAt = time.Now().Format(time.RFC3339)

	if err := u.repo.Update(assignment); err != nil {
		return err
	}

	_ = u.cache.Set(assignment)
	return nil
}
