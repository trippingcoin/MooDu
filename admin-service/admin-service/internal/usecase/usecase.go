package usecase

import (
	"admin/admin-service/internal/cache"
	nats "admin/admin-service/pkg/broker"
	"admin/admin-service/pkg/transactor"
	"admin/pb"
	"context"
)

type AdminRepository interface {
	CreateTranscriptRequest(studentID, purpose string) error
	ViewQueue() ([]*pb.QueueEntry, error)
	JoinQueue(studentID, reason string) error
	RegisterRetake(studentID, courseID, reason string) error
	GetSchedule(studentID string) ([]*pb.ScheduleEntry, error)
	UpdateSchedule(courseID, day, timeStr, room string) error
	SubmitCertificateRequest(studentID, certType, details string) error
}

type AdminUsecase struct {
	repo      AdminRepository
	publisher *nats.NATSPublisher
	callTx    transactor.WithinTransactionFunc
	cache     cache.AdminCache
}

func NewAdminUsecase(repo AdminRepository, publisher *nats.NATSPublisher, callTx transactor.WithinTransactionFunc, cache cache.AdminCache) *AdminUsecase {
	return &AdminUsecase{
		repo:      repo,
		publisher: publisher,
		callTx:    callTx,
		cache:     cache,
	}
}

func (u *AdminUsecase) CreateTranscriptRequest(studentID, purpose string) error {
	txFn := func(ctx context.Context) error {
		if err := u.repo.CreateTranscriptRequest(studentID, purpose); err != nil {
			return err
		}

		event := map[string]string{
			"student_id": studentID,
			"purpose":    purpose,
			"event":      "transcript_requested",
		}
		if err := u.publisher.Publish("admin.transcript.created", event); err != nil {
			return err
		}
		return nil
	}
	return u.callTx(context.Background(), txFn)
}

func (u *AdminUsecase) ViewQueue() (*pb.QueueList, error) {
	if cached, ok := u.cache.GetQueue(); ok {
		return cached, nil
	}

	entries, err := u.repo.ViewQueue()
	if err != nil {
		return nil, err
	}

	queueList := &pb.QueueList{Entries: entries}
	_ = u.cache.SetQueue(queueList)

	return queueList, nil
}

func (u *AdminUsecase) JoinQueue(studentID, reason string) error {
	return u.repo.JoinQueue(studentID, reason)
}

func (u *AdminUsecase) RegisterRetake(studentID, courseID, reason string) error {
	return u.repo.RegisterRetake(studentID, courseID, reason)
}

func (u *AdminUsecase) GetSchedule(studentID string) (*pb.ScheduleResponse, error) {
	if cached, ok := u.cache.GetSchedule(studentID); ok {
		return cached, nil
	}

	entries, err := u.repo.GetSchedule(studentID)
	if err != nil {
		return nil, err
	}

	scheduleResp := &pb.ScheduleResponse{Entries: entries}
	_ = u.cache.SetSchedule(studentID, scheduleResp)

	return scheduleResp, nil
}

func (u *AdminUsecase) UpdateSchedule(courseID, day, timeStr, room string) error {
	return u.repo.UpdateSchedule(courseID, day, timeStr, room)
}

func (u *AdminUsecase) SubmitCertificateRequest(studentID, certType, details string) error {
	return u.repo.SubmitCertificateRequest(studentID, certType, details)
}
