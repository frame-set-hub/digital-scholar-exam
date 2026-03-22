package usecase

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"digital-scholar-exam/backend/internal/models"
)

// Exam บริหาร use case ข้อสอบ
type Exam struct {
	questions QuestionStore
	results   ExamResultStore
}

// NewExam ...
func NewExam(q QuestionStore, r ExamResultStore) *Exam {
	return &Exam{questions: q, results: r}
}

// QuestionDTO ส่งออกทาง API (ไม่มีเฉลย)
type QuestionDTO struct {
	ID       uint        `json:"id"`
	Prompt   string      `json:"prompt"`
	Subtitle *string     `json:"subtitle"`
	Options  []OptionDTO `json:"options"`
}

// OptionDTO ...
type OptionDTO struct {
	ID     string `json:"id"`
	Letter string `json:"letter"`
	Text   string `json:"text"`
}

// SubmitResponse ผลหลังส่งข้อสอบ
type SubmitResponse struct {
	CandidateName string `json:"candidateName"`
	Score         int    `json:"score"`
	Total         int    `json:"total"`
}

// LeaderboardEntryDTO อันดับผู้สอบสำหรับ API (ไม่รวมคำตอบดิบ)
type LeaderboardEntryDTO struct {
	Rank          int    `json:"rank"`
	CandidateName string `json:"candidateName"`
	Score         int    `json:"score"`
	Total         int    `json:"total"`
	CreatedAt     string `json:"createdAt"`
}

// GetQuestions ดึงข้อสอบสำหรับหน้า IT 10-1 (ไม่รวม correctOptionId)
func (e *Exam) GetQuestions(ctx context.Context) ([]QuestionDTO, error) {
	qs, err := e.questions.GetQuestions(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]QuestionDTO, 0, len(qs))
	for _, q := range qs {
		opts := make([]OptionDTO, 0, len(q.Options))
		for _, o := range q.Options {
			opts = append(opts, OptionDTO{ID: o.ID, Letter: o.Letter, Text: o.Text})
		}
		out = append(out, QuestionDTO{
			ID:       q.ID,
			Prompt:   q.Prompt,
			Subtitle: q.Subtitle,
			Options:  opts,
		})
	}
	return out, nil
}

// SubmitExam ดึงเฉลยจาก DB ตรวจคำตอบ บวกคะแนนต่อข้อที่ถูก แล้วให้ repository บันทึกชื่อกับคะแนนรวม
func (e *Exam) SubmitExam(ctx context.Context, candidateName string, answers map[string]string) (*SubmitResponse, error) {
	name := strings.TrimSpace(candidateName)
	if name == "" {
		return nil, ErrCandidateNameRequired
	}

	qs, err := e.questions.GetQuestions(ctx)
	if err != nil {
		return nil, err
	}
	score := ScoreAnswers(qs, answers)
	total := len(qs)

	payload, err := json.Marshal(answers)
	if err != nil {
		return nil, err
	}

	exists, err := e.results.CandidateNameExists(ctx, name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrDuplicateCandidateName
	}

	res := &models.ExamResult{
		CandidateName: name,
		Score:           score,
		Total:           total,
		AnswersJSON:     string(payload),
	}
	if err := e.results.SaveExamResult(ctx, res); err != nil {
		return nil, err
	}

	return &SubmitResponse{
		CandidateName: name,
		Score:         score,
		Total:         total,
	}, nil
}

// GetLeaderboard ดึงอันดับจาก repository แล้วส่งเฉพาะฟิลด์ที่จำเป็น
func (e *Exam) GetLeaderboard(ctx context.Context, limit int) ([]LeaderboardEntryDTO, error) {
	limit = normalizeLeaderboardLimit(limit)
	rows, err := e.results.GetLeaderboard(ctx, limit)
	if err != nil {
		return nil, err
	}
	out := make([]LeaderboardEntryDTO, 0, len(rows))
	for i, row := range rows {
		out = append(out, LeaderboardEntryDTO{
			Rank:          i + 1,
			CandidateName: row.CandidateName,
			Score:         row.Score,
			Total:         row.Total,
			CreatedAt:     row.CreatedAt.UTC().Format(time.RFC3339),
		})
	}
	return out, nil
}

func normalizeLeaderboardLimit(limit int) int {
	const defaultLimit = 20
	const maxLimit = 20
	if limit <= 0 {
		return defaultLimit
	}
	if limit > maxLimit {
		return maxLimit
	}
	return limit
}

// ScoreAnswers นับคะแนนจากคำตอบ (คีย์ของ map เป็น string ของ question id เช่น "1","2")
func ScoreAnswers(questions []models.Question, answers map[string]string) int {
	score := 0
	for _, q := range questions {
		key := strconv.FormatUint(uint64(q.ID), 10)
		if answers[key] == q.CorrectOptionID {
			score++
		}
	}
	return score
}
