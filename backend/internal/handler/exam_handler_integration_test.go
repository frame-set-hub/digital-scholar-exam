//go:build integration

package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"digital-scholar-exam/backend/internal/repository"
	"digital-scholar-exam/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func TestGetLeaderboard_HTTP_forCandidate_returnsYourEntry(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller")
	}
	backendDir := filepath.Join(filepath.Dir(thisFile), "..", "..")
	dbPath := filepath.Join(backendDir, "data", "exam.db")

	db, err := repository.OpenSQLite(dbPath)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	rRepo := repository.NewExamResultGorm(db)
	qRepo := repository.NewQuestionGorm(db)
	examUC := usecase.NewExam(qRepo, rRepo)
	examH := NewExamHTTP(examUC)

	r := gin.New()
	r.GET("/api/leaderboard", examH.GetLeaderboard)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(
		http.MethodGet,
		"/api/leaderboard?forCandidate=zero%20%E0%B8%A1%E0%B8%B7%E0%B8%AD1",
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status %d body %s", w.Code, w.Body.String())
	}
	var payload struct {
		Entries   []interface{} `json:"entries"`
		YourEntry interface{}   `json:"yourEntry"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &payload); err != nil {
		t.Fatal(err)
	}
	if payload.YourEntry == nil {
		t.Fatalf("expected yourEntry in JSON, body=%s", w.Body.String())
	}
}
