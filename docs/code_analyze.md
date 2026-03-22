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
| 3–28 | Define `routes`: `/` → lazy `ExamView`, `/result` → `ResultView`, `/leaderboard` → `LeaderboardView`, catch-all → `/` |
| 31–34 | `afterEach` sets `document.title` from `meta.title` |
| 36 | `export default router` — imported in `main.js` and `examStore.js` |

### 4. Browser HTTP — `frontend/src/api/client.js`

| Lines (approx.) | What it does |
|---------------------|--------|
| 6–9 | `apiBase()` reads `import.meta.env.VITE_API_BASE_URL` |
| 11–15 | `apiUrl(path)` joins base or uses relative `/api/...` |
| 17–38 | `fetchJSON` — sets `Content-Type` when body exists, `JSON.parse`, throws on `!res.ok` |

Together with **proxy** in `frontend/vite.config.js`: `server.port` from `DEV_SERVER_PORT` in `.env` (default `5173`); `server.proxy['/api']` forwards to `API_PROXY_TARGET` (default `http://localhost:8080`)

### 5. Shared state — `frontend/src/stores/examStore.js`

| Lines (approx.) | What it does |
|---------------------|--------|
| 1–4 | Import Pinia, Vue `ref`/`computed`, `router`, `apiUrl`/`fetchJSON` |
| 6–129 | `defineStore('exam', () => { ... })` — no mock questions in bundle; API only |
| 7–10 | state: `candidateName`, `questions`, `answers`, `score` |
| 12–14 | `leaderboard`, `leaderboardState`, `leaderboardError` |
| 16–17 | `loadState`, `loadError` |
| 19–20 | `loadQuestionsInflight` prevents duplicate concurrent GETs |
| 22 | `totalQuestions` = `questions.length` |
| 24–26 | `setAnswer` |
| 32–63 | `loadQuestions` — `GET /api/questions` only; on failure clears `questions` + sets `loadError` |
| 65–80 | `loadLeaderboard` — `GET /api/leaderboard` → `entries` stored in `leaderboard` |
| 82–88 | `answersForSubmit` — string keys per API contract |
| 90–100 | `submitExam` — `POST /api/submit` then `router.push` to `result` |
| 102–110 | `resetExam` — clears name/answers/score/leaderboard, returns to exam (keeps `questions`) |
| 112–128 | `return` — what components consume |

### 6. Exam view — `frontend/src/views/ExamView.vue` (**~260 lines**)

Same file as in the repo — split by `<script>` / `<template>` / `<style>` blocks.

#### `<script setup>` — lines 1–112

| Lines | What it does |
|--------|--------|
| 1–4 | Import `ref`, `computed`, `onMounted`, `nextTick` from Vue · `storeToRefs` from Pinia · `useExamStore` |
| 6–7 | Create `exam` · `storeToRefs(exam)` → `candidateName`, `questions`, `answers`, `loadState`, `loadError` |
| 9–13 | `formError` · `showUnansweredHighlight` · `sectionRefs` map · `setSectionRef(questionId, el)` for `:ref` callbacks |
| 23–25 | `onMounted(() => exam.loadQuestions())` |
| 27–30 | `allAnswered` (computed) — false if no questions; else every `q.id` has a value in `answers` |
| 32–42 | `isSelected` / `selectOption` / `isQuestionUnanswered` |
| 44–52 | `questionSectionClasses` — default card chrome vs red border/background when highlight + unanswered |
| 54–81 | `handleSubmit` — clear errors/highlight · trim name (empty → message, return) · if not `allAnswered`: set message + `showUnansweredHighlight`, `nextTick` + `scrollIntoView({ behavior: 'smooth', block: 'center' })` on first unanswered section, return · else `await exam.submitExam()` · `catch` maps fetch/TypeError to user-facing text |
| 83–111 | `optionCardClasses` / `indicatorClasses` / `optionTextClasses` — Tailwind class arrays for choice UI |
| 112 | Close `</script>` |

#### `<template>` — lines 114–250

| Lines | What it does |
|--------|--------|
| 115–117 | Root `div` (`min-h-screen`, `bg-background`) · `<main>` `max-w-3xl`, padding |
| 118–124 | `v-if="loadError"` — amber alert, `role="status"` |
| 125–161 | Header + candidate name: badge, title, module line · `input#candidate-name` · focus underline |
| 163–172 | Loading spinner + copy |
| 175–217 | `v-else` question list: each `section` has `:id`, `:ref` → `setSectionRef`, `:class="questionSectionClasses(q.id)"` · options loop · `v-if` in-card alert when highlight + unanswered |
| 219–240 | Submit block: `formError` with `animate-pulse` · Submit button |
| 241–249 | Fixed blur decorations |

#### `<style scoped>` — lines 252–260

| Lines | What it does |
|--------|--------|
| 252–260 | `.material-symbols-outlined` — `font-variation-settings` |

### 7. Result view — `frontend/src/views/ResultView.vue`

| Part | Lines (approx.) | What it does |
|------|---------------------|--------|
| `<script setup>` | 1–5 | Import Vue, `useRouter`, Pinia, store |
| | 11–15 | If no `score` → `replace` to `exam` |
| | 17–25 | Score progress circle calculation |
| | 27–33 | `retake` → `resetExam()` · `goLeaderboard` → `router.push` to `leaderboard` |
| `<template>` | ~36+ | Name, `score / totalQuestions`, View Leaderboard + Retake Exam buttons |

### 8. Leaderboard view — `frontend/src/views/LeaderboardView.vue`

| Part | Lines (approx.) | What it does |
|------|---------------------|--------|
| `<script setup>` | 1–4 | Import Vue, Pinia, `useExamStore` |
| | 6–8 | `storeToRefs` — `leaderboard`, `leaderboardState`, `leaderboardError` |
| | 10–12 | `onMounted` → `loadLeaderboard()` |
| | 14–18 | Computed: ranks 1–3 and remaining rows (`slice(3)`) |
| | 20–34 | `formatScore` / `formatDate` · `backToExam` → `resetExam()` |
| `<template>` | ~35–230 | Header + loading/error/empty states · podium (1 / 2 / 3+ entries) · rank 4+ list via `v-for` · Back to Exam button |
| `<style scoped>` | End of file | Gradient for ranks 1–3, Material Symbols |

---

## Backend (run `go run ./cmd/api` from `backend/` or per project)

### 9. Entry — `backend/cmd/api/main.go`

| Lines (approx.) | What it does |
|---------------------|--------|
| 3–16 | Imports: `handler`, `repository`, `usecase`, `gin`, `godotenv` |
| 18–22 | `main()` calls `run()` |
| 24–25 | `godotenv.Load()` — optional `backend/.env` (`PORT`, `DATABASE_DIR`) |
| 27–34 | `resolveDataDir` inputs, `MkdirAll`, DSN `exam.db` |
| 36–45 | `OpenSQLite` → `AutoMigrate` → `EnsureSeedQuestions` |
| 47–51 | DI: `NewQuestionGorm`, `NewExamResultGorm` → `usecase.NewExam` → `handler.NewExamHTTP` |
| 53–61 | `gin.Default()`, `corsMiddleware`, `RegisterRoutes`, `Run` — `:8080` or `PORT` from env / `.env` |
| 85–96 | `corsMiddleware` — allows `GET`/`POST`/`OPTIONS` for `/api` |

### 10. Route registration — `backend/internal/handler/router.go`

| Lines | What it does |
|--------|--------|
| 8–14 | Group `/api`: `GET /questions`, `POST /submit`, `GET /leaderboard` |

### 11. HTTP + JSON — `backend/internal/handler/exam_handler.go`

| Lines (approx.) | What it does |
|---------------------|--------|
| 12–20 | struct `ExamHTTP`, constructor |
| 22–29 | `GetQuestions` → use case → `{ "questions": ... }` |
| 32–36 | `SubmitBody` — `candidateName`, `answers` |
| 38–55 | `Submit` — bind JSON, require non-empty `answers` → use case → 200 |
| 58–71 | `GetLeaderboard` — query `limit` (optional) → `{ "entries": ... }` |

### 12. Business rules — `backend/internal/usecase/exam_usecase.go`

| Lines (approx.) | What it does |
|---------------------|--------|
| 12–20 | struct `Exam` references `QuestionStore`, `ExamResultStore` (interfaces from `ports.go`) |
| 23–52 | DTOs for API + `SubmitResponse` + `LeaderboardEntryDTO` |
| 54–74 | `GetQuestions` — map `Question` → DTO **without answers** |
| 76–105 | `SubmitExam` — load questions → `ScoreAnswers` → build `ExamResult` + `SaveExamResult` |
| 107–125 | `GetLeaderboard` — load from store → assign `rank`, format `CreatedAt` as RFC3339 |
| 127–137 | `normalizeLeaderboardLimit` — default and max cap at 20 |
| 139–149 | `ScoreAnswers` — compare `answers["id"]` to `CorrectOptionID` |

### 13. Use case ports — `backend/internal/usecase/ports.go`

| Lines | What it does |
|--------|--------|
| 9–17 | `QuestionStore`, `ExamResultStore` (`SaveExamResult`, `GetLeaderboard`) — repository implements |

### 14. DB models — `backend/internal/models/question.go`, `exam_result.go`

- `question.go`: `Question`, `Option`, `CorrectOptionID` in DB
- `exam_result.go`: name, score, total, `AnswersJSON`, `CreatedAt`

### 15. GORM / SQLite — `backend/internal/repository/`

| File | What it does |
|------|--------|
| `gorm.go` | `OpenSQLite` |
| `migrate.go` | `AutoMigrate` Question, Option, ExamResult |
| `seed.go` | `EnsureSeedQuestions`, `seedQuestions()` — sample questions for API |
| `question_gorm.go` | `GetQuestions` — `Preload("Options")`, `Order("sort_order")` |
| `exam_result_gorm.go` | `SaveExamResult` — `Create` · `GetLeaderboard` — `ORDER BY score DESC`, `created_at ASC`, `Limit` |

### 16. Use case tests — `backend/internal/usecase/exam_usecase_test.go`

| What it does |
|--------|
| Mocks `QuestionStore` / `ExamResultStore`, tests `ScoreAnswers`, `SubmitExam`, `GetLeaderboard` |

---

## Suggested reading order

**Frontend:** `main.js` → `App.vue` → `router/index.js` → `api/client.js` + `vite.config.js` → `stores/examStore.js` → `views/ExamView.vue` → `views/ResultView.vue` → `views/LeaderboardView.vue`

**Backend:** `cmd/api/main.go` → `handler/router.go` → `handler/exam_handler.go` → `usecase/exam_usecase.go` + `ports.go` → `models/*` → `repository/*`

Then read the high-level flow in [architech.md](./architech.md)
