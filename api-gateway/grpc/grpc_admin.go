package grpc

import (
	pb "api/pb/adminpb"
	"log"

	"google.golang.org/grpc"
)

var AdminClient pb.AdminServiceClient

func InitAdminClient(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to admin service: %v", err)
	}
	AdminClient = pb.NewAdminServiceClient(conn)
}
