package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"digital-scholar-exam/backend/internal/models"
	"digital-scholar-exam/backend/internal/repository"
	"digital-scholar-exam/backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

// stubQuestionStore — GetLeaderboard ไม่ใช้คำถาม แค่ต้องมีตัวแทน QuestionStore
type stubQuestionStore struct{}

func (stubQuestionStore) GetQuestions(ctx context.Context) ([]models.Question, error) {
	return nil, nil
}

func TestExamHTTP_GetLeaderboard_yourEntryFirstInJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := repository.OpenSQLite(":memory:")
	require.NoError(t, err)
	require.NoError(t, repository.AutoMigrate(db))

	t0 := time.Date(2026, 3, 24, 12, 0, 0, 0, time.UTC)
	require.NoError(t, db.Create(&models.ExamResult{CandidateName: "Alice", Score: 7, Total: 7, CreatedAt: t0}).Error)

	ex := usecase.NewExam(&stubQuestionStore{}, repository.NewExamResultGorm(db))
	h := NewExamHTTP(ex)
	router := gin.New()
	router.GET("/api/leaderboard", h.GetLeaderboard)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/leaderboard?forCandidate=Alice", nil)
	require.NoError(t, err)
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	body := w.Body.String()
	require.True(t, strings.HasPrefix(body, `{"yourEntry":`), "yourEntry ต้องอยู่ต้น JSON เพื่อดูด้วย curl | head")
	require.Contains(t, body, `"entries"`)

	var payload struct {
		YourEntry map[string]interface{} `json:"yourEntry"`
		Entries   []interface{}          `json:"entries"`
	}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &payload))
	require.NotNil(t, payload.YourEntry)
	require.Len(t, payload.Entries, 1)
}

func TestExamHTTP_GetLeaderboard_yourEntryNullWithoutForCandidate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := repository.OpenSQLite(":memory:")
	require.NoError(t, err)
	require.NoError(t, repository.AutoMigrate(db))

	t0 := time.Date(2026, 3, 24, 12, 0, 0, 0, time.UTC)
	require.NoError(t, db.Create(&models.ExamResult{CandidateName: "Bob", Score: 3, Total: 7, CreatedAt: t0}).Error)

	ex := usecase.NewExam(&stubQuestionStore{}, repository.NewExamResultGorm(db))
	h := NewExamHTTP(ex)
	router := gin.New()
	router.GET("/api/leaderboard", h.GetLeaderboard)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/leaderboard", nil)
	require.NoError(t, err)
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	require.True(t, strings.HasPrefix(w.Body.String(), `{"yourEntry":null`))

	var payload struct {
		YourEntry *usecase.LeaderboardYourEntryDTO `json:"yourEntry"`
		Entries   []interface{}                    `json:"entries"`
	}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &payload))
	require.Nil(t, payload.YourEntry)
	require.Len(t, payload.Entries, 1)
}

func TestExamHTTP_GetLeaderboard_forCandidateURLDecoded(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := repository.OpenSQLite(":memory:")
	require.NoError(t, err)
	require.NoError(t, repository.AutoMigrate(db))

	t0 := time.Date(2026, 3, 24, 12, 0, 0, 0, time.UTC)
	name := "zero มือ1"
	require.NoError(t, db.Create(&models.ExamResult{CandidateName: name, Score: 1, Total: 7, CreatedAt: t0}).Error)

	ex := usecase.NewExam(&stubQuestionStore{}, repository.NewExamResultGorm(db))
	h := NewExamHTTP(ex)
	router := gin.New()
	router.GET("/api/leaderboard", h.GetLeaderboard)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(
		http.MethodGet,
		"/api/leaderboard?forCandidate=zero%20%E0%B8%A1%E0%B8%B7%E0%B8%AD1",
		nil,
	)
	require.NoError(t, err)
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var payload struct {
		YourEntry struct {
			CandidateName string `json:"candidateName"`
			Rank          int    `json:"rank"`
		} `json:"yourEntry"`
	}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &payload))
	require.Equal(t, name, payload.YourEntry.CandidateName)
	require.Equal(t, 1, payload.YourEntry.Rank)
}
