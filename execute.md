# Digital Scholar Exam — Execution Progress

Use this file to track project progress (supports opening a new session) — **update whenever a sub-task completes**.

## Checklist

- [x] Phase 1: Frontend Structure (Vue 3, Pinia, Tailwind, Docs init)
- [x] Phase 2: UI Integration (ExamView, ResultView, Single-choice logic)
- [x] Phase 3: Backend Initialization (Golang, Gin, SQLite, Pragmatic Clean Architecture)
- [x] Phase 4: API & Database Implementation (Mock Questions, Submit Exam)
- [x] Phase 5: Unit Testing Setup (testify/mock for use case — score calculation)
- [x] Phase 6: FE & BE Integration
- [x] Phase 7: Leaderboard UI & API (Fetch top scores, Sorting, Routing)

## Notes

| Phase | Latest detail |
|-------|----------------|
| 3 | `cmd/api/main.go`, `internal/{models,repository,usecase,handler}`, GORM + SQLite, DI |
<<<<<<< HEAD
| 4 | Repository: `GetQuestions`, `SaveExamResult` — Usecase: `GetQuestions`, `SubmitExam` + `ScoreAnswers` — Handler: `GET/POST` ผูก usecase |
| 5 | `exam_usecase_test.go`: mock repository, เทสคะแนนเต็ม / ศูนย์ / บางส่วน + `SubmitExam` |
| 6 | Frontend: `GET /api/questions` + `POST /api/submit` (Vite proxy `/api` → :8080) — ไม่มี mock ข้อสอบใน bundle |
| 7 | `GET /api/leaderboard` — เรียง `exam_results` ตามคะแนน (มาก→น้อย) แล้ว `created_at` (เก่าก่อน); FE route `/leaderboard` + Pinia `loadLeaderboard()` |
=======
| 4 | Repository: `GetQuestions`, `SaveExamResult` — Use case: `GetQuestions`, `SubmitExam` + `ScoreAnswers` — Handler: `GET/POST` wired to use case |
| 5 | `exam_usecase_test.go`: mock repository, tests for full score / zero / partial + `SubmitExam` |
| 6 | Frontend: `GET /api/questions` + `POST /api/submit` (Vite proxy `/api` → :8080) — no bundled mock questions |
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

## Run backend (dev)

```bash
cd backend
go run ./cmd/api
```

<<<<<<< HEAD
- API: `http://localhost:8080` — `GET /api/questions`, `POST /api/submit`, `GET /api/leaderboard`
- SQLite: `backend/data/exam.db` (สร้างอัตโนมัติ)
=======
- API: `http://localhost:8080` — `GET /api/questions`, `POST /api/submit`
- SQLite: `backend/data/exam.db` (created automatically)
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

## Documentation

- Index: [`docs/README.md`](./docs/README.md) · API: [`docs/api.md`](./docs/api.md)
- Short run commands: [`frontend/README.md`](./frontend/README.md), [`backend/README.md`](./backend/README.md)
