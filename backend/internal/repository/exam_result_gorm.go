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

// CandidateNameExists ตรวจว่ามีผลสอบด้วยชื่อนี้แล้วหรือไม่ (เทียบตรงหลัง trim ฝั่ง use case)
func (r *ExamResultGorm) CandidateNameExists(ctx context.Context, candidateName string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.ExamResult{}).
		Where("candidate_name = ?", candidateName).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
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
