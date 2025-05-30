package main

import (
	cache "admin/admin-service/internal/cache/redis"
	handler "admin/admin-service/internal/handler/grpc"
	repo "admin/admin-service/internal/repo/mongodb"
	"admin/admin-service/internal/usecase"
	nats "admin/admin-service/pkg/broker"
	"admin/admin-service/pkg/mongo"
	"admin/pb"
	"log"
	"net"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func main() {
	db, err := mongo.ConnectMongo()
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}

	mongoRepo := repo.NewMongoAdminRepo(db)
	transactor := mongo.NewTransactor(db.Client())

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	adminCache := cache.NewRedisAdminCache(rdb)

	publisher, err := nats.NewNATSPublisher("nats://localhost:4222")
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}

	adminUC := usecase.NewAdminUsecase(mongoRepo, publisher, transactor.WithinTransaction, adminCache)

	adminHandler := handler.NewAdminHandler(adminUC)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServiceServer(grpcServer, adminHandler)

	log.Println("AdminService gRPC server is running on :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
