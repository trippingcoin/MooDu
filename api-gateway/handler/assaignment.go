package handler

import (
	assignment "api/course"
	"net/http"

	pb "api/pb"

	"github.com/gin-gonic/gin"
)

func RegisterAssignmentRoutes(r *gin.Engine) {
	r.POST("/assignments", CreateAssignment)
	r.PUT("/assignments/:id", UpdateAssignment)
	r.DELETE("/assignments/:id", DeleteAssignment)
	r.GET("/assignments/:id", GetAssignment)
	r.GET("/assignments", ListAssignments)
	r.POST("/assignments/:id/submit", AddSubmissions)
}

func CreateAssignment(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		CourseID    string `json:"course_id"`
		Deadline    string `json:"deadline"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := assignment.AssignmentClient.CreateAssignment(c, &pb.CreateAssignmentRequest{
		Title:       req.Title,
		Description: req.Description,
		CourseId:    req.CourseID,
		Deadline:    req.Deadline,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp.Assignment)
}

func UpdateAssignment(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		CourseID    string `json:"course_id"`
		Deadline    string `json:"deadline"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := assignment.AssignmentClient.UpdateAssignment(c, &pb.UpdateAssignmentRequest{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		CourseId:    req.CourseID,
		Deadline:    req.Deadline,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Assignment)
}

func DeleteAssignment(c *gin.Context) {
	id := c.Param("id")
	_, err := assignment.AssignmentClient.DeleteAssignment(c, &pb.DeleteAssignmentRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "assignment deleted"})
}

func GetAssignment(c *gin.Context) {
	id := c.Param("id")
	resp, err := assignment.AssignmentClient.GetAssignment(c, &pb.GetAssignmentRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Assignment)
}

func ListAssignments(c *gin.Context) {
	resp, err := assignment.AssignmentClient.ListAssignments(c, &pb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Assignments)
}

func AddSubmissions(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		SubmissionIDs []string `json:"submission_ids"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := assignment.AssignmentClient.AddSubmissions(c, &pb.AddSubmissionsRequest{
		AssignmentId:  id,
		SubmissionIds: req.SubmissionIDs,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Assignment)
}
