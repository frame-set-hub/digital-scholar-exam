# Digital Scholar Exam — Execution Progress

ใช้ไฟล์นี้ติดตามความคืบหน้าโปรเจกต์ (รองรับการเปิด New Session) — **อัปเดตทุกครั้งที่ Task ย่อยสำเร็จ**

## Checklist

- [x] Phase 1: Frontend Structure (Vue 3, Pinia, Tailwind, Docs init)
- [x] Phase 2: UI Integration (ExamView, ResultView, Single-choice logic)
- [x] Phase 3: Backend Initialization (Golang, Gin, SQLite, Pragmatic Clean Architecture)
- [x] Phase 4: API & Database Implementation (Mock Questions, Submit Exam)
- [x] Phase 5: Unit Testing Setup (testify/mock สำหรับ Usecase — คำนวณคะแนน)
- [x] Phase 6: FE & BE Integration
- [x] Phase 7: Leaderboard UI & API (Fetch top scores, Sorting, Routing)
- [x] Phase 8: ExamView — Submit validation UX (ไฮไลต์ข้อที่ยังไม่ตอบทุกข้อ + เลื่อนไปข้อแรกที่ว่าง)

## Notes

| Phase | รายละเอียดล่าสุด |
|-------|-------------------|
| 3 | `cmd/api/main.go`, `internal/{models,repository,usecase,handler}`, GORM + SQLite, DI |
| 4 | Repository: `GetQuestions`, `SaveExamResult` — Usecase: `GetQuestions`, `SubmitExam` + `ScoreAnswers` — Handler: `GET/POST` ผูก usecase |
| 5 | `exam_usecase_test.go`: mock repository, เทสคะแนนเต็ม / ศูนย์ / บางส่วน + `SubmitExam` |
| 6 | Frontend: `GET /api/questions` + `POST /api/submit` (Vite proxy `/api` → :8080) — ไม่มี mock ข้อสอบใน bundle |
| 7 | `GET /api/leaderboard` — เรียง `exam_results` ตามคะแนน (มาก→น้อย) แล้ว `created_at` (เก่าก่อน); FE route `/leaderboard` + Pinia `loadLeaderboard()` |
| 8 | `ExamView.vue`: กด Submit แต่ยังตอบไม่ครบ — `showUnansweredHighlight` + `questionSectionClasses` / ข้อความในการ์ด; `sectionRefs` + `scrollIntoView({ behavior: 'smooth', block: 'center' })` ไปข้อแรกที่ `answers[q.id] == null`; `formError` ใต้ปุ่มมี `animate-pulse` |

## รัน Backend (dev)

```bash
cd backend
go run ./cmd/api
```

- API: `http://localhost:8080` — `GET /api/questions`, `POST /api/submit`, `GET /api/leaderboard`
- SQLite: `backend/data/exam.db` (สร้างอัตโนมัติ)

## เอกสาร

- ดัชนี: [`docs/README.md`](./docs/README.md) · API: [`docs/api.md`](./docs/api.md)
- คำสั่งรันแพ็กเกจสั้นๆ: [`frontend/README.md`](./frontend/README.md), [`backend/README.md`](./backend/README.md)
