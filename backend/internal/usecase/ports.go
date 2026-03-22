package usecase

import (
	"context"

	"digital-scholar-exam/backend/internal/models"
)

// QuestionStore ดึงรายการข้อสอบพร้อมเฉลยในฐานข้อมูล (repository implement)
type QuestionStore interface {
	GetQuestions(ctx context.Context) ([]models.Question, error)
}

// ExamResultStore บันทึกผลการสอบลง SQLite
type ExamResultStore interface {
	CandidateNameExists(ctx context.Context, candidateName string) (bool, error)
	SaveExamResult(ctx context.Context, r *models.ExamResult) error
	GetLeaderboard(ctx context.Context, limit int) ([]models.ExamResult, error)
	// CandidateRank คำนวณอันดับรวม (1 = สูงสุด) ตาม score DESC แล้ว created_at ASC — ไม่พบชื่อแล้ว found=false
	CandidateRank(ctx context.Context, candidateName string) (rank int, row models.ExamResult, found bool, err error)
}
