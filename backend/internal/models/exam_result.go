package models

import "time"

// ExamResult บันทึกผลการสอบ
type ExamResult struct {
	ID            uint   `gorm:"primaryKey"`
	CandidateName string `gorm:"size:255;not null"`
	Score         int    `gorm:"not null"`
	Total         int    `gorm:"not null"`
	AnswersJSON   string `gorm:"type:text"` // optional: เก็บ payload ของคำตอบ
	CreatedAt     time.Time
}
