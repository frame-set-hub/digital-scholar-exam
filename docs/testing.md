# Testing Report (Frontend + Backend)

## สารบัญ

- [Testing Report (Frontend + Backend)](#testing-report-frontend--backend)
  - [สารบัญ](#สารบัญ)
  - [สถานะปัจจุบัน](#สถานะปัจจุบัน)
  - [Backend — Go / Usecase](#backend--go--usecase)
  - [Backend — Integration / API (แนะนำเติม)](#backend--integration--api-แนะนำเติม)
  - [Frontend — Unit / Component](#frontend--unit--component)
  - [แนวทาง CI](#แนวทาง-ci)

## สถานะปัจจุบัน

| วันที่ | รอบ / รายการ | Frontend | Backend | หมายเหตุ |
|--------|----------------|----------|---------|----------|
| — | รันครั้งล่าสุด (กรอกมือ) | รอเพิ่ม Vitest | `go test ./...` (ดู [`../backend/README.md`](../backend/README.md)) | อัปเดตเมื่อรัน CI |

## Backend — Go / Usecase

| รายการ | คำสั่ง / ตำแหน่ง | สถานะ |
|--------|------------------|--------|
| Unit test usecase | `cd backend && go test ./... -count=1` | มี — `internal/usecase/exam_usecase_test.go` |
| เนื้อหาที่ทดสอบ | `ScoreAnswers` (เต็ม / ศูนย์ / บางส่วน), `SubmitExam` ด้วย **mock** `QuestionStore` / `ExamResultStore` (testify/mock) | มี |
| เครื่องมือ | `testing`, `github.com/stretchr/testify/assert`, `github.com/stretchr/testify/mock` | |

**ตัวอย่างรัน:**

```bash
cd backend
go test ./... -count=1 -v
```

## Backend — Integration / API (แนะนำเติม)

| ชุดทดสอบ | เครื่องมือ | สถานะ |
|----------|------------|--------|
| Handler + DB in-memory หรือ temp SQLite | `httptest`, GORM `:memory:` | รอเพิ่ม (optional) |
| E2E ยิง `GET/POST` กับ server จริง | curl / Playwright API | รอ Phase 6 |

## Frontend — Unit / Component

| ชุดทดสอบ | เครื่องมือ | สถานะ |
|----------|------------|--------|
| Pinia `examStore` (คำนวณคะแนน, reset) | Vitest | รอเพิ่ม |
| `ExamView` / `ResultView` | Vitest + Vue Test Utils | รอเพิ่ม |

เมื่อเพิ่มแล้วให้บันทึกคำสั่ง (เช่น `cd frontend && npm run test`) และผล pass/fail ในตารางด้านบน

## แนวทาง CI

- Job 1: `cd backend && go test ./...`
- Job 2: `cd frontend && npm ci && npm run build` (และ `npm run test` เมื่อมี)
- แยก coverage รายงานตามโฟลเดอร์ `backend/` / `frontend/`
