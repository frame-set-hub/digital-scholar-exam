package handler

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes ลงทะเบียน API ภายใต้ /api
func RegisterRoutes(r *gin.Engine, exam *ExamHTTP) {
	api := r.Group("/api")
	{
		api.GET("/questions", exam.GetQuestions)
		api.POST("/submit", exam.Submit)
		api.GET("/leaderboard", exam.GetLeaderboard)
	}
}
