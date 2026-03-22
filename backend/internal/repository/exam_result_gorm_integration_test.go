//go:build integration

package repository

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"
)

// ตรวจว่า CandidateRank เจอชื่อ UTF-8 ตรงกับ SQLite — ถ้า fail แสดงว่า DB หรือสตริงไม่ตรง
func TestCandidateRank_zeroThaiName(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller")
	}
	// .../backend/internal/repository -> backend/data/exam.db
	backendDir := filepath.Join(filepath.Dir(thisFile), "..", "..")
	dbPath := filepath.Join(backendDir, "data", "exam.db")

	db, err := OpenSQLite(dbPath)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	r := NewExamResultGorm(db)

	name := "zero มือ1"
	rank, row, found, err := r.CandidateRank(context.Background(), name)
	if err != nil {
		t.Fatalf("CandidateRank: %v", err)
	}
	if !found {
		t.Fatalf("expected found for %q", name)
	}
	if rank < 1 {
		t.Fatalf("rank=%d", rank)
	}
	if row.CandidateName != name {
		t.Fatalf("row name %q != %q", row.CandidateName, name)
	}
}
