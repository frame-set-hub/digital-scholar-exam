package models

// Question คำถามหนึ่งข้อ (มีตัวเลือกหลายข้อ และเก็บ id ของข้อที่ถูกต้อง)
type Question struct {
	ID              uint     `gorm:"primaryKey"`
	Prompt          string   `gorm:"size:512;not null"`
	Subtitle        *string  `gorm:"size:512"`
	SortOrder       int      `gorm:"not null;index"`
	CorrectOptionID string   `gorm:"size:32;not null"`
	Options         []Option `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE"`
}

// Option ตัวเลือกคำตอบ (primary key เป็น string ตรงกับฝั่ง frontend เช่น 1a, 2b)
type Option struct {
	ID         string `gorm:"primaryKey;size:32"`
	QuestionID uint   `gorm:"not null;index"`
	Letter     string `gorm:"size:4;not null"`
	Text       string `gorm:"size:512;not null"`
}
