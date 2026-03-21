package repository

import (
	"context"

	"digital-scholar-exam/backend/internal/models"

	"gorm.io/gorm"
)

// QuestionGorm ค้นหาข้อสอบด้วย GORM
type QuestionGorm struct {
	db *gorm.DB
}

// NewQuestionGorm ...
func NewQuestionGorm(db *gorm.DB) *QuestionGorm {
	return &QuestionGorm{db: db}
}

// GetQuestions ดึงคำถามทั้งหมดเรียงตาม SortOrder และ preload ตัวเลือก (รวม CorrectOptionID สำหรับตรวจคะแนน)
func (r *QuestionGorm) GetQuestions(ctx context.Context) ([]models.Question, error) {
	var out []models.Question
	err := r.db.WithContext(ctx).
		Preload("Options").
		Order("sort_order ASC").
		Find(&out).Error
	return out, err
}
