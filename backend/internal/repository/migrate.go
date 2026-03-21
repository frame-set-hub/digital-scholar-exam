package repository

import (
	"digital-scholar-exam/backend/internal/models"

	"gorm.io/gorm"
)

// AutoMigrate สร้างตารางจาก models
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Question{},
		&models.Option{},
		&models.ExamResult{},
	)
}
