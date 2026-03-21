# Code map — file layout and line ranges

This document follows **startup order**: where imports come from, what each file does, and **line ranges** to open alongside your editor.

High-level **flows, use cases, and data** are in [architech.md](./architech.md)  
API contracts are in [api.md](./api.md)

---

## Frontend (run `npm run dev` in `frontend/`)

### 1. Bundle entry — `frontend/src/main.js`

| Lines (approx.) | What it does |
|---------------------|--------|
| 1 | Import `createApp` from Vue |
| 2 | Import `createPinia` |
| 3 | Import root `App.vue` |
| 4 | Import `router` from `./router` |
| 5 | Import app-wide CSS `./assets/main.css` |
| 7–11 | Create app → `use(Pinia)` → `use(router)` → `mount('#app')` |

Order matters: **Pinia first**, then router so every view can use the store

### 2. Root layout — `frontend/src/App.vue`

| Part | What it does |
|------|--------|
| `<template>` | Wraps `RouterView` — no global menu; each route is full-page |

### 3. Routes — `frontend/src/router/index.js`

| Lines (approx.) | What it does |
|---------------------|--------|
| 1 | `createRouter`, `createWebHistory` |
| 3–22 | Define `routes`: `/` → lazy `ExamView`, `/result` → `ResultView`, catch-all → `/` |
| 25–28 | `afterEach` sets `document.title` from `meta.title` |
| 30 | `export default router` — imported in `main.js` and `examStore.js` |

### 4. Browser HTTP — `frontend/src/api/client.js`

| Lines (approx.) | What it does |
|---------------------|--------|
| 6–9 | `apiBase()` reads `import.meta.env.VITE_API_BASE_URL` |
| 11–15 | `apiUrl(path)` joins base or uses relative `/api/...` |
| 17–38 | `fetchJSON` — sets `Content-Type` when body exists, `JSON.parse`, throws on `!res.ok` |

Together with **proxy** in `frontend/vite.config.js`: `server.proxy['/api']` forwards to `API_PROXY_TARGET` (default `http://localhost:8080`)

### 5. Shared state — `frontend/src/stores/examStore.js`

| Lines (approx.) | What it does |
|---------------------|--------|
| 1–4 | Import Pinia, Vue `ref`/`computed`, `router`, `apiUrl`/`fetchJSON` |
| 6–101 | `defineStore('exam', () => { ... })` — no mock questions in bundle; API only |
| 7–10 | state: `candidateName`, `questions`, `answers`, `score` |
| 12–13 | `loadState`, `loadError` |
| 15–16 | `loadQuestionsInflight` prevents duplicate concurrent GETs |
| 18 | `totalQuestions` = `questions.length` |
| 20–22 | `setAnswer` |
| 28–59 | `loadQuestions` — `GET /api/questions` only; on failure clears `questions` + sets `loadError` |
| 61–67 | `answersForSubmit` — string keys per API |
| 69–79 | `submitExam` — `POST /api/submit` then `router.push` to `result` |
| 81–86 | `resetExam` — clear name/answers/score, back to exam (keeps `questions`) |
| 88–100 | `return` — what components consume |

### 6. Exam view — `frontend/src/views/ExamView.vue`

| Part | Lines (approx.) | What it does |
|------|---------------------|--------|
| `<script setup>` | 1–4 | Import Vue, Pinia `storeToRefs`, `useExamStore` |
| | 6–7 | Refs from store |
| | 11–13 | `onMounted` → `exam.loadQuestions()` |
| | 15–18 | `allAnswered` — at least one question and every question answered |
| | 19–25 | `selectOption` / `isSelected` |
| | 27–46 | `handleSubmit` — validate name + all answered → `submitExam` + catch errors |
| | 48–76 | Tailwind class helpers for option cards |
| `<template>` | ~79+ | Main layout, `loadError` banner, heading, name field |
| | ~128–137 | Spinner when `loadState === 'loading'` |
| | ~139–172 | `v-for` questions and option buttons |
| | ~174–191 | Submit button |

### 7. Result view — `frontend/src/views/ResultView.vue`

| Part | Lines (approx.) | What it does |
|------|---------------------|--------|
| `<script setup>` | 1–5 | Import Vue, `useRouter`, Pinia, store |
| | 11–15 | If no `score` → `replace` to `exam` |
| | 17–25 | Score progress circle |
| | 27–29 | `retake` → `resetExam()` |
| `<template>` | ~32+ | Name, `score / totalQuestions`, Retake button |

---

## Backend (run `go run ./cmd/api` from `backend/` or per project)

### 8. Entry — `backend/cmd/api/main.go`

| Lines (approx.) | What it does |
|---------------------|--------|
| 8–14 | import: `handler`, `repository`, `usecase`, `gin` |
| 16–20 | `main()` calls `run()` |
| 22–27 | Create `data/` folder, path `data/exam.db` |
| 29–35 | `OpenSQLite` → `AutoMigrate` |
| 36–37 | `EnsureSeedQuestions` — seed if missing |
| 40–44 | DI: `NewQuestionGorm`, `NewExamResultGorm` → `usecase.NewExam` → `handler.NewExamHTTP` |
| 46–48 | `gin.Default()`, `corsMiddleware`, `RegisterRoutes` |
| 50–54 | Port `:8080` or `PORT` |

### 9. Route registration — `backend/internal/handler/router.go`

| Lines | What it does |
|--------|--------|
| 8–13 | Group `/api`: `GET /questions`, `POST /submit` |

### 10. HTTP + JSON — `backend/internal/handler/exam_handler.go`

| Lines (approx.) | What it does |
|---------------------|--------|
| 11–19 | struct `ExamHTTP`, constructor |
| 21–29 | `GetQuestions` → use case → `{ "questions": ... }` |
| 31–35 | `SubmitBody` — `candidateName`, `answers` |
| 37–55 | `Submit` — bind JSON, require non-empty `answers` → use case → 200 |

### 11. Business rules — `backend/internal/usecase/exam_usecase.go`

| Lines (approx.) | What it does |
|---------------------|--------|
| 11–20 | struct `Exam` references `QuestionStore`, `ExamResultStore` (interfaces from `ports.go`) |
| 22–42 | DTOs for API + `SubmitResponse` |
| 44–64 | `GetQuestions` — map `Question` → DTO **without answers** |
| 66–95 | `SubmitExam` — load questions → `ScoreAnswers` → build `ExamResult` + `SaveExamResult` |
| 97–107 | `ScoreAnswers` — compare `answers["id"]` to `CorrectOptionID` |

### 12. Use case ports — `backend/internal/usecase/ports.go`

| Lines | What it does |
|--------|--------|
| 9–17 | `QuestionStore`, `ExamResultStore` — repository implements |

### 13. DB models — `backend/internal/models/question.go`, `exam_result.go`

- `question.go`: `Question`, `Option`, `CorrectOptionID` in DB
- `exam_result.go`: name, score, total, `AnswersJSON`, `CreatedAt`

### 14. GORM / SQLite — `backend/internal/repository/`

| File | What it does |
|------|--------|
| `gorm.go` | `OpenSQLite` |
| `migrate.go` | `AutoMigrate` Question, Option, ExamResult |
| `seed.go` | `EnsureSeedQuestions`, `seedQuestions()` — sample questions for API |
| `question_gorm.go` | `GetQuestions` — `Preload("Options")`, `Order("sort_order")` |
| `exam_result_gorm.go` | `SaveExamResult` — `Create` |

### 15. Use case tests — `backend/internal/usecase/exam_usecase_test.go`

| What it does |
|--------|
| Mocks `QuestionStore` / `ExamResultStore`, tests `ScoreAnswers` and `SubmitExam` |

---

## Suggested reading order

**Frontend:** `main.js` → `App.vue` → `router/index.js` → `api/client.js` + `vite.config.js` → `stores/examStore.js` → `views/ExamView.vue` → `views/ResultView.vue`

**Backend:** `cmd/api/main.go` → `handler/router.go` → `handler/exam_handler.go` → `usecase/exam_usecase.go` + `ports.go` → `models/*` → `repository/*`

Then read the high-level flow in [architech.md](./architech.md)
