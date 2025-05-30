package usecase_test

import (
	"admin/admin-service/internal/usecase"
	"admin/pb"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ---- Mocks ----

type MockRepo struct{ mock.Mock }
type MockPublisher struct{ mock.Mock }
type MockTx struct{}
type MockCache struct{ mock.Mock }

func (m *MockRepo) CreateTranscriptRequest(sid, purpose string) error {
	args := m.Called(sid, purpose)
	return args.Error(0)
}
func (m *MockRepo) ViewQueue() ([]*pb.QueueEntry, error) {
	args := m.Called()
	return args.Get(0).([]*pb.QueueEntry), args.Error(1)
}
func (m *MockRepo) JoinQueue(sid, reason string) error {
	args := m.Called(sid, reason)
	return args.Error(0)
}
func (m *MockRepo) RegisterRetake(sid, cid, reason string) error {
	args := m.Called(sid, cid, reason)
	return args.Error(0)
}
func (m *MockRepo) GetSchedule(sid string) ([]*pb.ScheduleEntry, error) {
	args := m.Called(sid)
	return args.Get(0).([]*pb.ScheduleEntry), args.Error(1)
}
func (m *MockRepo) UpdateSchedule(cid, day, timeStr, room string) error {
	args := m.Called(cid, day, timeStr, room)
	return args.Error(0)
}
func (m *MockRepo) SubmitCertificateRequest(sid, typ, details string) error {
	args := m.Called(sid, typ, details)
	return args.Error(0)
}

func (m *MockPublisher) Publish(subj string, msg interface{}) error {
	return nil
}

func (tx *MockTx) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}

func (m *MockCache) GetQueue() (*pb.QueueList, bool)                            { return nil, false }
func (m *MockCache) SetQueue(*pb.QueueList) error                               { return nil }
func (m *MockCache) GetSchedule(studentID string) (*pb.ScheduleResponse, bool)  { return nil, false }
func (m *MockCache) SetSchedule(studentID string, s *pb.ScheduleResponse) error { return nil }

// ---- Tests ----

func TestCreateTranscriptRequest(t *testing.T) {
	repo := new(MockRepo)
	cache := new(MockCache)
	pub := new(MockPublisher)
	tx := &MockTx{}

	uc := usecase.NewAdminUsecase(repo, pub, tx.Run, cache)

	repo.On("CreateTranscriptRequest", "123", "job").Return(nil)

	err := uc.CreateTranscriptRequest("123", "job")
	assert.NoError(t, err)
	repo.AssertCalled(t, "CreateTranscriptRequest", "123", "job")
}

func TestViewQueue_CacheMiss(t *testing.T) {
	repo := new(MockRepo)
	cache := new(MockCache)
	pub := new(MockPublisher)
	tx := &MockTx{}
	uc := usecase.NewAdminUsecase(repo, pub, tx.Run, cache)

	repo.On("ViewQueue").Return([]*pb.QueueEntry{{StudentId: "1"}}, nil)

	result, err := uc.ViewQueue()
	assert.NoError(t, err)
	assert.Len(t, result.Entries, 1)
}

func TestGetSchedule_CacheMiss(t *testing.T) {
	repo := new(MockRepo)
	cache := new(MockCache)
	pub := new(MockPublisher)
	tx := &MockTx{}
	uc := usecase.NewAdminUsecase(repo, pub, tx.Run, cache)

	repo.On("GetSchedule", "1").Return([]*pb.ScheduleEntry{{CourseId: "c1"}}, nil)

	resp, err := uc.GetSchedule("1")
	assert.NoError(t, err)
	assert.Equal(t, "c1", resp.Entries[0].CourseId)
}

func TestJoinQueue(t *testing.T) {
	repo := new(MockRepo)
	cache := new(MockCache)
	pub := new(MockPublisher)
	tx := &MockTx{}
	uc := usecase.NewAdminUsecase(repo, pub, tx.Run, cache)

	repo.On("JoinQueue", "1", "delay").Return(nil)

	err := uc.JoinQueue("1", "delay")
	assert.NoError(t, err)
}

func TestRegisterRetake(t *testing.T) {
	repo := new(MockRepo)
	cache := new(MockCache)
	pub := new(MockPublisher)
	tx := &MockTx{}
	uc := usecase.NewAdminUsecase(repo, pub, tx.Run, cache)

	repo.On("RegisterRetake", "1", "math", "failed").Return(nil)

	err := uc.RegisterRetake("1", "math", "failed")
	assert.NoError(t, err)
}

func TestSubmitCertificateRequest(t *testing.T) {
	repo := new(MockRepo)
	cache := new(MockCache)
	pub := new(MockPublisher)
	tx := &MockTx{}
	uc := usecase.NewAdminUsecase(repo, pub, tx.Run, cache)

	repo.On("SubmitCertificateRequest", "1", "study", "urgent").Return(nil)
	err := uc.SubmitCertificateRequest("1", "study", "urgent")
	assert.NoError(t, err)
}
