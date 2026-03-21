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
<<<<<<< HEAD
| 3–28 | นิยาม `routes`: `/` → lazy load `ExamView`, `/result` → `ResultView`, `/leaderboard` → `LeaderboardView`, catch-all → `/` |
| 31–34 | `afterEach` ตั้ง `document.title` จาก `meta.title` |
| 36 | `export default router` — ถูก import ใน `main.js` และใน `examStore.js` |
=======
| 3–22 | Define `routes`: `/` → lazy `ExamView`, `/result` → `ResultView`, catch-all → `/` |
| 25–28 | `afterEach` sets `document.title` from `meta.title` |
| 30 | `export default router` — imported in `main.js` and `examStore.js` |
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

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
<<<<<<< HEAD
| 1–4 | import Pinia, Vue `ref`/`computed`, `router`, `apiUrl`/`fetchJSON` |
| 6–129 | `defineStore('exam', () => { ... })` — ไม่มีข้อสอบ mock ใน bundle; ดึงจาก API เท่านั้น |
| 7–10 | state: `candidateName`, `questions`, `answers`, `score` |
| 12–14 | `leaderboard`, `leaderboardState`, `leaderboardError` |
| 16–17 | `loadState`, `loadError` |
| 19–20 | `loadQuestionsInflight` กันยิง GET ซ้ำพร้อมกัน |
| 22 | `totalQuestions` = `questions.length` |
| 24–26 | `setAnswer` |
| 32–63 | `loadQuestions` — `GET /api/questions` เท่านั้น; ล้มเหลวเคลียร์ `questions` + ตั้ง `loadError` |
| 65–80 | `loadLeaderboard` — `GET /api/leaderboard` → `entries` ลง `leaderboard` |
| 82–88 | `answersForSubmit` — คีย์เป็น string ตามสัญญา API |
| 90–100 | `submitExam` — `POST /api/submit` แล้ว `router.push` ไป `result` |
| 102–110 | `resetExam` — เคลียร์ชื่อ/คำตอบ/คะแนน/leaderboard กลับหน้าสอบ (ไม่เคลียร์ `questions`) |
| 112–128 | `return` สิ่งที่คอมโพเนนต์ใช้ได้ |

### 6. หน้าทำข้อสอบ — `frontend/src/views/ExamView.vue` (**212 บรรทัด**)

ไฟล์เดียวกับโค้ดใน repo — แบ่งตามบล็อก `<script>` / `<template>` / `<style>`

#### `<script setup>` — บรรทัด 1–78

| บรรทัด | ทำอะไร |
|--------|--------|
| 1–4 | import `ref`, `computed`, `onMounted` จาก Vue · `storeToRefs` จาก Pinia · `useExamStore` |
| 6–7 | สร้าง `exam` · `storeToRefs(exam)` → `candidateName`, `questions`, `answers`, `loadState`, `loadError` |
| 9 | `formError = ref('')` สำหรับ validation ฝั่งฟอร์มและข้อความ error ตอนส่งข้อสอบ |
| 11–13 | `onMounted(() => exam.loadQuestions())` |
| 15–18 | `allAnswered` (computed) — ถ้าไม่มีข้อ return false · ไม่เช่นนั้นทุกข้อต้องมีค่าใน `answers` |
| 20–22 | `isSelected(questionId, optionId)` — เทียบ `answers[questionId] === optionId` |
| 24–26 | `selectOption` → `exam.setAnswer(questionId, optionId)` |
| 28–47 | `handleSubmit` — เคลียร์ `formError` · ตรวจชื่อไม่ว่าง · ตรวจ `allAnswered` · `await exam.submitExam()` · `catch` ตั้งข้อความไทยเมื่อ network/TypeError |
| 49–59 | `optionCardClasses` — คืน array คลาส Tailwind สำหรับการ์ดตัวเลือก (border, พื้นหลัง, เงาเมื่อเลือก) |
| 61–69 | `indicatorClasses` — วงกลมตัวอักษรตัวเลือก (A/B/C) |
| 71–77 | `optionTextClasses` — ข้อความคำอธิบายตัวเลือก |
| 78 | ปิด `</script>` |

#### `<template>` — บรรทัด 80–202

| บรรทัด | ทำอะไร |
|--------|--------|
| 80–82 | root `div` (`min-h-screen`, `bg-background`) · เปิด `<main>` คอนเทนเนอร์ `max-w-3xl`, padding |
| 84–90 | `v-if="loadError"` — กล่องแจ้งเตือน (amber) แสดง `loadError`, `role="status"` |
| 92–127 | บล็อกหัวหน้า + ชื่อผู้สอบ: badge “Live Session”, หัวข้อ “IT 10-1 Exam”, คำอธิบายโมดูล · `label` + `input#candidate-name` `v-model="candidateName"` · เส้นใต้เมื่อ focus |
| 129–138 | `v-if="loadState === 'loading'"` — spinner + ข้อความ “กำลังโหลดข้อสอบจากเซิร์ฟเวอร์…” |
| 141–173 | `v-else` + `space-y-12` — `v-for="(q, index) in questions"` แต่ละ `section`: หมายเลขข้อ · `q.prompt` / `q.subtitle` · `v-for="opt in q.options"` ปุ่ม `button` เรียก `selectOption` · ผูกคลาสจาก `optionCardClasses` / `indicatorClasses` / `optionTextClasses` |
| 175–192 | `v-if="loadState !== 'loading'"` — แสดง `formError` (สีแดง) · ปุ่ม Submit (`:disabled="questions.length === 0"`) · `@click="handleSubmit"` · gradient + `material-symbols-outlined` arrow |
| 193 | ปิด `</main>` |
| 195–201 | `div` ตกแต่งพื้นหลังแบบ fixed + blur (มุมขวาบน / ซ้ายล่าง) |
| 202 | ปิด root `</div>` |

#### `<style scoped>` — บรรทัด 204–212

| บรรทัด | ทำอะไร |
|--------|--------|
| 204–212 | `.material-symbols-outlined` — ตั้ง `font-variation-settings` (FILL, wght, GRAD, opsz) |
=======
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
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

### 7. Result view — `frontend/src/views/ResultView.vue`

| Part | Lines (approx.) | What it does |
|------|---------------------|--------|
<<<<<<< HEAD
| `<script setup>` | 1–5 | import Vue, `useRouter`, Pinia, store |
| | 11–15 | ถ้าไม่มี `score` → `replace` กลับ `exam` |
| | 17–25 | คำนวณวงกลมความคืบหน้าคะแนน |
| | 27–33 | `retake` → `resetExam()` · `goLeaderboard` → `router.push` ไป `leaderboard` |
| `<template>` | ต่อจาก ~36 | แสดงชื่อ, คะแนน `score / totalQuestions`, ปุ่ม View Leaderboard + Retake Exam |

### 8. หน้ากระดานอันดับ — `frontend/src/views/LeaderboardView.vue`

| ส่วน | บรรทัด (โดยประมาณ) | ทำอะไร |
|------|---------------------|--------|
| `<script setup>` | 1–4 | import Vue, Pinia, `useExamStore` |
| | 6–8 | `storeToRefs` — `leaderboard`, `leaderboardState`, `leaderboardError` |
| | 10–12 | `onMounted` → `loadLeaderboard()` |
| | 14–18 | computed: อันดับ 1–3 และแถวที่เหลือ (`slice(3)`) |
| | 20–34 | `formatScore` / `formatDate` · `backToExam` → `resetExam()` |
| `<template>` | ~35–230 | header + สถานะโหลด/error/ว่าง · podium (1 / 2 / 3+ คน) · รายการอันดับ 4+ ด้วย `v-for` · ปุ่ม Back to Exam |
| `<style scoped>` | ท้ายไฟล์ | gradient อันดับ 1–3, Material Symbols |
=======
| `<script setup>` | 1–5 | Import Vue, `useRouter`, Pinia, store |
| | 11–15 | If no `score` → `replace` to `exam` |
| | 17–25 | Score progress circle |
| | 27–29 | `retake` → `resetExam()` |
| `<template>` | ~32+ | Name, `score / totalQuestions`, Retake button |
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

---

## Backend (run `go run ./cmd/api` from `backend/` or per project)

<<<<<<< HEAD
### 9. Entry process — `backend/cmd/api/main.go`
=======
### 8. Entry — `backend/cmd/api/main.go`
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

| Lines (approx.) | What it does |
|---------------------|--------|
<<<<<<< HEAD
| 3–15 | import: `handler`, `repository`, `usecase`, `gin` |
| 17–21 | `main()` เรียก `run()` |
| 23–31 | `resolveDataDir`, สร้างโฟลเดอร์ข้อมูล, path `exam.db` |
| 33–42 | `OpenSQLite` → `AutoMigrate` → `EnsureSeedQuestions` |
| 44–48 | DI: `NewQuestionGorm`, `NewExamResultGorm` → `usecase.NewExam` → `handler.NewExamHTTP` |
| 50–58 | `gin.Default()`, `corsMiddleware`, `RegisterRoutes`, `Run` — พอร์ต `:8080` หรือ `PORT` |
| 82–93 | `corsMiddleware` — อนุญาต `GET`/`POST`/`OPTIONS` สำหรับ `/api` |

### 10. ลงทะเบียน route — `backend/internal/handler/router.go`
=======
| 8–14 | import: `handler`, `repository`, `usecase`, `gin` |
| 16–20 | `main()` calls `run()` |
| 22–27 | Create `data/` folder, path `data/exam.db` |
| 29–35 | `OpenSQLite` → `AutoMigrate` |
| 36–37 | `EnsureSeedQuestions` — seed if missing |
| 40–44 | DI: `NewQuestionGorm`, `NewExamResultGorm` → `usecase.NewExam` → `handler.NewExamHTTP` |
| 46–48 | `gin.Default()`, `corsMiddleware`, `RegisterRoutes` |
| 50–54 | Port `:8080` or `PORT` |

### 9. Route registration — `backend/internal/handler/router.go`
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

| Lines | What it does |
|--------|--------|
<<<<<<< HEAD
| 8–14 | กลุ่ม `/api`: `GET /questions`, `POST /submit`, `GET /leaderboard` |
=======
| 8–13 | Group `/api`: `GET /questions`, `POST /submit` |
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

### 11. HTTP + JSON — `backend/internal/handler/exam_handler.go`

| Lines (approx.) | What it does |
|---------------------|--------|
<<<<<<< HEAD
| 12–20 | struct `ExamHTTP`, constructor |
| 22–29 | `GetQuestions` → usecase → `{ "questions": ... }` |
| 32–36 | `SubmitBody` — `candidateName`, `answers` |
| 38–55 | `Submit` — bind JSON, ตรวจ `answers` ไม่ว่าง → usecase → 200 |
| 58–71 | `GetLeaderboard` — query `limit` (optional) → `{ "entries": ... }` |

### 12. กฎธุรกิจ — `backend/internal/usecase/exam_usecase.go`
=======
| 11–19 | struct `ExamHTTP`, constructor |
| 21–29 | `GetQuestions` → use case → `{ "questions": ... }` |
| 31–35 | `SubmitBody` — `candidateName`, `answers` |
| 37–55 | `Submit` — bind JSON, require non-empty `answers` → use case → 200 |

### 11. Business rules — `backend/internal/usecase/exam_usecase.go`
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

| Lines (approx.) | What it does |
|---------------------|--------|
<<<<<<< HEAD
| 12–20 | struct `Exam` อ้าง `QuestionStore`, `ExamResultStore` (interface จาก `ports.go`) |
| 23–52 | DTO ส่งออก API + `SubmitResponse` + `LeaderboardEntryDTO` |
| 54–74 | `GetQuestions` — map `Question` → DTO **ไม่ใส่เฉลย** |
| 76–105 | `SubmitExam` — โหลดคำถาม → `ScoreAnswers` → สร้าง `ExamResult` + `SaveExamResult` |
| 107–125 | `GetLeaderboard` — `GetLeaderboard` จาก store → ใส่ `rank`, `CreatedAt` เป็น RFC3339 |
| 127–137 | `normalizeLeaderboardLimit` — ค่าเริ่มต้นและเพดาน 20 |
| 139–149 | `ScoreAnswers` — เทียบ `answers["id"]` กับ `CorrectOptionID` |

### 13. Interface ชั้น usecase — `backend/internal/usecase/ports.go`
=======
| 11–20 | struct `Exam` references `QuestionStore`, `ExamResultStore` (interfaces from `ports.go`) |
| 22–42 | DTOs for API + `SubmitResponse` |
| 44–64 | `GetQuestions` — map `Question` → DTO **without answers** |
| 66–95 | `SubmitExam` — load questions → `ScoreAnswers` → build `ExamResult` + `SaveExamResult` |
| 97–107 | `ScoreAnswers` — compare `answers["id"]` to `CorrectOptionID` |

### 12. Use case ports — `backend/internal/usecase/ports.go`
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

| Lines | What it does |
|--------|--------|
<<<<<<< HEAD
| 9–17 | `QuestionStore`, `ExamResultStore` (`SaveExamResult`, `GetLeaderboard`) — repository ต้อง implement |

### 14. โมเดล DB — `backend/internal/models/question.go`, `exam_result.go`
=======
| 9–17 | `QuestionStore`, `ExamResultStore` — repository implements |

### 13. DB models — `backend/internal/models/question.go`, `exam_result.go`
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

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

<<<<<<< HEAD
### 16. ทดสอบ usecase — `backend/internal/usecase/exam_usecase_test.go`
=======
### 15. Use case tests — `backend/internal/usecase/exam_usecase_test.go`
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

| What it does |
|--------|
<<<<<<< HEAD
| mock `QuestionStore` / `ExamResultStore`, ทดสอบ `ScoreAnswers`, `SubmitExam`, `GetLeaderboard` |
=======
| Mocks `QuestionStore` / `ExamResultStore`, tests `ScoreAnswers` and `SubmitExam` |
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

---

## Suggested reading order

**Frontend:** `main.js` → `App.vue` → `router/index.js` → `api/client.js` + `vite.config.js` → `stores/examStore.js` → `views/ExamView.vue` → `views/ResultView.vue` → `views/LeaderboardView.vue`

**Backend:** `cmd/api/main.go` → `handler/router.go` → `handler/exam_handler.go` → `usecase/exam_usecase.go` + `ports.go` → `models/*` → `repository/*`

Then read the high-level flow in [architech.md](./architech.md)
