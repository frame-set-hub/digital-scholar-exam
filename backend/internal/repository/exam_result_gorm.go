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
