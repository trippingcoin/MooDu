package main

import (
	"api/grpc"
	"api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	grpc.InitCourseClient("localhost:50051")
	grpc.InitAssignmentClient("localhost:50051")
	grpc.InitUserClient("localhost:50052")
	grpc.InitAdminClient("localhost:50053")

	handler.RegisterCourseRoutes(r)
	handler.RegisterAssignmentRoutes(r)
	handler.RegisterUserRoutes(r)
	handler.RegisterAdminRoutes(r)

	r.Run(":8080")
}
