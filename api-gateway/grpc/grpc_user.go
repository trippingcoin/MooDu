package grpc

import (
	pb "api/pb/userpb"
	"log"

	"google.golang.org/grpc"
)

var UserClient pb.UserServiceClient

func InitUserClient(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	UserClient = pb.NewUserServiceClient(conn)
}
