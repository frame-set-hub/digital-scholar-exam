package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"digital-scholar-exam/backend/internal/handler"
	"digital-scholar-exam/backend/internal/repository"
	"digital-scholar-exam/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	dataDir := filepath.Join(".", "data")
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return err
	}
	dsn := filepath.Join(dataDir, "exam.db")

	db, err := repository.OpenSQLite(dsn)
	if err != nil {
		return err
	}
	if err := repository.AutoMigrate(db); err != nil {
		return err
	}
	if err := repository.SeedQuestionsIfEmpty(db); err != nil {
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
