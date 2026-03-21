# Testing Report (Frontend + Backend)

## Table of contents

- [Testing Report (Frontend + Backend)](#testing-report-frontend--backend)
  - [Table of contents](#table-of-contents)
  - [Current status](#current-status)
  - [Backend — Go / Use case](#backend--go--use-case)
  - [Backend — Integration / API (recommended next)](#backend--integration--api-recommended-next)
  - [Frontend — Unit / Component](#frontend--unit--component)
  - [CI guidance](#ci-guidance)

## Current status

| Date | Run / item | Frontend | Backend | Notes |
|--------|----------------|----------|---------|----------|
| — | Last manual run | Vitest pending | `go test ./...` (see [`../backend/README.md`](../backend/README.md)) | Update when CI runs |

## Backend — Go / Use case

| Item | Command / location | Status |
|--------|------------------|--------|
| Use case unit tests | `cd backend && go test ./... -count=1` | Yes — `internal/usecase/exam_usecase_test.go` |
| What is tested | `ScoreAnswers` + `SubmitExam` with **mock** `QuestionStore` / `ExamResultStore` (testify/mock) — includes **edge cases** below | Yes |
| Tools | `testing`, `github.com/stretchr/testify/assert`, `github.com/stretchr/testify/mock` | |

**Edge cases (`internal/usecase/exam_usecase_test.go`):**

| Case | Assertion |
|------|----------------|
| All correct | `ScoreAnswers` + `SubmitExam` — `score == total` (full score) |
| All wrong | `ScoreAnswers` + `SubmitExam` — `score == 0` |
| Incomplete answers | 2 questions, submit 1 — scoring works, no error; unanswered rows get no points (`ScoreAnswers` / `SubmitExam` partial) |
| Invalid / non-existent option ID | No panic / no error; that question scores zero (`ScoreAnswers` + `SubmitExam` zero) |
| Answer keys for unknown questions | Odd keys ignored when scoring (`WrongKeysIgnoredForUnknownQuestion`) |

**Example run:**

```bash
cd backend
go test ./... -count=1 -v
```

## Backend — Integration / API (recommended next)

| Suite | Tool | Status |
|----------|------------|--------|
| Handler + DB in-memory or temp SQLite | `httptest`, GORM `:memory:` | Pending (optional) |
| E2E against real server | curl / Playwright API | When pipeline / E2E stack exists |

## Frontend — Unit / Component

| Suite | Tool | Status |
|----------|------------|--------|
| Pinia `examStore` (score logic, reset) | Vitest | Pending |
| `ExamView` / `ResultView` | Vitest + Vue Test Utils | Pending |

When added, record commands (e.g. `cd frontend && npm run test`) and pass/fail in the table above.

## CI guidance

- Job 1: `cd backend && go test ./...`
- Job 2: `cd frontend && npm ci && npm run build` (and `npm run test` when present)
- Split coverage reports under `backend/` / `frontend/`
