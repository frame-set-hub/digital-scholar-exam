# Digital Scholar Exam — Execution Progress

ใช้ไฟล์นี้ติดตามความคืบหน้าโปรเจกต์ (รองรับการเปิด New Session) — **อัปเดตทุกครั้งที่ Task ย่อยสำเร็จ**

## Checklist

- [x] Phase 1: Frontend Structure (Vue 3, Pinia, Tailwind, Docs init)
- [x] Phase 2: UI Integration (ExamView, ResultView, Single-choice logic)
- [x] Phase 3: Backend Initialization (Golang, Gin, SQLite, Pragmatic Clean Architecture)
- [x] Phase 4: API & Database Implementation (Mock Questions, Submit Exam)
- [x] Phase 5: Unit Testing Setup (testify/mock สำหรับ Usecase — คำนวณคะแนน)
- [x] Phase 6: FE & BE Integration

## Notes

| Phase | รายละเอียดล่าสุด |
|-------|-------------------|
| 3 | `cmd/api/main.go`, `internal/{models,repository,usecase,handler}`, GORM + SQLite, DI |
| 4 | Repository: `GetQuestions`, `SaveExamResult` — Usecase: `GetQuestions`, `SubmitExam` + `ScoreAnswers` — Handler: `GET/POST` ผูก usecase |
| 5 | `exam_usecase_test.go`: mock repository, เทสคะแนนเต็ม / ศูนย์ / บางส่วน + `SubmitExam` |
| 6 | Frontend: `GET /api/questions` + `POST /api/submit` (Vite proxy `/api` → :8080) — ไม่มี mock ข้อสอบใน bundle |

## รัน Backend (dev)

```bash
cd backend
go run ./cmd/api
```

- API: `http://localhost:8080` — `GET /api/questions`, `POST /api/submit`
- SQLite: `backend/data/exam.db` (สร้างอัตโนมัติ)

## เอกสาร

- ดัชนี: [`docs/README.md`](./docs/README.md) · API: [`docs/api.md`](./docs/api.md)
- คำสั่งรันแพ็กเกจสั้นๆ: [`frontend/README.md`](./frontend/README.md), [`backend/README.md`](./backend/README.md)
