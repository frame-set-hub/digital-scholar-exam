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
	SaveExamResult(ctx context.Context, r *models.ExamResult) error
}
