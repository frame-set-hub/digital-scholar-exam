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
- [x] Phase 8: Submit validation & UX — highlight unanswered questions + scroll to first gap; duplicate-name API (`409`) / name errors on field + scroll; Enter on name focuses Submit

## Backlog (planned — not started)

- [ ] **Leaderboard “เห็นตัวเอง”เมื่ออันดับเกิน top N** — ตอนนี้ `GET /api/leaderboard` จำกัด 20 แถว (default/max); ถ้าคะแนนของเราอยู่นอก 20 อันดับแรก หน้า `/leaderboard` จะไม่แสดงแถวของตัวเอง — ต้องออกแบบ API/UX เพิ่ม (เช่น query ตาม `candidateName` + rank, หรือแถว “Your position” แยก) สอดคล้องกับ [`docs/planning.md`](./docs/planning.md)

## Notes

| Phase | Latest detail |
|-------|----------------|
| 3 | `cmd/api/main.go`, `internal/{models,repository,usecase,handler}`, GORM + SQLite, DI |
| 4 | Repository: `GetQuestions`, `SaveExamResult` — Use case: `GetQuestions`, `SubmitExam` + `ScoreAnswers` — Handler: `GET/POST` wired to use case |
| 5 | `exam_usecase_test.go`: mock repository, tests for full score / zero / partial + `SubmitExam` |
| 6 | Frontend: `GET /api/questions` + `POST /api/submit` (Vite proxy `/api` → :8080) — no bundled mock questions |
| 7 | `GET /api/leaderboard` — sort `exam_results` by score (high→low) then `created_at` (earliest first); FE route `/leaderboard` + Pinia `loadLeaderboard()` |
| 8 | `ExamView.vue`: incomplete — `showUnansweredHighlight` + per-question red border + in-card copy; `sectionRefs` + smooth scroll to first gap; `submitError` under button (pulse when validation). Name/API — `nameError`, duplicate `409` / `400`; `fetchJSON` exposes `err.status`; Enter on name → focus Submit |

## Run backend (dev)

```bash
cd backend
go run ./cmd/api
```

- API: `http://localhost:8080` — `GET /api/questions`, `POST /api/submit`, `GET /api/leaderboard`
- SQLite: `backend/data/exam.db` (created automatically)

## Documentation

- Index: [`docs/README.md`](./docs/README.md) · API: [`docs/api.md`](./docs/api.md)
- Short run commands: [`frontend/README.md`](./frontend/README.md), [`backend/README.md`](./backend/README.md)
