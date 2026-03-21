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
| 3–22 | นิยาม `routes`: `/` → lazy load `ExamView`, `/result` → `ResultView`, catch-all → `/` |
| 25–28 | `afterEach` ตั้ง `document.title` จาก `meta.title` |
| 30 | `export default router` — ถูก import ใน `main.js` และใน `examStore.js` |

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
| 6–101 | `defineStore('exam', () => { ... })` — ไม่มีข้อสอบ mock ใน bundle; ดึงจาก API เท่านั้น |
| 7–10 | state: `candidateName`, `questions`, `answers`, `score` |
| 12–13 | `loadState`, `loadError` |
| 15–16 | `loadQuestionsInflight` กันยิง GET ซ้ำพร้อมกัน |
| 18 | `totalQuestions` = `questions.length` |
| 20–22 | `setAnswer` |
| 28–59 | `loadQuestions` — `GET /api/questions` เท่านั้น; ล้มเหลวเคลียร์ `questions` + ตั้ง `loadError` |
| 61–67 | `answersForSubmit` — คีย์เป็น string ตามสัญญา API |
| 69–79 | `submitExam` — `POST /api/submit` แล้ว `router.push` ไป `result` |
| 81–86 | `resetExam` — เคลียร์ชื่อ/คำตอบ/คะแนน กลับหน้าสอบ (ไม่เคลียร์ `questions`) |
| 88–100 | `return` สิ่งที่คอมโพเนนต์ใช้ได้ |

### 6. หน้าทำข้อสอบ — `frontend/src/views/ExamView.vue`

| ส่วน | บรรทัด (โดยประมาณ) | ทำอะไร |
|------|---------------------|--------|
| `<script setup>` | 1–4 | import Vue, Pinia `storeToRefs`, `useExamStore` |
| | 6–7 | ดึง ref จาก store |
| | 11–13 | `onMounted` → `exam.loadQuestions()` |
| | 15–18 | `allAnswered` — มีข้ออย่างน้อยหนึ่งข้อ และทุกข้อมีคำตอบ |
| | 19–25 | เลือกข้อ `selectOption` / `isSelected` |
| | 27–46 | `handleSubmit` — validate ชื่อ + ครบข้อ → `submitExam` + จับ error |
| | 48–76 | ฟังก์ชันคลาส Tailwind สำหรับการ์ดตัวเลือก |
| `<template>` | เริ่ม ~79 | layout หลัก, แบนเนอร์ `loadError`, หัวข้อ, ช่องชื่อ |
| | ~128–137 | spinner เมื่อ `loadState === 'loading'` |
| | ~139–172 | `v-for` ข้อและปุ่มตัวเลือก |
| | ~174–191 | ปุ่ม Submit |

### 7. หน้าผล — `frontend/src/views/ResultView.vue`

| ส่วน | บรรทัด (โดยประมาณ) | ทำอะไร |
|------|---------------------|--------|
| `<script setup>` | 1–5 | import Vue, `useRouter`, Pinia, store |
| | 11–15 | ถ้าไม่มี `score` → `replace` กลับ `exam` |
| | 17–25 | คำนวณวงกลมความคืบหน้าคะแนน |
| | 27–29 | `retake` → `resetExam()` |
| `<template>` | ต่อจาก ~32 | แสดงชื่อ, คะแนน `score / totalQuestions`, ปุ่ม Retake |

---

## ฝั่ง Backend (รัน `go run ./cmd/api` จาก `backend/` หรือตามที่โปรเจกต์กำหนด)

### 8. Entry process — `backend/cmd/api/main.go`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 8–14 | import: `handler`, `repository`, `usecase`, `gin` |
| 16–20 | `main()` เรียก `run()` |
| 22–27 | สร้างโฟลเดอร์ `data/`, path `data/exam.db` |
| 29–35 | `OpenSQLite` → `AutoMigrate` |
| 36–37 | `EnsureSeedQuestions` — ใส่ข้อที่ยังไม่มีใน DB |
| 40–44 | DI: `NewQuestionGorm`, `NewExamResultGorm` → `usecase.NewExam` → `handler.NewExamHTTP` |
| 46–48 | `gin.Default()`, `corsMiddleware`, `RegisterRoutes` |
| 50–54 | พอร์ต `:8080` หรือ `PORT` |

### 9. ลงทะเบียน route — `backend/internal/handler/router.go`

| บรรทัด | ทำอะไร |
|--------|--------|
| 8–13 | กลุ่ม `/api`: `GET /questions`, `POST /submit` |

### 10. HTTP + JSON — `backend/internal/handler/exam_handler.go`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 11–19 | struct `ExamHTTP`, constructor |
| 21–29 | `GetQuestions` → usecase → `{ "questions": ... }` |
| 31–35 | `SubmitBody` — `candidateName`, `answers` |
| 37–55 | `Submit` — bind JSON, ตรวจ `answers` ไม่ว่าง → usecase → 200 |

### 11. กฎธุรกิจ — `backend/internal/usecase/exam_usecase.go`

| บรรทัด (โดยประมาณ) | ทำอะไร |
|---------------------|--------|
| 11–20 | struct `Exam` อ้าง `QuestionStore`, `ExamResultStore` (interface จาก `ports.go`) |
| 22–42 | DTO ส่งออก API + `SubmitResponse` |
| 44–64 | `GetQuestions` — map `Question` → DTO **ไม่ใส่เฉลย** |
| 66–95 | `SubmitExam` — โหลดคำถาม → `ScoreAnswers` → สร้าง `ExamResult` + `SaveExamResult` |
| 97–107 | `ScoreAnswers` — เทียบ `answers["id"]` กับ `CorrectOptionID` |

### 12. Interface ชั้น usecase — `backend/internal/usecase/ports.go`

| บรรทัด | ทำอะไร |
|--------|--------|
| 9–17 | `QuestionStore`, `ExamResultStore` — repository ต้อง implement |

### 13. โมเดล DB — `backend/internal/models/question.go`, `exam_result.go`

- `question.go`: `Question`, `Option`, ฟิลด์ `CorrectOptionID` ฝั่ง DB
- `exam_result.go`: บันทึกชื่อ, คะแนน, รวมข้อ, `AnswersJSON`, `CreatedAt`

### 14. GORM / SQLite — `backend/internal/repository/`

| ไฟล์ | ทำอะไร |
|------|--------|
| `gorm.go` | `OpenSQLite` |
| `migrate.go` | `AutoMigrate` ตาราง Question, Option, ExamResult |
| `seed.go` | `EnsureSeedQuestions`, `seedQuestions()` — ข้อสอบตัวอย่างใน SQLite สำหรับ API |
| `question_gorm.go` | `GetQuestions` — `Preload("Options")`, `Order("sort_order")` |
| `exam_result_gorm.go` | `SaveExamResult` — `Create` |

### 15. ทดสอบ usecase — `backend/internal/usecase/exam_usecase_test.go`

| ทำอะไร |
|--------|
| mock `QuestionStore` / `ExamResultStore`, ทดสอบ `ScoreAnswers` และ `SubmitExam` |

---

## สรุปลำดับ “เปิดอ่าน” แนะนำ

**Frontend:** `main.js` → `App.vue` → `router/index.js` → `api/client.js` + `vite.config.js` → `stores/examStore.js` → `views/ExamView.vue` → `views/ResultView.vue`

**Backend:** `cmd/api/main.go` → `handler/router.go` → `handler/exam_handler.go` → `usecase/exam_usecase.go` + `ports.go` → `models/*` → `repository/*`

จากนั้นอ่าน flow ภาพรวมใน [architech.md](./architech.md)
