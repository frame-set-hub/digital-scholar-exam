package repository

import (
	"digital-scholar-exam/backend/internal/models"

	"gorm.io/gorm"
)

// EnsureSeedQuestions ใส่ข้อสอบตัวอย่างที่ยังไม่มีใน DB (แหล่งข้อมูลจริงสำหรับ API)
func EnsureSeedQuestions(db *gorm.DB) error {
	qs := seedQuestions()
	for i := range qs {
		var n int64
		if err := db.Model(&models.Question{}).Where("id = ?", qs[i].ID).Count(&n).Error; err != nil {
			return err
		}
		if n > 0 {
			continue
		}
		if err := db.Create(&qs[i]).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedQuestions() []models.Question {
	s1 := "จงหาค่า X"
	s4 := "การคูณพื้นฐาน"
	s5 := "หารลงตัว"
	s7 := "คำย่อทางเว็บ"
	return []models.Question{
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
		{
			ID:              4,
			Prompt:          "5 × 4 = ?",
			Subtitle:        &s4,
			SortOrder:       4,
			CorrectOptionID: "4c",
			Options: []models.Option{
				{ID: "4a", QuestionID: 4, Letter: "A", Text: "18"},
				{ID: "4b", QuestionID: 4, Letter: "B", Text: "19"},
				{ID: "4c", QuestionID: 4, Letter: "C", Text: "20"},
				{ID: "4d", QuestionID: 4, Letter: "D", Text: "24"},
			},
		},
		{
			ID:              5,
			Prompt:          "100 ÷ 4 = ?",
			Subtitle:        &s5,
			SortOrder:       5,
			CorrectOptionID: "5b",
			Options: []models.Option{
				{ID: "5a", QuestionID: 5, Letter: "A", Text: "20"},
				{ID: "5b", QuestionID: 5, Letter: "B", Text: "25"},
				{ID: "5c", QuestionID: 5, Letter: "C", Text: "30"},
				{ID: "5d", QuestionID: 5, Letter: "D", Text: "40"},
			},
		},
		{
			ID:              6,
			Prompt:          "เลขคู่ที่น้อยที่สุดในตัวเลือกคือข้อใด",
			Subtitle:        nil,
			SortOrder:       6,
			CorrectOptionID: "6c",
			Options: []models.Option{
				{ID: "6a", QuestionID: 6, Letter: "A", Text: "3"},
				{ID: "6b", QuestionID: 6, Letter: "B", Text: "5"},
				{ID: "6c", QuestionID: 6, Letter: "C", Text: "8"},
				{ID: "6d", QuestionID: 6, Letter: "D", Text: "11"},
			},
		},
		{
			ID:              7,
			Prompt:          "HTTP ย่อมาจากอะไร",
			Subtitle:        &s7,
			SortOrder:       7,
			CorrectOptionID: "7b",
			Options: []models.Option{
				{ID: "7a", QuestionID: 7, Letter: "A", Text: "Hypertext Markup Language"},
				{ID: "7b", QuestionID: 7, Letter: "B", Text: "Hypertext Transfer Protocol"},
				{ID: "7c", QuestionID: 7, Letter: "C", Text: "High-speed Text Protocol"},
				{ID: "7d", QuestionID: 7, Letter: "D", Text: "Hyperlink Transfer Process"},
			},
		},
	}
}
