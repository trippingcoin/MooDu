package grpcH

import (
	"context"

	"cs/course-service/internal/model"
	"cs/course-service/internal/usecase"
	pb "cs/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HandlerAssignment struct {
	pb.UnimplementedAssignmentServiceServer
	usecase *usecase.AssignmentUsecase
}

func NewAS(usecase *usecase.AssignmentUsecase) *HandlerAssignment {
	return &HandlerAssignment{usecase: usecase}
}

func (h *HandlerAssignment) CreateAssignment(ctx context.Context, req *pb.CreateAssignmentRequest) (*pb.AssignmentResponse, error) {
	a := &model.Assignment{
		Title:       req.Title,
		Description: req.Description,
		CourseID:    req.CourseId,
		DueDate:     req.Deadline,
	}
	err := h.usecase.Create(ctx, a)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AssignmentResponse{
		Assignment: &pb.Assignment{
			Id:          a.ID,
			Title:       a.Title,
			Description: a.Description,
			CourseId:    a.CourseID,
			Deadline:    a.DueDate,
		},
	}, nil
}

func (h *HandlerAssignment) UpdateAssignment(ctx context.Context, req *pb.UpdateAssignmentRequest) (*pb.AssignmentResponse, error) {
	id := req.GetId()
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid assignment ID")
	}

	a := &model.Assignment{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		CourseID:    req.CourseId,
		DueDate:     req.Deadline,
	}
	err = h.usecase.Update(a)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AssignmentResponse{
		Assignment: &pb.Assignment{
			Id:          a.ID,
			Title:       a.Title,
			Description: a.Description,
			CourseId:    a.CourseID,
			Deadline:    a.DueDate,
		},
	}, nil
}

func (h *HandlerAssignment) DeleteAssignment(ctx context.Context, req *pb.DeleteAssignmentRequest) (*pb.Empty, error) {
	err := h.usecase.Delete(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Empty{}, nil
}

func (h *HandlerAssignment) GetAssignment(ctx context.Context, req *pb.GetAssignmentRequest) (*pb.AssignmentResponse, error) {
	a, err := h.usecase.GetByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &pb.AssignmentResponse{
		Assignment: &pb.Assignment{
			Id:          a.ID,
			Title:       a.Title,
			Description: a.Description,
			CourseId:    a.CourseID,
			Deadline:    a.DueDate,
		},
	}, nil
}

func (h *HandlerAssignment) ListAssignments(ctx context.Context, _ *pb.Empty) (*pb.AssignmentList, error) {
	assignments, err := h.usecase.List()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var res []*pb.Assignment
	for _, a := range assignments {
		res = append(res, &pb.Assignment{
			Id:          a.ID,
			Title:       a.Title,
			Description: a.Description,
			CourseId:    a.CourseID,
			Deadline:    a.DueDate,
		})
	}
	return &pb.AssignmentList{Assignments: res}, nil
}

func (h *HandlerAssignment) AddSubmissions(ctx context.Context, req *pb.AddSubmissionsRequest) (*pb.AssignmentResponse, error) {
	if req.AssignmentId == "" || len(req.SubmissionIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "assignment ID and submission IDs are required")
	}

	err := h.usecase.AddSubmissions(req.AssignmentId, req.SubmissionIds)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	assignment, err := h.usecase.GetByID(req.AssignmentId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.AssignmentResponse{
		Assignment: &pb.Assignment{
			Id:          assignment.ID,
			Title:       assignment.Title,
			Description: assignment.Description,
			CourseId:    assignment.CourseID,
			Deadline:    assignment.DueDate,
		},
	}, nil
}
