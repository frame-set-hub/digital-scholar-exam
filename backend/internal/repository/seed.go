package repository

import (
	"digital-scholar-exam/backend/internal/models"

	"gorm.io/gorm"
)

// SeedQuestionsIfEmpty ใส่ mock ข้อสอบเมื่อยังไม่มีข้อมูล (ตรงกับ frontend MOCK_QUESTIONS)
func SeedQuestionsIfEmpty(db *gorm.DB) error {
	var n int64
	if err := db.Model(&models.Question{}).Count(&n).Error; err != nil {
		return err
	}
	if n > 0 {
		return nil
	}

	s1 := "จงหาค่า X"
	questions := []models.Question{
		{
			ID:              1,
			Prompt:          "ข้อใดต่างจากข้ออื่น",
			Subtitle:        nil,
			SortOrder:       1,
			CorrectOptionID: "1c",
			Options: []models.Option{
				{ID: "1a", QuestionID: 1, Letter: "A", Text: "3"},
				{ID: "1b", QuestionID: 1, Letter: "B", Text: "5"},
				{ID: "1c", QuestionID: 1, Letter: "C", Text: "9"},
				{ID: "1d", QuestionID: 1, Letter: "D", Text: "11"},
			},
		},
		{
			ID:              2,
			Prompt:          "X + 2 = 4",
			Subtitle:        &s1,
			SortOrder:       2,
			CorrectOptionID: "2b",
			Options: []models.Option{
				{ID: "2a", QuestionID: 2, Letter: "A", Text: "1"},
				{ID: "2b", QuestionID: 2, Letter: "B", Text: "2"},
				{ID: "2c", QuestionID: 2, Letter: "C", Text: "3"},
				{ID: "2d", QuestionID: 2, Letter: "D", Text: "4"},
			},
		},
		{
			ID:              3,
			Prompt:          "2 + 2 = ?",
			Subtitle:        nil,
			SortOrder:       3,
			CorrectOptionID: "3b",
			Options: []models.Option{
				{ID: "3a", QuestionID: 3, Letter: "A", Text: "3"},
				{ID: "3b", QuestionID: 3, Letter: "B", Text: "4"},
				{ID: "3c", QuestionID: 3, Letter: "C", Text: "5"},
				{ID: "3d", QuestionID: 3, Letter: "D", Text: "6"},
			},
		},
	}

	for i := range questions {
		if err := db.Create(&questions[i]).Error; err != nil {
			return err
		}
	}
	return nil
}
