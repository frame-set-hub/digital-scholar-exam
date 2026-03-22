package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"digital-scholar-exam/backend/internal/handler"
	"digital-scholar-exam/backend/internal/repository"
	"digital-scholar-exam/backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	_ = godotenv.Load()

	dataDir, err := resolveDataDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return fmt.Errorf("create data dir %q: %w", dataDir, err)
	}
	dsn := filepath.Join(dataDir, "exam.db")

	db, err := repository.OpenSQLite(dsn)
	if err != nil {
		return err
	}
	if err := repository.AutoMigrate(db); err != nil {
		return err
	}
	if err := repository.EnsureSeedQuestions(db); err != nil {
		return err
	}

	// Dependency injection
	qRepo := repository.NewQuestionGorm(db)
	rRepo := repository.NewExamResultGorm(db)
	examUC := usecase.NewExam(qRepo, rRepo)
	examH := handler.NewExamHTTP(examUC)

	r := gin.Default()
	r.Use(corsMiddleware())
	handler.RegisterRoutes(r, examH)

	addr := ":8080"
	if v := os.Getenv("PORT"); v != "" {
		addr = ":" + v
	}
	return r.Run(addr)
}

// resolveDataDir picks where SQLite lives: $DATABASE_DIR if set, else <process working dir>/data.
// Using an absolute path avoids silent failures when cwd is wrong or stale (e.g. shell still in a deleted folder).
func resolveDataDir() (string, error) {
	if v := os.Getenv("DATABASE_DIR"); v != "" {
		abs, err := filepath.Abs(v)
		if err != nil {
			return "", fmt.Errorf("DATABASE_DIR: %w", err)
		}
		return abs, nil
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory (cd into backend/ or set DATABASE_DIR): %w", err)
	}
	abs, err := filepath.Abs(filepath.Join(wd, "data"))
	if err != nil {
		return "", err
	}
	return abs, nil
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
