package main

import (
	"cs/course-service/internal/cache"
	grpcH "cs/course-service/internal/handler/grpc"
	"cs/course-service/internal/repo/mongodb"
	"cs/course-service/internal/usecase"
	"cs/course-service/pkg/broker/nats"
	"cs/pb"
	"log"
	"net"

	"cs/course-service/pkg/mongo"

	"google.golang.org/grpc"
)

func main() {

	db, err := mongo.ConnectMongo()
	if err != nil {
		log.Println("mongo: %w", err)
	}
	repo := mongodb.NewCourseRepo(db)

	courseCache := cache.NewInMemoryCourseCache()

	transactor := mongo.NewTransactor(db.Client())

	publisher, err := nats.NewNATSPublisher("nats://localhost:4222")
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	log.Println("Succesful to connect to Nats")

	useCase := usecase.New(repo, publisher, transactor.WithinTransaction, courseCache)
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
