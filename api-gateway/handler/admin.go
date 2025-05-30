package handler

import (
	admin "api/grpc"
	pb "api/pb/adminpb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.Engine) {
	r.POST("/admin/transcripts", CreateTranscriptRequest)
	r.GET("/admin/queue", ViewQueue)
	r.POST("/admin/queue", JoinQueue)
	r.POST("/admin/retakes", RegisterRetake)
	r.GET("/admin/schedule/:student_id", GetSchedule)
	r.PUT("/admin/schedule", UpdateSchedule)
	r.POST("/admin/certificates", SubmitCertificateRequest)
}

func CreateTranscriptRequest(c *gin.Context) {
	var req pb.TranscriptRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := admin.AdminClient.CreateTranscriptRequest(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "transcript request created"})
}

func ViewQueue(c *gin.Context) {
	resp, err := admin.AdminClient.ViewQueue(c, &pb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Entries)
}

func JoinQueue(c *gin.Context) {
	var req pb.QueueRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := admin.AdminClient.JoinQueue(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "joined queue"})
}

func RegisterRetake(c *gin.Context) {
	var req pb.RetakeRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := admin.AdminClient.RegisterRetake(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "retake registered"})
}

func GetSchedule(c *gin.Context) {
	studentID := c.Param("student_id")
	resp, err := admin.AdminClient.GetSchedule(c, &pb.ScheduleRequest{StudentId: studentID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Entries)
}

func UpdateSchedule(c *gin.Context) {
	var req pb.UpdateScheduleRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := admin.AdminClient.UpdateSchedule(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "schedule updated"})
}

func SubmitCertificateRequest(c *gin.Context) {
	var req pb.CertificateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := admin.AdminClient.SubmitCertificateRequest(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "certificate request submitted"})
}
