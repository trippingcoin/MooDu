package main

import (
	grpcH "cs/course-service/internal/handler/grpc"
	"cs/course-service/internal/repo/mongodb"
	"cs/course-service/internal/usecase"
	"cs/pb"
	"log"
	"net"

	"cs/course-service/pkg/mongo"

	"google.golang.org/grpc"
)

func main() {
	db := mongo.ConnectMongo()
	repo := mongodb.NewCourseRepo(db)
	useCase := usecase.New(repo)
	handler := grpcH.New(useCase)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterCourseServiceServer(server, handler)

	log.Println("CourseService gRPC server is running on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
