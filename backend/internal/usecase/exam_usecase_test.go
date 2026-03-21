package usecase_test

import (
	"context"
	"testing"

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

func (m *MockExamResultStore) SaveExamResult(ctx context.Context, r *models.ExamResult) error {
	args := m.Called(ctx, r)
	return args.Error(0)
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
