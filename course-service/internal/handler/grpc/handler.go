package grpcH

import (
	"context"
	"log"

	"cs/course-service/internal/model"
	"cs/course-service/internal/usecase"
	pb "cs/pb"
)

type Handler struct {
	pb.UnimplementedCourseServiceServer
	usecase *usecase.CourseUsecase
}

func New(usecase *usecase.CourseUsecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CourseResponse, error) {
	c := &model.Course{
		Title:       req.Title,
		Description: req.Description,
		TeacherID:   req.TeacherId,
	}
	err := h.usecase.Create(c)
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

func (h *Handler) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.CourseResponse, error) {
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

func (h *Handler) ListCourses(ctx context.Context, _ *pb.Empty) (*pb.CourseList, error) {
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
