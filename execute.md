# Digital Scholar Exam — Execution Progress

ใช้ไฟล์นี้ติดตามความคืบหน้าโปรเจกต์ (รองรับการเปิด New Session) — **อัปเดตทุกครั้งที่ Task ย่อยสำเร็จ**

## Checklist

- [x] Phase 1: Project Setup (Monorepo, Documentation)
- [x] Phase 2: UI Integration (ExamView, ResultView, Single-choice logic)
- [x] Phase 3: Backend Initialization (Golang, Gin, SQLite, Pragmatic Clean Architecture)
- [x] Phase 4: API & Database Implementation (Mock Questions, Submit Exam)
- [x] Phase 5: Unit Testing Setup (testify/mock สำหรับ use case — คำนวณคะแนน)
- [x] Phase 6: FE & BE Integration
- [x] Phase 7: Leaderboard UI & API (Fetch top scores, Sorting, Routing)
- [x] Phase 8: Submit validation & UX — ไฮไลต์ข้อที่ยังไม่ตอบ + เลื่อนไปช่องว่างแรก; API ชื่อซ้ำ (`409`) / error ชื่อบนฟิลด์ + เลื่อน; Enter ที่ชื่อ → โฟกัส Submit

## Backlog (วางแผน — ยังไม่เริ่ม)

- [ ] **Leaderboard “เห็นตัวเอง” เมื่ออันดับเกิน top N** — ตอนนี้ `GET /api/leaderboard` จำกัด 20 แถว (default/max); ถ้าคะแนนของเราอยู่นอก 20 อันดับแรก หน้า `/leaderboard` จะไม่แสดงแถวของตัวเอง — ต้องออกแบบ API/UX เพิ่ม (เช่น query ตาม `candidateName` + rank, หรือแถว “Your position” แยก) สอดคล้องกับ [`docs/planning.md`](./docs/planning.md)

## Notes

| Phase | รายละเอียดล่าสุด |
|-------|-------------------|
| 3 | `cmd/api/main.go`, `internal/{models,repository,usecase,handler}`, GORM + SQLite, DI |
| 4 | Repository: `GetQuestions`, `SaveExamResult` — Use case: `GetQuestions`, `SubmitExam` + `ScoreAnswers` — Handler: ผูก `GET/POST` กับ use case |
| 5 | `exam_usecase_test.go`: mock repository, เทสคะแนนเต็ม / ศูนย์ / บางส่วน + `SubmitExam` |
| 6 | Frontend: `GET /api/questions` + `POST /api/submit` (Vite proxy `/api` → :8080) — ไม่มี mock ข้อสอบใน bundle |
| 7 | `GET /api/leaderboard` — เรียง `exam_results` ตามคะแนน (มาก→น้อย) แล้ว `created_at` (เก่าก่อน); FE route `/leaderboard` + Pinia `loadLeaderboard()` |
| 8 | `ExamView.vue`: ยังตอบไม่ครบ — `showUnansweredHighlight` + กรอบแดงต่อข้อ + ข้อความในการ์ด; `sectionRefs` + เลื่อนแบบ smooth ไปช่องว่างแรก; `submitError` ใต้ปุ่ม (pulse เมื่อ validation). ชื่อ/API — `nameError`, ชื่อซ้ำ `409` / `400`; `fetchJSON` ใส่ `err.status`; Enter ที่ชื่อ → โฟกัส Submit |

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
