package handler

import (
	"errors"
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
		switch {
		case errors.Is(err, usecase.ErrCandidateNameRequired):
			c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณากรอกชื่อผู้สอบ"})
		case errors.Is(err, usecase.ErrDuplicateCandidateName):
			c.JSON(http.StatusConflict, gin.H{"error": "ชื่อนี้ถูกใช้ส่งข้อสอบแล้ว — กรุณาใช้ชื่ออื่น"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit"})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}

// leaderboardResponse — yourEntry อยู่ก่อน entries ใน JSON เพื่อให้ดูด้วย curl | head ได้ (entries ยาวมาก)
type leaderboardResponse struct {
	YourEntry *usecase.LeaderboardYourEntryDTO `json:"yourEntry"`
	Entries   []usecase.LeaderboardEntryDTO  `json:"entries"`
}

// GetLeaderboard GET /api/leaderboard
func (h *ExamHTTP) GetLeaderboard(c *gin.Context) {
	limit := 0
	if q := c.Query("limit"); q != "" {
		if n, err := strconv.Atoi(q); err == nil {
			limit = n
		}
	}
	forCandidate := c.Query("forCandidate")
	entries, your, err := h.uc.GetLeaderboard(c.Request.Context(), limit, forCandidate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load leaderboard"})
		return
	}
	// yourEntry เป็น null เมื่อไม่ส่ง forCandidate หรือไม่พบชื่อ — ส่งคีย์เสมอเพื่อให้ curl/FE ดีบักได้ชัด
	c.JSON(http.StatusOK, leaderboardResponse{YourEntry: your, Entries: entries})
}
