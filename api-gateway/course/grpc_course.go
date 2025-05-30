package course

import (
	pb "api/pb"
	"log"

	"google.golang.org/grpc"
)

var CourseClient pb.CourseServiceClient

func InitCourseClient(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to course service: %v", err)
	}
	CourseClient = pb.NewCourseServiceClient(conn)
}

var AssignmentClient pb.AssignmentServiceClient

func InitAssignmentClient(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to assignment service: %v", err)
	}
	AssignmentClient = pb.NewAssignmentServiceClient(conn)
}
