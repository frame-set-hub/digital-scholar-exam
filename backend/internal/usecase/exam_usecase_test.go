package usecase_test

import (
	"context"
	"testing"
	"time"

	"digital-scholar-exam/backend/internal/models"
	"digital-scholar-exam/backend/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQuestionStore จำลอง Repository ข้อสอบ (testify/mock)
type MockQuestionStore struct {
	mock.Mock
}

func (m *MockQuestionStore) GetQuestions(ctx context.Context) ([]models.Question, error) {
	args := m.Called(ctx)
	v, _ := args.Get(0).([]models.Question)
	return v, args.Error(1)
}

// MockExamResultStore จำลอง Repository บันทึกผล
type MockExamResultStore struct {
	mock.Mock
}

func (m *MockExamResultStore) CandidateNameExists(ctx context.Context, name string) (bool, error) {
	args := m.Called(ctx, name)
	return args.Bool(0), args.Error(1)
}

func (m *MockExamResultStore) SaveExamResult(ctx context.Context, r *models.ExamResult) error {
	args := m.Called(ctx, r)
	return args.Error(0)
}

func (m *MockExamResultStore) GetLeaderboard(ctx context.Context, limit int) ([]models.ExamResult, error) {
	args := m.Called(ctx, limit)
	v, _ := args.Get(0).([]models.ExamResult)
	return v, args.Error(1)
}

func sampleQuestions() []models.Question {
	s := "จงหาค่า X"
	return []models.Question{
		{ID: 1, CorrectOptionID: "1c"},
		{ID: 2, Subtitle: &s, CorrectOptionID: "2b"},
		{ID: 3, CorrectOptionID: "3b"},
	}
}

func TestScoreAnswers_AllCorrect_FullScore(t *testing.T) {
	qs := sampleQuestions()
	ans := map[string]string{"1": "1c", "2": "2b", "3": "3b"}
	assert.Equal(t, 3, usecase.ScoreAnswers(qs, ans))
}

func TestScoreAnswers_AllWrong_ZeroScore(t *testing.T) {
	qs := sampleQuestions()
	ans := map[string]string{"1": "1a", "2": "2a", "3": "3a"}
	assert.Equal(t, 0, usecase.ScoreAnswers(qs, ans))
}

func TestScoreAnswers_Partial(t *testing.T) {
	qs := sampleQuestions()
	ans := map[string]string{"1": "1c", "2": "2a", "3": "3b"}
	assert.Equal(t, 2, usecase.ScoreAnswers(qs, ans))
}

func TestScoreAnswers_WrongKeysIgnoredForUnknownQuestion(t *testing.T) {
	qs := []models.Question{{ID: 1, CorrectOptionID: "1c"}}
	ans := map[string]string{"1": "1c", "99": "x"}
	assert.Equal(t, 1, usecase.ScoreAnswers(qs, ans))
}

func TestExam_SubmitExam_FullScore(t *testing.T) {
	mq := new(MockQuestionStore)
	mr := new(MockExamResultStore)
	mq.On("GetQuestions", mock.Anything).Return(sampleQuestions(), nil)
	mr.On("CandidateNameExists", mock.Anything, "Alice").Return(false, nil)
	mr.On("SaveExamResult", mock.Anything, mock.MatchedBy(func(r *models.ExamResult) bool {
		return r.CandidateName == "Alice" && r.Score == 3 && r.Total == 3
	})).Return(nil)

	ex := usecase.NewExam(mq, mr)
	res, err := ex.SubmitExam(context.Background(), "Alice", map[string]string{
		"1": "1c", "2": "2b", "3": "3b",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Alice", res.CandidateName)
	assert.Equal(t, 3, res.Score)
	assert.Equal(t, 3, res.Total)

	mq.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestExam_SubmitExam_ZeroScore(t *testing.T) {
	mq := new(MockQuestionStore)
	mr := new(MockExamResultStore)
	mq.On("GetQuestions", mock.Anything).Return(sampleQuestions(), nil)
	mr.On("CandidateNameExists", mock.Anything, "Bob").Return(false, nil)
	mr.On("SaveExamResult", mock.Anything, mock.MatchedBy(func(r *models.ExamResult) bool {
		return r.CandidateName == "Bob" && r.Score == 0 && r.Total == 3
	})).Return(nil)

	ex := usecase.NewExam(mq, mr)
	res, err := ex.SubmitExam(context.Background(), "Bob", map[string]string{
		"1": "1a", "2": "2a", "3": "3a",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Bob", res.CandidateName)
	assert.Equal(t, 0, res.Score)
	assert.Equal(t, 3, res.Total)

	mq.AssertExpectations(t)
	mr.AssertExpectations(t)
}

// --- Edge cases: full / zero / missing answers / invalid option IDs ---

func TestScoreAnswers_MissingAnswers_OnlySubmittedRowsGraded(t *testing.T) {
	// ข้อสอบ 2 ข้อ ส่งมาแค่ข้อเดียว (ถูก) — ข้อที่ไม่ส่งถือว่าไม่ตรงเฉลย ไม่ panic
	two := sampleQuestions()[:2]
	ans := map[string]string{"1": "1c"}
	assert.Equal(t, 1, usecase.ScoreAnswers(two, ans))
}

func TestScoreAnswers_InvalidOptionID_NoPoint(t *testing.T) {
	qs := []models.Question{{ID: 1, CorrectOptionID: "1c"}}
	ans := map[string]string{"1": "not-a-real-option-id"}
	assert.Equal(t, 0, usecase.ScoreAnswers(qs, ans))
}

func TestExam_SubmitExam_MissingAnswers_PartialScore(t *testing.T) {
	mq := new(MockQuestionStore)
	mr := new(MockExamResultStore)
	two := sampleQuestions()[:2]
	mq.On("GetQuestions", mock.Anything).Return(two, nil)
	mr.On("CandidateNameExists", mock.Anything, "Partial").Return(false, nil)
	mr.On("SaveExamResult", mock.Anything, mock.MatchedBy(func(r *models.ExamResult) bool {
		return r.CandidateName == "Partial" && r.Score == 1 && r.Total == 2
	})).Return(nil)

	ex := usecase.NewExam(mq, mr)
	res, err := ex.SubmitExam(context.Background(), "Partial", map[string]string{"1": "1c"})
	assert.NoError(t, err)
	assert.Equal(t, 1, res.Score)
	assert.Equal(t, 2, res.Total)

	mq.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestExam_SubmitExam_EmptyNameAfterTrim(t *testing.T) {
	ex := usecase.NewExam(new(MockQuestionStore), new(MockExamResultStore))
	res, err := ex.SubmitExam(context.Background(), "   ", map[string]string{"1": "1c"})
	assert.ErrorIs(t, err, usecase.ErrCandidateNameRequired)
	assert.Nil(t, res)
}

func TestExam_SubmitExam_DuplicateName(t *testing.T) {
	mq := new(MockQuestionStore)
	mr := new(MockExamResultStore)
	mq.On("GetQuestions", mock.Anything).Return(sampleQuestions(), nil)
	mr.On("CandidateNameExists", mock.Anything, "Alice").Return(true, nil)

	ex := usecase.NewExam(mq, mr)
	res, err := ex.SubmitExam(context.Background(), "Alice", map[string]string{
		"1": "1c", "2": "2b", "3": "3b",
	})
	assert.ErrorIs(t, err, usecase.ErrDuplicateCandidateName)
	assert.Nil(t, res)
	mr.AssertNotCalled(t, "SaveExamResult", mock.Anything, mock.Anything)
	mq.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestExam_SubmitExam_InvalidOptionIDs_NoErrorZeroScore(t *testing.T) {
	mq := new(MockQuestionStore)
	mr := new(MockExamResultStore)
	mq.On("GetQuestions", mock.Anything).Return(sampleQuestions(), nil)
	mr.On("CandidateNameExists", mock.Anything, "BadIds").Return(false, nil)
	mr.On("SaveExamResult", mock.Anything, mock.MatchedBy(func(r *models.ExamResult) bool {
		return r.CandidateName == "BadIds" && r.Score == 0 && r.Total == 3
	})).Return(nil)

	ex := usecase.NewExam(mq, mr)
	res, err := ex.SubmitExam(context.Background(), "BadIds", map[string]string{
		"1": "garbage-uuid", "2": "not-real", "3": "x",
	})
	assert.NoError(t, err)
	assert.Equal(t, 0, res.Score)
	assert.Equal(t, 3, res.Total)

	mq.AssertExpectations(t)
	mr.AssertExpectations(t)
}

func TestExam_GetLeaderboard(t *testing.T) {
	mr := new(MockExamResultStore)
	t0 := time.Date(2026, 3, 24, 14, 30, 0, 0, time.UTC)
	t1 := time.Date(2026, 3, 24, 14, 15, 0, 0, time.UTC)
	mr.On("GetLeaderboard", mock.Anything, 20).Return([]models.ExamResult{
		{CandidateName: "Sophia", Score: 5, Total: 5, CreatedAt: t0},
		{CandidateName: "Alex", Score: 4, Total: 5, CreatedAt: t1},
	}, nil)

	ex := usecase.NewExam(new(MockQuestionStore), mr)
	entries, err := ex.GetLeaderboard(context.Background(), 0)
	assert.NoError(t, err)
	assert.Len(t, entries, 2)
	assert.Equal(t, 1, entries[0].Rank)
	assert.Equal(t, "Sophia", entries[0].CandidateName)
	assert.Equal(t, 5, entries[0].Score)
	assert.Equal(t, t0.UTC().Format(time.RFC3339), entries[0].CreatedAt)

	mr.AssertExpectations(t)
}
