package repository

import (
	"context"
	"testing"
	"time"

	"digital-scholar-exam/backend/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupMemoryDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err)
	require.NoError(t, AutoMigrate(db))
	return db
}

func TestExamResultGorm_GetLeaderboard_orderAndLimit(t *testing.T) {
	db := setupMemoryDB(t)

	t1 := time.Date(2026, 3, 24, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2026, 3, 24, 11, 0, 0, 0, time.UTC)
	require.NoError(t, db.Create(&models.ExamResult{CandidateName: "Low", Score: 1, Total: 7, CreatedAt: t1}).Error)
	require.NoError(t, db.Create(&models.ExamResult{CandidateName: "High", Score: 9, Total: 7, CreatedAt: t2}).Error)

	r := NewExamResultGorm(db)
	rows, err := r.GetLeaderboard(context.Background(), 20)
	require.NoError(t, err)
	require.Len(t, rows, 2)
	assert.Equal(t, "High", rows[0].CandidateName)
	assert.Equal(t, 9, rows[0].Score)
	assert.Equal(t, "Low", rows[1].CandidateName)
}

func TestExamResultGorm_CandidateRank_singleRow(t *testing.T) {
	db := setupMemoryDB(t)

	t0 := time.Date(2026, 3, 24, 12, 0, 0, 0, time.UTC)
	require.NoError(t, db.Create(&models.ExamResult{CandidateName: "Solo", Score: 5, Total: 7, CreatedAt: t0}).Error)

	r := NewExamResultGorm(db)
	rank, row, found, err := r.CandidateRank(context.Background(), "Solo")
	require.NoError(t, err)
	require.True(t, found)
	assert.Equal(t, 1, rank)
	assert.Equal(t, "Solo", row.CandidateName)
}

func TestExamResultGorm_CandidateRank_tieBreak_createdAt(t *testing.T) {
	db := setupMemoryDB(t)

	tOld := time.Date(2026, 3, 24, 10, 0, 0, 0, time.UTC)
	tNew := time.Date(2026, 3, 24, 11, 0, 0, 0, time.UTC)
	require.NoError(t, db.Create(&models.ExamResult{CandidateName: "First", Score: 5, Total: 7, CreatedAt: tOld}).Error)
	require.NoError(t, db.Create(&models.ExamResult{CandidateName: "Second", Score: 5, Total: 7, CreatedAt: tNew}).Error)

	r := NewExamResultGorm(db)
	rank, _, found, err := r.CandidateRank(context.Background(), "Second")
	require.NoError(t, err)
	require.True(t, found)
	assert.Equal(t, 2, rank)
}

func TestExamResultGorm_CandidateRank_utf8Name(t *testing.T) {
	db := setupMemoryDB(t)

	t0 := time.Date(2026, 3, 24, 12, 0, 0, 0, time.UTC)
	name := "zero มือ1"
	require.NoError(t, db.Create(&models.ExamResult{CandidateName: name, Score: 1, Total: 7, CreatedAt: t0}).Error)

	r := NewExamResultGorm(db)
	rank, row, found, err := r.CandidateRank(context.Background(), name)
	require.NoError(t, err)
	require.True(t, found)
	assert.Equal(t, 1, rank)
	assert.Equal(t, name, row.CandidateName)
}

func TestExamResultGorm_CandidateRank_notFound(t *testing.T) {
	db := setupMemoryDB(t)

	r := NewExamResultGorm(db)
	_, _, found, err := r.CandidateRank(context.Background(), "Nope")
	require.NoError(t, err)
	assert.False(t, found)
}
