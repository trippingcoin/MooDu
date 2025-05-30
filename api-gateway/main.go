package main

import (
	"api/course"
	"api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	course.InitCourseClient("localhost:50051")
	course.InitAssignmentClient("localhost:50051")
	handler.RegisterCourseRoutes(r)
	handler.RegisterAssignmentRoutes(r)

	r.Run(":8080")
}
