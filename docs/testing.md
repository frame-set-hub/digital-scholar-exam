# Testing Report (Frontend + Backend)

## สารบัญ

- [Testing Report (Frontend + Backend)](#testing-report-frontend--backend)
  - [สารบัญ](#สารบัญ)
  - [สถานะปัจจุบัน](#สถานะปัจจุบัน)
  - [Backend — Go / Use case](#backend--go--use-case)
  - [Backend — Repository (SQLite in-memory)](#backend--repository-sqlite-in-memory)
  - [Backend — Handler (HTTP + JSON)](#backend--handler-http--json)
  - [Backend — Integration (`exam.db` จริง)](#backend--integration-examdb-จริง)
  - [Backend — Integration / API (E2E แนะนำต่อไป)](#backend--integration--api-e2e-แนะนำต่อไป)
  - [Frontend — Unit / Component](#frontend--unit--component)
  - [แนวทาง CI](#แนวทาง-ci)

## สถานะปัจจุบัน

| วันที่ | รัน / รายการ | Frontend | Backend | หมายเหตุ |
|--------|----------------|----------|---------|----------|
| 2026-03-22 | รันครั้งล่าสุด (กรอกมือ) | รอเพิ่ม Vitest | `go test ./...` + `go test -tags=integration …` (ดู [`../backend/README.md`](../backend/README.md)) | Leaderboard: usecase + repository + handler + integration tag |

## Backend — Go / Use case

| รายการ | คำสั่ง / ตำแหน่ง | สถานะ |
|--------|------------------|--------|
| Unit test use case | `cd backend && go test ./... -count=1` | มี — `internal/usecase/exam_usecase_test.go` |
| เนื้อหาที่ทดสอบ | `ScoreAnswers` + `SubmitExam` + `GetLeaderboard` (รวม `forCandidate` / `yourEntry` / `InTopList`) ด้วย **mock** `QuestionStore` / `ExamResultStore` (testify/mock) — รวม **edge cases** ด้านล่าง | มี |
| เพิ่มเติม `GetLeaderboard` | `forCandidate` ไม่พบใน DB; ชื่อว่างหลัง trim; `CandidateRank` error จาก repo | มี |
| เครื่องมือ | `testing`, `github.com/stretchr/testify/assert`, `github.com/stretchr/testify/mock` | |

**Edge cases (`internal/usecase/exam_usecase_test.go`):**

| กรณี | สิ่งที่ยืนยัน |
|------|----------------|
| ตอบถูกหมด | `ScoreAnswers` + `SubmitExam` — `score == total` (คะแนนเต็ม) |
| ตอบผิดหมด | `ScoreAnswers` + `SubmitExam` — `score == 0` |
| ส่งคำตอบไม่ครบข้อ | ข้อสอบ 2 ข้อ ส่งมาแค่ 1 ข้อ — คำนวณได้ ไม่ error; ข้อที่ไม่ส่งไม่ได้คะแนน (`ScoreAnswers` / `SubmitExam` partial) |
| option ID ไม่ถูกต้อง / ไม่มีจริง | ไม่ panic / ไม่ error; ข้อนั้นไม่ได้คะแนน (`ScoreAnswers` + `SubmitExam` ศูนย์คะแนน) |
| คีย์คำตอบที่ไม่ตรงข้อสอบ | คีย์แปลกๆ ถูกละเว้นเมื่อนับคะแนน (`WrongKeysIgnoredForUnknownQuestion`) |
| ชื่อว่างหลัง trim | `SubmitExam("   ", …)` → `ErrCandidateNameRequired` |
| ชื่อซ้ำ | `CandidateNameExists` เป็น true → `ErrDuplicateCandidateName`; ไม่เรียก `SaveExamResult` |
| Leaderboard `forCandidate` ไม่พบ | `CandidateRank` → `found: false` → `yourEntry` เป็น nil |
| Leaderboard `forCandidate` เป็นช่องว่าง | trim แล้วว่าง → ไม่เรียก `CandidateRank` |
| Leaderboard `CandidateRank` error | ส่ง error ต่อจาก use case |

**ตัวอย่างรัน:**

```bash
cd backend
go test ./... -count=1 -v
```

## Backend — Repository (SQLite in-memory)

| รายการ | คำสั่ง / ตำแหน่ง | สถานะ |
|--------|------------------|--------|
| Unit test GORM `ExamResultGorm` | `cd backend && go test ./internal/repository/ -count=1` | มี — `internal/repository/exam_result_gorm_test.go` |
| เนื้อหาที่ทดสอบ | `GetLeaderboard` เรียงคะแนน; `CandidateRank` แถวเดียว / เสมอกันตาม `created_at` / ชื่อ UTF-8 (`zero มือ1`) / ไม่พบชื่อ | มี |
| หมายเหตุ | เปิด DB ด้วย `gorm` + `sqlite.Open(":memory:")` + `Logger: Silent` ในเทส — ไม่พึ่งไฟล์ `exam.db` | |

## Backend — Handler (HTTP + JSON)

| รายการ | คำสั่ง / ตำแหน่ง | สถานะ |
|--------|------------------|--------|
| Unit test `GET /api/leaderboard` | `cd backend && go test ./internal/handler/ -count=1` | มี — `internal/handler/exam_handler_leaderboard_test.go` |
| เนื้อหาที่ทดสอบ | Response ขึ้นต้นด้วย `yourEntry` ก่อน `entries`; `yourEntry` เป็น `null` เมื่อไม่มี `forCandidate`; query `forCandidate` แบบ URL-encoded ภาษาไทย | มี |
| เครื่องมือ | `httptest`, `gin`, in-memory SQLite + `stubQuestionStore` + `usecase.Exam` จริง | |

## Backend — Integration (`exam.db` จริง)

| รายการ | คำสั่ง | สถานะ |
|--------|--------|--------|
| Handler + DB จาก path คงที่ | `cd backend && go test -tags=integration ./internal/handler/ ./internal/repository/ -count=1` | มี — `exam_handler_integration_test.go`, `exam_result_gorm_integration_test.go` |
| เงื่อนไข | ต้องมีไฟล์ [`backend/data/exam.db`](../backend/data/exam.db) (หรือ DB ที่ path เดียวกับที่เทส resolve) | ไม่รันใน CI ถ้าไม่มีไฟล์ / ไม่ใส่ tag |

## Backend — Integration / API (E2E แนะนำต่อไป)

| ชุดทดสอบ | เครื่องมือ | สถานะ |
|----------|------------|--------|
| E2E ยิง `GET/POST` กับ server จริง | curl / Playwright API | เมื่อมี pipeline / สแตก E2E |

## Frontend — Unit / Component

| ชุดทดสอบ | เครื่องมือ | สถานะ |
|----------|------------|--------|
| Pinia `examStore` (คำนวณคะแนน, reset) | Vitest | รอเพิ่ม |
| `ExamView` / `ResultView` | Vitest + Vue Test Utils | รอเพิ่ม |

เมื่อเพิ่มแล้วให้บันทึกคำสั่ง (เช่น `cd frontend && npm run test`) และผล pass/fail ในตารางด้านบน

## แนวทาง CI

- Job 1: `cd backend && go test ./...` (ครอบ usecase + repository + handler แบบ in-memory — **ไม่**ต้องมี `exam.db`)
- Job 2 (optional): `cd backend && go test -tags=integration ./internal/handler/ ./internal/repository/` เมื่อต้องการยืนยันกับ `data/exam.db` ใน workspace / artifact
- Job 3: `cd frontend && npm ci && npm run build` (และ `npm run test` เมื่อมี)
- แยกรายงาน coverage ตามโฟลเดอร์ `backend/` / `frontend/`
