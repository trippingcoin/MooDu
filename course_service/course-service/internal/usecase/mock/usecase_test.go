package usecase_test

import (
	"context"
	"errors"
	"testing"

	"cs/course-service/internal/model"
	"cs/course-service/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ---- Mocks ----

type MockRepo struct{ mock.Mock }

type MockPublisher struct{ mock.Mock }

type MockTx struct{}

type MockCache struct{ mock.Mock }

func (m *MockRepo) Create(a *model.Assignment) error { args := m.Called(a); return args.Error(0) }
func (m *MockRepo) Update(a *model.Assignment) error { args := m.Called(a); return args.Error(0) }
func (m *MockRepo) Delete(id string) error           { args := m.Called(id); return args.Error(0) }
func (m *MockRepo) GetByID(id string) (*model.Assignment, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Assignment), args.Error(1)
}
func (m *MockRepo) List() ([]*model.Assignment, error) {
	args := m.Called()
	return args.Get(0).([]*model.Assignment), args.Error(1)
}
func (m *MockRepo) AddSubmission(aid, sid string) error {
	args := m.Called(aid, sid)
	return args.Error(0)
}
func (m *MockRepo) AddSubmissions(aid string, sids []string) error {
	args := m.Called(aid, sids)
	return args.Error(0)
}

func (m *MockPublisher) PublishAssignmentCreated(a *model.Assignment) error { return nil }
func (m *MockPublisher) PublishAssignmentUpdated(a *model.Assignment) error { return nil }
func (m *MockPublisher) PublishAssignmentDeleted(id string) error           { return nil }
func (m *MockPublisher) PublishCourseCreated(a *model.Course) error         { return nil }
func (m *MockPublisher) PublishCourseUpdated(a *model.Course) error         { return nil }
func (m *MockPublisher) PublishCourseDeleted(id string) error               { return nil }

func (tx *MockTx) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}

func (c *MockCache) Get(id string) (*model.Assignment, error) { return nil, errors.New("cache miss") }
func (c *MockCache) Set(a *model.Assignment) error            { return nil }
func (c *MockCache) Delete(id string) error                   { return nil }
func (c *MockCache) List() ([]*model.Assignment, error)       { return nil, errors.New("cache miss") }
func (c *MockCache) SetList(list []*model.Assignment) error   { return nil }

// ---- Tests ----

func TestCreateAssignment_Success(t *testing.T) {
	repo := new(MockRepo)
	cache := new(MockCache)
	pub := new(MockPublisher)
	tx := &MockTx{}

	uc := usecase.NewAssignemntUc(repo, pub, tx.Run, cache)

	assignment := &model.Assignment{Title: "Test", CourseID: "C1"}
	repo.On("Create", assignment).Return(nil)

	err := uc.Create(context.Background(), assignment)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestAddSubmissions_Success(t *testing.T) {
	repo := new(MockRepo)
	cache := new(MockCache)
	pub := new(MockPublisher)
	tx := &MockTx{}

	uc := usecase.NewAssignemntUc(repo, pub, tx.Run, cache)

	assignment := &model.Assignment{ID: "A1", Submissions: []string{"s1"}}
	repo.On("GetByID", "A1").Return(assignment, nil)
	repo.On("Update", mock.Anything).Return(nil)

	err := uc.AddSubmissions("A1", []string{"s2", "s3"})
	assert.NoError(t, err)
	repo.AssertCalled(t, "Update", mock.MatchedBy(func(a *model.Assignment) bool {
		return len(a.Submissions) == 3
	}))
}

func TestGetByID_FallbackToRepo(t *testing.T) {
	repo := new(MockRepo)
	cache := new(MockCache)
	pub := new(MockPublisher)
	tx := &MockTx{}

	uc := usecase.NewAssignemntUc(repo, pub, tx.Run, cache)

	expected := &model.Assignment{ID: "A1"}
	repo.On("GetByID", "A1").Return(expected, nil)

	a, err := uc.GetByID("A1")
	assert.NoError(t, err)
	assert.Equal(t, "A1", a.ID)
}
