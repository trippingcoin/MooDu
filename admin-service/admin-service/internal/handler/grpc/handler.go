package handler

import (
	"context"

	"admin/admin-service/internal/usecase"
	"admin/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AdminHandler struct {
	pb.UnimplementedAdminServiceServer
	uc *usecase.AdminUsecase
}

func NewAdminHandler(uc *usecase.AdminUsecase) *AdminHandler {
	return &AdminHandler{uc: uc}
}

func (h *AdminHandler) CreateTranscriptRequest(ctx context.Context, req *pb.TranscriptRequest) (*pb.Empty, error) {
	if err := h.uc.CreateTranscriptRequest(req.StudentId, req.Purpose); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Empty{}, nil
}

func (h *AdminHandler) ViewQueue(ctx context.Context, _ *pb.Empty) (*pb.QueueList, error) {
	list, err := h.uc.ViewQueue()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return list, nil
}

func (h *AdminHandler) JoinQueue(ctx context.Context, req *pb.QueueRequest) (*pb.Empty, error) {
	if err := h.uc.JoinQueue(req.StudentId, req.Reason); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Empty{}, nil
}

func (h *AdminHandler) RegisterRetake(ctx context.Context, req *pb.RetakeRequest) (*pb.Empty, error) {
	if err := h.uc.RegisterRetake(req.StudentId, req.CourseId, req.Reason); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Empty{}, nil
}

func (h *AdminHandler) GetSchedule(ctx context.Context, req *pb.ScheduleRequest) (*pb.ScheduleResponse, error) {
	schedule, err := h.uc.GetSchedule(req.StudentId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return schedule, nil
}

func (h *AdminHandler) UpdateSchedule(ctx context.Context, req *pb.UpdateScheduleRequest) (*pb.Empty, error) {
	err := h.uc.UpdateSchedule(req.CourseId, req.Day, req.Time, req.Room)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Empty{}, nil
}

func (h *AdminHandler) SubmitCertificateRequest(ctx context.Context, req *pb.CertificateRequest) (*pb.Empty, error) {
	if err := h.uc.SubmitCertificateRequest(req.StudentId, req.CertificateType, req.Details); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Empty{}, nil
}
