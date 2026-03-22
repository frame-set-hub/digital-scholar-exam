# Testing Report (Frontend + Backend)

## สารบัญ

- [Testing Report (Frontend + Backend)](#testing-report-frontend--backend)
  - [สารบัญ](#สารบัญ)
  - [สถานะปัจจุบัน](#สถานะปัจจุบัน)
  - [Backend — Go / Use case](#backend--go--use-case)
  - [Backend — Integration / API (แนะนำต่อไป)](#backend--integration--api-แนะนำต่อไป)
  - [Frontend — Unit / Component](#frontend--unit--component)
  - [แนวทาง CI](#แนวทาง-ci)

## สถานะปัจจุบัน

| วันที่ | รัน / รายการ | Frontend | Backend | หมายเหตุ |
|--------|----------------|----------|---------|----------|
| — | รันครั้งล่าสุด (กรอกมือ) | รอเพิ่ม Vitest | `go test ./...` (ดู [`../backend/README.md`](../backend/README.md)) | อัปเดตเมื่อ CI รัน |

## Backend — Go / Use case

| รายการ | คำสั่ง / ตำแหน่ง | สถานะ |
|--------|------------------|--------|
| Unit test use case | `cd backend && go test ./... -count=1` | มี — `internal/usecase/exam_usecase_test.go` |
| เนื้อหาที่ทดสอบ | `ScoreAnswers` + `SubmitExam` + `GetLeaderboard` (รวม `forCandidate` / `yourEntry` / `InTopList`) ด้วย **mock** `QuestionStore` / `ExamResultStore` (testify/mock) — รวม **edge cases** ด้านล่าง | มี |
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

**ตัวอย่างรัน:**

```bash
cd backend
go test ./... -count=1 -v
```

## Backend — Integration / API (แนะนำต่อไป)

| ชุดทดสอบ | เครื่องมือ | สถานะ |
|----------|------------|--------|
| Handler + DB in-memory หรือ temp SQLite | `httptest`, GORM `:memory:` | รอเพิ่ม (optional) |
| E2E ยิง `GET/POST` กับ server จริง | curl / Playwright API | เมื่อมี pipeline / สแตก E2E |

## Frontend — Unit / Component

| ชุดทดสอบ | เครื่องมือ | สถานะ |
|----------|------------|--------|
| Pinia `examStore` (คำนวณคะแนน, reset) | Vitest | รอเพิ่ม |
| `ExamView` / `ResultView` | Vitest + Vue Test Utils | รอเพิ่ม |

เมื่อเพิ่มแล้วให้บันทึกคำสั่ง (เช่น `cd frontend && npm run test`) และผล pass/fail ในตารางด้านบน

## แนวทาง CI

- Job 1: `cd backend && go test ./...`
- Job 2: `cd frontend && npm ci && npm run build` (และ `npm run test` เมื่อมี)
- แยกรายงาน coverage ตามโฟลเดอร์ `backend/` / `frontend/`
