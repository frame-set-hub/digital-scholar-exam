# Code map — โครงสร้างไฟล์และช่วงบรรทัด

เอกสารนี้ไล่ตาม **ลำดับการเริ่มทำงาน**: import มาจากไหน แต่ละไฟล์ทำอะไร และ **ช่วงบรรทัด** ที่ควรเปิดคู่กับ editor

ภาพรวม **flow, use case และข้อมูล** อยู่ใน [architech.md](./architech.md)  
สัญญา API อยู่ใน [api.md](./api.md)

---

## Frontend (รัน `npm run dev` ใน `frontend/`)

### 1. จุดเริ่ม bundle — `frontend/src/main.js`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 1 | import `createApp` จาก Vue |
| 2 | import `createPinia` |
| 3 | import root `App.vue` |
| 4 | import `router` จาก `./router` |
| 5 | import CSS ทั่วทั้งแอป `./assets/main.css` |
| 7–11 | สร้างแอป → `use(Pinia)` → `use(router)` → `mount('#app')` |

ลำดับสำคัญ: **Pinia ก่อน** แล้วค่อย router เพื่อให้ทุกหน้าใช้ store ได้

### 2. Root layout — `frontend/src/App.vue`

| ส่วน | ทำอะไร |
|------|--------|
| `<template>` | ห่อ `RouterView` — ไม่มีเมนู global; แต่ละ route เต็มหน้า |

### 3. เส้นทาง — `frontend/src/router/index.js`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 1 | `createRouter`, `createWebHistory` |
| 3–28 | นิยาม `routes`: `/` → lazy `ExamView`, `/result` → `ResultView`, `/leaderboard` → `LeaderboardView`, catch-all → `/` |
| 31–34 | `afterEach` ตั้ง `document.title` จาก `meta.title` |
| 36 | `export default router` — import ใน `main.js` และ `examStore.js` |

### 4. HTTP ฝั่งเบราว์เซอร์ — `frontend/src/api/client.js`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 6–9 | `apiBase()` อ่าน `import.meta.env.VITE_API_BASE_URL` |
| 11–15 | `apiUrl(path)` ต่อ base หรือใช้ path สัมพัทธ์ `/api/...` |
| 17–41 | `fetchJSON` — ใส่ `Content-Type` เมื่อมี body, `JSON.parse`, โยนเมื่อ `!res.ok` ด้วย `Error` ที่มี `status` และ `data` จาก response |

ร่วมกับ **proxy** ใน `frontend/vite.config.js`: `server.port` จาก `DEV_SERVER_PORT` ใน `.env` (ค่าเริ่ม `5173`); `server.proxy['/api']` ส่งต่อไป `API_PROXY_TARGET` (ค่าเริ่ม `http://localhost:8080`)

### 5. State กลาง — `frontend/src/stores/examStore.js`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 1–4 | import Pinia, Vue `ref`/`computed`, `router`, `apiUrl`/`fetchJSON` |
| 6–141 | `defineStore('exam', () => { ... })` — ไม่มีข้อสอบ mock ใน bundle; ดึงจาก API เท่านั้น |
| 7–10 | state: `candidateName`, `questions`, `answers`, `score` |
| 12–16 | `leaderboard`, `leaderboardYourEntry`, `leaderboardState`, `leaderboardError` |
| 18–19 | `loadState`, `loadError` |
| 21–22 | `loadQuestionsInflight` กันยิง GET ซ้ำพร้อมกัน |
| 24 | `totalQuestions` = `questions.length` |
| 26–28 | `setAnswer` |
| 34–65 | `loadQuestions` — เฉพาะ `GET /api/questions`; ล้มเหลวเคลียร์ `questions` + ตั้ง `loadError` |
| 67–98 | `loadLeaderboard` — `resolveLeaderboardCandidateName()` (store หรือ `route.query.forCandidate`) → `GET /api/leaderboard?forCandidate=` → `entries` + `yourEntry` |
| 92–98 | `answersForSubmit` — คีย์เป็น string ตามสัญญา API |
| 100–110 | `submitExam` — `POST /api/submit` แล้ว `router.push` ไป `result` |
| 112–121 | `resetExam` — เคลียร์ชื่อ/คำตอบ/คะแนน/leaderboard/`leaderboardYourEntry` กลับหน้าสอบ (คง `questions`) |
| 123–140 | `return` — สิ่งที่คอมโพเนนต์ใช้ |

### 6. หน้าสอบ — `frontend/src/views/ExamView.vue` (**~330 บรรทัด**)

ไฟล์เดียวกับใน repo — แบ่งตามบล็อก `<script>` / `<template>` / `<style>`

#### `<script setup>`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 1–4 | import `ref`, `computed`, `onMounted`, `nextTick` จาก Vue · `storeToRefs` · `useExamStore` |
| 6–7 | `exam` · `storeToRefs` → `candidateName`, `questions`, `answers`, `loadState`, `loadError` |
| 9–28 | `nameError` / `submitError` · `showUnansweredHighlight` · `sectionRefs` + `setSectionRef` · `nameBlockRef` / `submitSectionRef` / `submitBtnRef` |
| 31–33 | `onMounted(() => exam.loadQuestions())` |
| 35–46 | `allAnswered` · `isSelected` · `selectOption` |
| 48–61 | `isQuestionUnanswered` · `questionSectionClasses` (กรอบแดงเมื่อ highlight + ยังไม่ตอบ) |
| 63–77 | `scrollToName` · `scrollToSubmit` · `focusSubmitButton` (Enter ที่ชื่อ → โฟกัส Submit) |
| 79–128 | `handleSubmit` — เคลียร์ error ชื่อ/ส่งและ highlight · trim ชื่อ (ว่าง → `nameError` + เลื่อนไปชื่อ) · ถ้าไม่ `allAnswered`: `submitError` + `showUnansweredHighlight` + เลื่อนไปข้อแรกที่ว่าง · ไม่เช่นนั้น `submitExam()` · `catch`: `409`/`400` ข้อความชื่อ → `nameError` + เลื่อนไปชื่อ; อื่น → `submitError` + เลื่อนไปบริเวณส่ง |
| 130–158 | `optionCardClasses` / `indicatorClasses` / `optionTextClasses` |

#### `<template>`

| ส่วน | ทำอะไร |
|------|--------|
| Root | แบนเนอร์ amber เมื่อ `loadError` · หัว + `ref="nameBlockRef"` ช่องชื่อ (`nameError`, `aria-*`, Enter → `focusSubmitButton`) |
| คำถาม | `v-for` section พร้อม `:ref` → `setSectionRef`, `questionSectionClasses`, ข้อความเตือนในการ์ดเมื่อ highlight + ยังไม่ตอบ |
| ส่ง | `ref="submitSectionRef"` · `submitError` เหนือปุ่ม (`animate-pulse` เมื่อ `showUnansweredHighlight`) · `ref="submitBtnRef"` บนปุ่ม |

#### `<style scoped>`

| บรรทัด | ทำอะไร |
|--------|--------|
| ท้ายไฟล์ | `.material-symbols-outlined` — `font-variation-settings` |

### 7. หน้าผล — `frontend/src/views/ResultView.vue`

| ส่วน | บรรทัด (โดยประมาณ) | ทำอะไร |
|------|---------------------|--------|
| `<script setup>` | 1–5 | import Vue, `useRouter`, Pinia, store |
| | 11–15 | ถ้าไม่มี `score` → `replace` ไป `exam` |
| | 17–25 | คำนวณวงกลมความคืบหน้าคะแนน |
| | 27–33 | `retake` → `resetExam()` · `goLeaderboard` → `router.push` ไป `leaderboard` |
| `<template>` | ~36+ | ชื่อ, `score / totalQuestions`, ปุ่ม View Leaderboard + Retake Exam |

### 8. หน้า Leaderboard — `frontend/src/views/LeaderboardView.vue` (**~318 บรรทัด**)

| ส่วน | บรรทัด (โดยประมาณ) | ทำอะไร |
|------|---------------------|--------|
| `<script setup>` | 1–4 | import Vue, Pinia, `useExamStore` |
| | 7–8 | `storeToRefs` — รวม `leaderboardYourEntry`, `candidateName` |
| | 10–15 | `showMeHint` — แสดงคำแนะนำเมื่อไม่มีชื่อใน store (จึงไม่มี `forCandidate`) |
| | 17–19 | `onMounted` → `loadLeaderboard()` |
| | 20–23 | computed: อันดับ 1–3 และแถวที่เหลือ (`slice(3)`) |
| | 25–39 | `formatScore` / `formatDate` · `backToExam` → `resetExam()` |
| `<template>` | ~50+ | หัวข้อ · การ์ด **Me** เหนือ podium เมื่อมี `leaderboardYourEntry` · podium · รายการ · ปุ่ม Back to Exam |
| `<style scoped>` | ท้ายไฟล์ | gradient อันดับ 1–3, Material Symbols |

---

## Backend (รัน `go run ./cmd/api` จาก `backend/` หรือตามโปรเจกต์)

### 9. Entry — `backend/cmd/api/main.go`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 3–16 | import: `handler`, `repository`, `usecase`, `gin`, `godotenv` |
| 18–22 | `main()` เรียก `run()` |
| 24–25 | `godotenv.Load()` — `backend/.env` (optional) (`PORT`, `DATABASE_DIR`) |
| 27–34 | `resolveDataDir`, `MkdirAll`, DSN `exam.db` |
| 36–45 | `OpenSQLite` → `AutoMigrate` → `EnsureSeedQuestions` |
| 47–51 | DI: `NewQuestionGorm`, `NewExamResultGorm` → `usecase.NewExam` → `handler.NewExamHTTP` |
| 53–61 | `gin.Default()`, `corsMiddleware`, `RegisterRoutes`, `Run` — `:8080` หรือ `PORT` จาก env / `.env` |
| 85–96 | `corsMiddleware` — อนุญาต `GET`/`POST`/`OPTIONS` สำหรับ `/api` |

### 10. ลงทะเบียน route — `backend/internal/handler/router.go`

| บรรทัด | ทำอะไร |
|--------|--------|
| 8–14 | กลุ่ม `/api`: `GET /questions`, `POST /submit`, `GET /leaderboard` |

### 11. HTTP + JSON — `backend/internal/handler/exam_handler.go`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 12–20 | struct `ExamHTTP`, constructor |
| 22–29 | `GetQuestions` → use case → `{ "questions": ... }` |
| 32–36 | `SubmitBody` — `candidateName`, `answers` |
| 38–64 | `Submit` — bind JSON, บังคับ `answers` ไม่ว่าง → `SubmitExam` → `errors.Is` แมป `ErrCandidateNameRequired` → 400, `ErrDuplicateCandidateName` → 409, อื่น → 500 |
| 66–85 | `GetLeaderboard` — query `limit`, `forCandidate` → `{ "entries": ... }` และอาจมี `yourEntry` |

### 12. กฎธุรกิจ — `backend/internal/usecase/exam_usecase.go` + `errors.go`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 12–20 | struct `Exam` อ้าง `QuestionStore`, `ExamResultStore` (interface จาก `ports.go`) |
| `errors.go` | `ErrCandidateNameRequired`, `ErrDuplicateCandidateName` |
| 23–52 | DTO สำหรับ API + `SubmitResponse` + `LeaderboardEntryDTO` |
| 54–74 | `GetQuestions` — map `Question` → DTO **ไม่มีเฉลย** |
| 76–118 | `SubmitExam` — `strings.TrimSpace` ชื่อ → ว่าง → error · โหลดคำถาม → `ScoreAnswers` → `CandidateNameExists` → ซ้ำ → error · สร้าง `ExamResult` + `SaveExamResult` |
| 132–180 | `GetLeaderboard` — โหลดจาก store → ใส่ `rank`, จัดรูป `CreatedAt`; ถ้ามี `forCandidate` → `CandidateRank` → `LeaderboardYourEntryDTO` (`inTopList`) |
| | `normalizeLeaderboardLimit` · `ScoreAnswers` |

### 13. Ports ของ use case — `backend/internal/usecase/ports.go`

| บรรทัด | ทำอะไร |
|--------|--------|
| 9–21 | `QuestionStore`, `ExamResultStore` (`CandidateNameExists`, `SaveExamResult`, `GetLeaderboard`, `CandidateRank`) — repository implement |

### 14. โมเดล DB — `backend/internal/models/question.go`, `exam_result.go`

- `question.go`: `Question`, `Option`, `CorrectOptionID` ใน DB
- `exam_result.go`: ชื่อ, คะแนน, รวมข้อ, `AnswersJSON`, `CreatedAt`

### 15. GORM / SQLite — `backend/internal/repository/`

| ไฟล์ | ทำอะไร |
|------|--------|
| `gorm.go` | `OpenSQLite` |
| `migrate.go` | `AutoMigrate` Question, Option, ExamResult |
| `seed.go` | `EnsureSeedQuestions`, `seedQuestions()` — ข้อสอบตัวอย่างสำหรับ API |
| `question_gorm.go` | `GetQuestions` — `Preload("Options")`, `Order("sort_order")` |
| `exam_result_gorm.go` | `CandidateNameExists` — `WHERE candidate_name = ?` · `SaveExamResult` — `Create` · `GetLeaderboard` — `ORDER BY score DESC`, `created_at ASC`, `Limit` · `CandidateRank` — นับแถวที่ดีกว่าแถวของชื่อนั้น + 1 |

### 16. ทดสอบ use case — `backend/internal/usecase/exam_usecase_test.go`

| ทำอะไร |
|--------|
| mock `QuestionStore` / `ExamResultStore`, ทดสอบ `ScoreAnswers`, `SubmitExam`, `GetLeaderboard` |

---

## ลำดับการอ่านที่แนะนำ

**Frontend:** `main.js` → `App.vue` → `router/index.js` → `api/client.js` + `vite.config.js` → `stores/examStore.js` → `views/ExamView.vue` → `views/ResultView.vue` → `views/LeaderboardView.vue`

**Backend:** `cmd/api/main.go` → `handler/router.go` → `handler/exam_handler.go` → `usecase/exam_usecase.go` + `ports.go` → `models/*` → `repository/*`

จากนั้นอ่าน flow ภาพรวมใน [architech.md](./architech.md)
