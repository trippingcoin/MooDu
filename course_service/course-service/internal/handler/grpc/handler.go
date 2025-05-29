package grpcH

import (
	"context"
	"log"

	"cs/course-service/internal/model"
	"cs/course-service/internal/usecase"
	pb "cs/pb_cs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HandlerCourse struct {
	pb.UnimplementedCourseServiceServer
	usecase *usecase.CourseUsecase
}

func NewCO(usecase *usecase.CourseUsecase) *HandlerCourse {
	return &HandlerCourse{usecase: usecase}
}

func (h *HandlerCourse) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CourseResponse, error) {
	c := &model.Course{
		Title:       req.Title,
		Description: req.Description,
		TeacherID:   req.TeacherId,
	}
	err := h.usecase.Create(ctx, c)
	if err != nil {
		return nil, err
	}
	return &pb.CourseResponse{
		Course: &pb.Course{
			Id:          c.ID,
			Title:       c.Title,
			Description: c.Description,
			TeacherId:   c.TeacherID,
		},
	}, nil
}

func (h *HandlerCourse) UpdateCourse(ctx context.Context, req *pb.UpdateCourseRequest) (*pb.UpdateCourseResponse, error) {
	if _, err := primitive.ObjectIDFromHex(req.GetId()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid course ID")
	}

	course := &model.Course{
		ID:          req.GetId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		TeacherID:   req.GetInstructor(),
	}

	if err := h.usecase.UpdateCourse(course); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateCourseResponse{
		Course: &pb.Course{
			Id:          course.ID,
			Title:       course.Title,
			Description: course.Description,
			TeacherId:   course.TeacherID,
		},
	}, nil
}

func (h *HandlerCourse) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.CourseResponse, error) {
	c, err := h.usecase.GetByID(req.Id)
	if err != nil {
		log.Printf("Error retrieving course with ID %s: %v", req.Id, err)
		return nil, err
	}
	log.Printf("Retrieved course: %+v", c)
	return &pb.CourseResponse{
		Course: &pb.Course{
			Id:          c.ID,
			Title:       c.Title,
			Description: c.Description,
			TeacherId:   c.TeacherID,
		},
	}, nil
}

func (h *HandlerCourse) ListCourses(ctx context.Context, _ *pb.Empty) (*pb.CourseList, error) {
	courses, err := h.usecase.List()
	if err != nil {
		return nil, err
	}
	var res []*pb.Course
	for _, c := range courses {
		res = append(res, &pb.Course{
			Id:          c.ID,
			Title:       c.Title,
			Description: c.Description,
			TeacherId:   c.TeacherID,
		})
	}
	return &pb.CourseList{Courses: res}, nil
}

func (h *HandlerCourse) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	if _, err := primitive.ObjectIDFromHex(req.Id); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid course ID")
	}

	err := h.usecase.Delete(req.Id)
	if err != nil {
		if err.Error() == "course not found" {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteCourseResponse{
		Message: "course deleted successfully",
	}, nil
}
