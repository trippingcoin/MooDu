package handler

import (
	course "api/grpc"
	"net/http"

	pb "api/pb/coursepb"

	"github.com/gin-gonic/gin"
)

func RegisterCourseRoutes(r *gin.Engine) {
	r.POST("/courses", CreateCourse)
	r.PUT("/courses/:id", UpdateCourse)
	r.GET("/courses/:id", GetCourse)
	r.GET("/courses", ListCourses)
	r.DELETE("/courses/:id", DeleteCourse)
}

func CreateCourse(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		TeacherID   string `json:"teacher_id"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := course.CourseClient.CreateCourse(c, &pb.CreateCourseRequest{
		Title:       req.Title,
		Description: req.Description,
		TeacherId:   req.TeacherID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp.Course)
}

func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Instructor  string `json:"instructor"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := course.CourseClient.UpdateCourse(c, &pb.UpdateCourseRequest{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		Instructor:  req.Instructor,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Course)
}

func GetCourse(c *gin.Context) {
	id := c.Param("id")
	resp, err := course.CourseClient.GetCourse(c, &pb.GetCourseRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Course)
}

func ListCourses(c *gin.Context) {
	resp, err := course.CourseClient.ListCourses(c, &pb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Courses)
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	resp, err := course.CourseClient.DeleteCourse(c, &pb.DeleteCourseRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": resp.Message})
}
