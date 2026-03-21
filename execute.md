# Digital Scholar Exam — Execution Progress

Use this file to track project progress (supports opening a new session) — **update whenever a sub-task completes**.

## Checklist

- [x] Phase 1: Frontend Structure (Vue 3, Pinia, Tailwind, Docs init)
- [x] Phase 2: UI Integration (ExamView, ResultView, Single-choice logic)
- [x] Phase 3: Backend Initialization (Golang, Gin, SQLite, Pragmatic Clean Architecture)
- [x] Phase 4: API & Database Implementation (Mock Questions, Submit Exam)
- [x] Phase 5: Unit Testing Setup (testify/mock for use case — score calculation)
- [x] Phase 6: FE & BE Integration

## Notes

| Phase | Latest detail |
|-------|----------------|
| 3 | `cmd/api/main.go`, `internal/{models,repository,usecase,handler}`, GORM + SQLite, DI |
| 4 | Repository: `GetQuestions`, `SaveExamResult` — Use case: `GetQuestions`, `SubmitExam` + `ScoreAnswers` — Handler: `GET/POST` wired to use case |
| 5 | `exam_usecase_test.go`: mock repository, tests for full score / zero / partial + `SubmitExam` |
| 6 | Frontend: `GET /api/questions` + `POST /api/submit` (Vite proxy `/api` → :8080) — no bundled mock questions |

## Run backend (dev)

```bash
cd backend
go run ./cmd/api
```

- API: `http://localhost:8080` — `GET /api/questions`, `POST /api/submit`
- SQLite: `backend/data/exam.db` (created automatically)

## Documentation

- Index: [`docs/README.md`](./docs/README.md) · API: [`docs/api.md`](./docs/api.md)
- Short run commands: [`frontend/README.md`](./frontend/README.md), [`backend/README.md`](./backend/README.md)
