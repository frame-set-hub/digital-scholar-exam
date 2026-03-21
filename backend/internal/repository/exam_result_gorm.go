package repository

import (
	"context"

	"digital-scholar-exam/backend/internal/models"

	"gorm.io/gorm"
)

// ExamResultGorm บันทึกผลการสอบ
type ExamResultGorm struct {
	db *gorm.DB
}

// NewExamResultGorm ...
func NewExamResultGorm(db *gorm.DB) *ExamResultGorm {
	return &ExamResultGorm{db: db}
}

// SaveExamResult บันทึกผลสอบลง SQLite
func (r *ExamResultGorm) SaveExamResult(ctx context.Context, res *models.ExamResult) error {
	return r.db.WithContext(ctx).Create(res).Error
}

// GetLeaderboard ดึงอันดับจากมากไปน้อย; คะแนนเท่ากันให้ผู้สอบก่อน (created_at เก่าก่อน)
func (r *ExamResultGorm) GetLeaderboard(ctx context.Context, limit int) ([]models.ExamResult, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 20 {
		limit = 20
	}
	var rows []models.ExamResult
	err := r.db.WithContext(ctx).
		Model(&models.ExamResult{}).
		Order("score DESC").
		Order("created_at ASC").
		Limit(limit).
		Find(&rows).Error
	return rows, err
}
