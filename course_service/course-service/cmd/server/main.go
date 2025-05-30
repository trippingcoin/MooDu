package main

import (
	cache "cs/course-service/internal/cache/inmemory"
	cacheR "cs/course-service/internal/cache/redis"

	grpcH "cs/course-service/internal/handler/grpc"
	"cs/course-service/internal/repo/mongodb"
	"cs/course-service/internal/usecase"
	"cs/course-service/pkg/broker/nats"
	pb "cs/pb"
	"log"
	"net"

	"cs/course-service/pkg/mongo"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func main() {
	db, err := mongo.ConnectMongo()
	if err != nil {
		log.Println("mongo: %w", err)
	}

	repo := mongodb.NewCourseRepo(db)
	repoAS := mongodb.NewAS(db)

	courseCache := cache.NewInMemoryCourseCache()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	assignmentCache := cacheR.NewAssignmentRedisCache(rdb)

	transactor := mongo.NewTransactor(db.Client())

	publisher, err := nats.NewNATSPublisher("nats://localhost:4222")
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	log.Println("Succesful to connect to Nats")

	useCaseAs := usecase.NewAssignemntUc(repoAS, publisher, transactor.WithinTransaction, assignmentCache)
	handlerAS := grpcH.NewAS(useCaseAs)
	useCase := usecase.New(repo, publisher, transactor.WithinTransaction, courseCache)
	handler := grpcH.NewCO(useCase)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterCourseServiceServer(server, handler)
	pb.RegisterAssignmentServiceServer(server, handlerAS)

	log.Println("CourseService gRPC server is running on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
