package handler

import (
	"net/http"
	"strconv"

	"digital-scholar-exam/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

// ExamHTTP รับ HTTP สำหรับข้อสอบ
type ExamHTTP struct {
	uc *usecase.Exam
}

// NewExamHTTP ...
func NewExamHTTP(uc *usecase.Exam) *ExamHTTP {
	return &ExamHTTP{uc: uc}
}

// GetQuestions GET /api/questions
func (h *ExamHTTP) GetQuestions(c *gin.Context) {
	items, err := h.uc.GetQuestions(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load questions"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"questions": items})
}

// SubmitBody JSON สำหรับ POST /api/submit
type SubmitBody struct {
	CandidateName string            `json:"candidateName" binding:"required"`
	Answers       map[string]string `json:"answers" binding:"required"`
}

// Submit POST /api/submit
func (h *ExamHTTP) Submit(c *gin.Context) {
	var body SubmitBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	if len(body.Answers) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "answers required"})
		return
	}

	res, err := h.uc.SubmitExam(c.Request.Context(), body.CandidateName, body.Answers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetLeaderboard GET /api/leaderboard
func (h *ExamHTTP) GetLeaderboard(c *gin.Context) {
	limit := 0
	if q := c.Query("limit"); q != "" {
		if n, err := strconv.Atoi(q); err == nil {
			limit = n
		}
	}
	entries, err := h.uc.GetLeaderboard(c.Request.Context(), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load leaderboard"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"entries": entries})
}
