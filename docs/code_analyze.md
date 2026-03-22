# คู่มืออ่านโค้ด — โครงสร้างไฟล์และบรรทัด (Code map)

เอกสารนี้ใช้ **ไล่ตามลำดับที่โปรแกรมเริ่มทำงาน** ว่า import มาจากไหน แต่ละไฟล์ทำอะไร และชี้ **ช่วงบรรทัด** ที่ควรเปิดอ่านคู่กับ editor

เรื่อง **flow การทำงานรวม, flow usecase, flow ข้อมูล** อยู่ใน [architech.md](./architech.md)  
สัญญา API อยู่ใน [api.md](./api.md)

---



## ฝั่ง Frontend (รัน `npm run dev` ที่ `frontend/`)

### 1. จุดเริ่ม bundle — `frontend/src/main.js`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 1 | นำเข้า `createApp` จาก Vue |
| 2 | นำเข้า `createPinia` |
| 3 | นำเข้าราก `App.vue` |
| 4 | นำเข้า `router` จาก `./router` |
| 5 | นำเข้า CSS ทั่วทั้งแอป `./assets/main.css` |
| 7–11 | สร้างแอป → `use(Pinia)` → `use(router)` → `mount('#app')` |

ลำดับสำคัญ: **Pinia ก่อน** แล้วค่อย router เพื่อให้ทุกหน้าใช้ store ได้

### 2. ราก layout — `frontend/src/App.vue`

| ส่วน | ทำอะไร |
|------|--------|
| `<template>` | ห่อ `RouterView` — ไม่มีเมนู global; แต่ละ route เป็นหน้าเต็ม |

### 3. เส้นทาง — `frontend/src/router/index.js`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 1 | `createRouter`, `createWebHistory` |
| 3–28 | นิยาม `routes`: `/` → lazy load `ExamView`, `/result` → `ResultView`, `/leaderboard` → `LeaderboardView`, catch-all → `/` |
| 31–34 | `afterEach` ตั้ง `document.title` จาก `meta.title` |
| 36 | `export default router` — ถูก import ใน `main.js` และใน `examStore.js` |

### 4. HTTP ฝั่งเบราว์เซอร์ — `frontend/src/api/client.js`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 6–9 | `apiBase()` อ่าน `import.meta.env.VITE_API_BASE_URL` |
| 11–15 | `apiUrl(path)` ต่อ base หรือใช้ path สัมพัทธ์ `/api/...` |
| 17–38 | `fetchJSON` — ใส่ `Content-Type` เมื่อมี body, `JSON.parse`, โยน Error เมื่อ `!res.ok` |

ควบคู่กับ **proxy** ใน `frontend/vite.config.js` (โฟลเดอร์ `frontend/`): คีย์ `server.proxy['/api']` ส่งต่อไป `API_PROXY_TARGET` (ค่าเริ่ม `http://localhost:8080`)

### 5. State กลาง — `frontend/src/stores/examStore.js`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
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

### 6. หน้าทำข้อสอบ — `frontend/src/views/ExamView.vue` (**260 บรรทัด**)

ไฟล์เดียวกับโค้ดใน repo — แบ่งตามบล็อก `<script>` / `<template>` / `<style>`

#### `<script setup>` — บรรทัด 1–112

| บรรทัด | ทำอะไร |
|--------|--------|
| 1–4 | import `ref`, `computed`, `onMounted`, `nextTick` จาก Vue · `storeToRefs` จาก Pinia · `useExamStore` |
| 6–7 | สร้าง `exam` · `storeToRefs(exam)` → `candidateName`, `questions`, `answers`, `loadState`, `loadError` |
| 9–21 | `formError` · `showUnansweredHighlight` (Submit ไม่ครบ → ไฮไลต์ทุกข้อที่ยังไม่ตอบ) · `sectionRefs` + `setSectionRef` สำหรับ `scrollIntoView` |
| 23–25 | `onMounted(() => exam.loadQuestions())` |
| 27–30 | `allAnswered` (computed) — ถ้าไม่มีข้อ return false · ไม่เช่นนั้นทุกข้อต้องมีค่าใน `answers` |
| 32–38 | `isSelected` · `selectOption` → `exam.setAnswer` |
| 40–52 | `isQuestionUnanswered` · `questionSectionClasses` — กรอบแดง/พื้นหลังเมื่อ `showUnansweredHighlight` และข้อนั้นยังไม่ตอบ |
| 54–81 | `handleSubmit` — เคลียร์ `formError` / `showUnansweredHighlight` · ตรวจชื่อ · ถ้าไม่ครบ: ตั้งข้อความ + เปิดไฮไลต์ · `nextTick` แล้ว `scrollIntoView` ไปข้อแรกที่ว่าง · ไม่เช่นนั้น `submitExam` + `catch` network |
| 83–111 | `optionCardClasses` / `indicatorClasses` / `optionTextClasses` — การ์ดตัวเลือก |
| 112 | ปิด `</script>` |

#### `<template>` — บรรทัด 114–250

| บรรทัด | ทำอะไร |
|--------|--------|
| 115–117 | root `div` · เปิด `<main>` คอนเทนเนอร์ `max-w-3xl`, padding |
| 118–124 | `v-if="loadError"` — กล่องแจ้งเตือน (amber), `role="status"` |
| 125–161 | หัวหน้า + ชื่อผู้สอบ · `input#candidate-name` `v-model="candidateName"` |
| 163–172 | `v-if="loadState === 'loading'"` — spinner + ข้อความโหลด |
| 175–217 | `v-else` — `v-for` แต่ละ `section`: `:id="'question-' + q.id"` · `:ref` callback · `:class="questionSectionClasses(q.id)"` · ตัวเลือก · `v-if="showUnansweredHighlight && answers[q.id] == null"` ข้อความแดงในการ์ด |
| 219–240 | `v-if="loadState !== 'loading'"` — `formError` (`animate-pulse`) · ปุ่ม Submit |
| 241 | ปิด `</main>` |
| 243–249 | `div` ตกแต่งพื้นหลัง fixed + blur |
| 250 | ปิด root `</div>` |

#### `<style scoped>` — บรรทัด 252–260

| บรรทัด | ทำอะไร |
|--------|--------|
| 252–260 | `.material-symbols-outlined` — ตั้ง `font-variation-settings` (FILL, wght, GRAD, opsz) |

### 7. หน้าผล — `frontend/src/views/ResultView.vue`

| ส่วน | บรรทัด (โดยประมาณ) | ทำอะไร |
|------|---------------------|--------|
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

---

## ฝั่ง Backend (รัน `go run ./cmd/api` จาก `backend/` หรือตามที่โปรเจกต์กำหนด)

### 9. Entry process — `backend/cmd/api/main.go`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 3–15 | import: `handler`, `repository`, `usecase`, `gin` |
| 17–21 | `main()` เรียก `run()` |
| 23–31 | `resolveDataDir`, สร้างโฟลเดอร์ข้อมูล, path `exam.db` |
| 33–42 | `OpenSQLite` → `AutoMigrate` → `EnsureSeedQuestions` |
| 44–48 | DI: `NewQuestionGorm`, `NewExamResultGorm` → `usecase.NewExam` → `handler.NewExamHTTP` |
| 50–58 | `gin.Default()`, `corsMiddleware`, `RegisterRoutes`, `Run` — พอร์ต `:8080` หรือ `PORT` |
| 82–93 | `corsMiddleware` — อนุญาต `GET`/`POST`/`OPTIONS` สำหรับ `/api` |

### 10. ลงทะเบียน route — `backend/internal/handler/router.go`

| บรรทัด | ทำอะไร |
|--------|--------|
| 8–14 | กลุ่ม `/api`: `GET /questions`, `POST /submit`, `GET /leaderboard` |

### 11. HTTP + JSON — `backend/internal/handler/exam_handler.go`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 12–20 | struct `ExamHTTP`, constructor |
| 22–29 | `GetQuestions` → usecase → `{ "questions": ... }` |
| 32–36 | `SubmitBody` — `candidateName`, `answers` |
| 38–55 | `Submit` — bind JSON, ตรวจ `answers` ไม่ว่าง → usecase → 200 |
| 58–71 | `GetLeaderboard` — query `limit` (optional) → `{ "entries": ... }` |

### 12. กฎธุรกิจ — `backend/internal/usecase/exam_usecase.go`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 12–20 | struct `Exam` อ้าง `QuestionStore`, `ExamResultStore` (interface จาก `ports.go`) |
| 23–52 | DTO ส่งออก API + `SubmitResponse` + `LeaderboardEntryDTO` |
| 54–74 | `GetQuestions` — map `Question` → DTO **ไม่ใส่เฉลย** |
| 76–105 | `SubmitExam` — โหลดคำถาม → `ScoreAnswers` → สร้าง `ExamResult` + `SaveExamResult` |
| 107–125 | `GetLeaderboard` — `GetLeaderboard` จาก store → ใส่ `rank`, `CreatedAt` เป็น RFC3339 |
| 127–137 | `normalizeLeaderboardLimit` — ค่าเริ่มต้นและเพดาน 20 |
| 139–149 | `ScoreAnswers` — เทียบ `answers["id"]` กับ `CorrectOptionID` |

### 13. Interface ชั้น usecase — `backend/internal/usecase/ports.go`

| บรรทัด | ทำอะไร |
|--------|--------|
| 9–17 | `QuestionStore`, `ExamResultStore` (`SaveExamResult`, `GetLeaderboard`) — repository ต้อง implement |

### 14. โมเดล DB — `backend/internal/models/question.go`, `exam_result.go`

- `question.go`: `Question`, `Option`, ฟิลด์ `CorrectOptionID` ฝั่ง DB
- `exam_result.go`: บันทึกชื่อ, คะแนน, รวมข้อ, `AnswersJSON`, `CreatedAt`

### 15. GORM / SQLite — `backend/internal/repository/`

| ไฟล์ | ทำอะไร |
|------|--------|
| `gorm.go` | `OpenSQLite` |
| `migrate.go` | `AutoMigrate` ตาราง Question, Option, ExamResult |
| `seed.go` | `EnsureSeedQuestions`, `seedQuestions()` — ข้อสอบตัวอย่างใน SQLite สำหรับ API |
| `question_gorm.go` | `GetQuestions` — `Preload("Options")`, `Order("sort_order")` |
| `exam_result_gorm.go` | `SaveExamResult` — `Create` · `GetLeaderboard` — `ORDER BY score DESC`, `created_at ASC`, `Limit` |

### 16. ทดสอบ usecase — `backend/internal/usecase/exam_usecase_test.go`

| ทำอะไร |
|--------|
| mock `QuestionStore` / `ExamResultStore`, ทดสอบ `ScoreAnswers`, `SubmitExam`, `GetLeaderboard` |

---

## สรุปลำดับ “เปิดอ่าน” แนะนำ

**Frontend:** `main.js` → `App.vue` → `router/index.js` → `api/client.js` + `vite.config.js` → `stores/examStore.js` → `views/ExamView.vue` → `views/ResultView.vue` → `views/LeaderboardView.vue`

**Backend:** `cmd/api/main.go` → `handler/router.go` → `handler/exam_handler.go` → `usecase/exam_usecase.go` + `ports.go` → `models/*` → `repository/*`

จากนั้นอ่าน flow ภาพรวมใน [architech.md](./architech.md)
