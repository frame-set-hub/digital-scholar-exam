# Planning & Future Design (FE + BE)

## สารบัญ

- [Planning \& Future Design (FE + BE)](#planning--future-design-fe--be)
  - [สารบัญ](#สารบัญ)
  - [สถานะปัจจุบัน](#สถานะปัจจุบัน)
  - [เป้าหมายระยะสั้น](#เป้าหมายระยะสั้น)
    - [Leaderboard: ตำแหน่งของคุณเมื่ออยู่นอก top 20](#leaderboard-ตำแหน่งของคุณเมื่ออยู่นอก-top-20)
  - [เป้าหมายระยะยาว](#เป้าหมายระยะยาว)
  - [การรองรับข้อสอบหลายชุด](#การรองรับข้อสอบหลายชุด)
  - [Backend ที่มีอยู่แล้ว (อ้างอิง)](#backend-ที่มีอยู่แล้ว-อ้างอิง)
  - [การผสาน Frontend กับ Backend (ปัจจุบัน)](#การผสาน-frontend-กับ-backend-ปัจจุบัน)
  - [UX และความปลอดภัย](#ux-และความปลอดภัย)
  - [การทดสอบอัตโนมัติ](#การทดสอบอัตโนมัติ)

## สถานะปัจจุบัน

- **Frontend:** flow ครบ — Pinia โหลดข้อสอบจาก API (`GET /api/questions`) เท่านั้น
- **Backend:** API `GET /api/questions`, `POST /api/submit`, SQLite + seed อัตโนมัติ — ดู [`api.md`](./api.md)
- **Phase 7:** เพิ่มระบบ Leaderboard จากตาราง `ExamResult` เรียงตามคะแนน (`Score DESC`) และเวลา (`CreatedAt ASC`) — API `GET /api/leaderboard`, หน้า `/leaderboard` ใน Vue

## เป้าหมายระยะสั้น

- Frontend ต้องมี backend ที่รันอยู่ (หรือ proxy ไป API) เพื่อโหลดข้อสอบและส่งคำตอบ
- Backend เป็นข้อมูลจริงและเป็นผู้ตัดสินคะแนนหลังผสานแล้ว

### Leaderboard: ตำแหน่งของคุณเมื่ออยู่นอก top 20

**ปัญหาเดิม:** `GET /api/leaderboard` คืนได้สูงสุด **20** แถว — ถ้าอันดับรวมเป็น **21+** จะไม่เห็นตัวเองในรายการ

**สถานะปัจจุบัน (ทำแล้ว):**

- **API:** query `forCandidate` — อันดับรวมคำนวณจาก `exam_results` แบบเดียวกับการจัดอันดับ (score DESC, created_at ASC); response ฟิลด์ `yourEntry` พร้อม `inTopList` — ดู [`api.md`](./api.md)
- **FE:** `loadLeaderboard()` ส่ง `forCandidate` เมื่อมี `candidateName` ใน Pinia; `LeaderboardView` แสดงแถบ **Your position** เมื่อ `yourEntry` มี `inTopList: false`

## เป้าหมายระยะยาว

เอกสารนี้กำหนดทิศทางเมื่อระบบโตเกิน “เครื่องเดียว / user เดียว” — ยังไม่ใช่สเปกลงมือทันที แต่ช่วยตัดสิน stack และลำดับงานภายหลัง

### Scale และโหลด

- **แนวนอน:** ทำ API ให้ **stateless** (session ไม่ผูก memory เครื่อง) เพื่อให้หลาย instance อยู่หลัง load balancer ได้
- **ฐานข้อมูล:** SQLite เหมาะ dev/สาธิต — เมื่อ concurrent เขียนสูงหรือต้อง backup/HA ให้เลือก **PostgreSQL** (หรือ MySQL); พิจารณา read replica ถ้าอ่านหนัก
- **ขีดจำกัด:** มักอยู่ที่ DB และการออกแบบ query/transaction ไม่ใช่ Go/Vue เอง

### ผู้ใช้หลายคน (และบทบาท)

- **ผู้สอบจำนวนมาก:** ต้องมี **ตัวตน** — ลงทะเบียน/ล็อกอิน หรือ SSO — ไม่ใช่แค่สตริงใน `candidateName`
- **บทบาท:** แยก **ผู้ดูแลข้อสอบ** กับ **ผู้สอบ**; อาจขยายเป็น org/คลาส (multi-tenant) ภายหลัง
- **Trade-off:** auth ฝั่ง backend (JWT / session cookie + HTTPS) ชัดกว่าเชื่อ token จาก client อย่างเดียว

### Deploy บน cloud / production

- **Frontend:** static build (Vite) ไป **CDN / object storage** (เช่น S3 + CloudFront) หรือ **Pages**; ตั้ง `VITE_API_BASE_URL` ชี้ API จริง
- **Backend:** binary Go ใน **container** (Docker) บน **ECS, Cloud Run, Fly.io, Railway** ฯลฯ — เลือกตามงบ ทีม และความคุ้นเคย
- **สภาพแวดล้อม:** แยก `dev` / `staging` / `prod`, เก็บ secret นอก repo (env แพลตฟอร์ม, vault)

### Tradeoff และการเลือก stack

| หัวข้อ | ทางเลือกหลัก | หมายเหตุ |
|--------|----------------|-----------|
| DB | SQLite → PostgreSQL | ย้ายเมื่อ concurrent write สำคัญ; migration มีเวอร์ชัน (เช่น golang-migrate) |
| Cache / session | ไม่มี → Redis | เมื่อ session คร่อมหลาย instance หรือ rate limit |
| Queue | ไม่มี → SQS / Rabbit / NATS | งานหนักหลัง submit (อีเมล, รายงาน) ไม่ควรบล็อก request |
| Monolith vs บริการ | คง Go monolith ก่อน | แยกเมื่อขอบเขตชัด (เช่น grading service) — อย่าแยกเร็วเกิน |
| Observability | log ธรรมดา → structured log + metrics + tracing | ช่วย debug บน production เมื่อ scale |

สรุป: **Go + Gin + GORM + Vue** ใช้เป็นฐานได้นาน — ส่วนที่เปลี่ยนก่อนมักเป็น **ฐานข้อมูล, auth, และ deploy** ไม่ใช่รีเขียนทั้งระบบ

## การรองรับข้อสอบหลายชุด

- นิยาม **Exam** เป็น entity: `id`, `title`, `slug`, `version`
- **Backend:** เพิ่มตาราง `exams`, `questions.exam_id`; API เช่น `GET /api/exams/:id/questions`
- **Frontend:** route `/exam/:examId`, store แยก catalog กับ session หรือ namespace ใน Pinia

## Backend ที่มีอยู่แล้ว (อ้างอิง)

| รายการ | รายละเอียด |
|--------|-------------|
| Entry | `backend/cmd/api/main.go` |
| DB | SQLite `backend/data/exam.db`, GORM `AutoMigrate` + seed เมื่อยังว่าง |
| API | `GET /api/questions` — ไม่ส่งเฉลย; `POST /api/submit` — `{ candidateName, answers }` → `{ candidateName, score, total }` และบันทึก `exam_results`; `GET /api/leaderboard` — อันดับผู้สอบ (ไม่ส่งคำตอบดิบ) |
| ชั้นโค้ด | `handler` → `usecase` → `repository` — ดู [architech.md](./architech.md) |

## การผสาน Frontend กับ Backend (ปัจจุบัน)

- Dev: Vite proxy `/api` → `http://localhost:8080` — `examStore.loadQuestions()` / `submitExam()` / `loadLeaderboard()` เรียก API
- โหลด: `GET /api/questions` — ล้มเหลวแสดง error; ไม่มีชุดข้อสอบในเครื่อง
- ส่ง: `POST /api/submit` — คะแนนจากเซิร์ฟเวอร์เท่านั้น
- Leaderboard: `GET /api/leaderboard` — `LeaderboardView` เรียก `loadLeaderboard()` แสดงรายการจาก `exam_results` (ไม่ส่งคำตอบดิบ)
- Production: ตั้ง `VITE_API_BASE_URL` หรือ reverse proxy ร่วม host — ดู `frontend/.env.example`

## UX และความปลอดภัย

- การยืนยันตัวผู้สอบ (ถ้ามี) — มักทำที่ BE (JWT / session) แล้ว FE ส่ง token
- จำกัดเวลา, autosave ระหว่างทำข้อสอบ
- คะแนนที่เชื่อถือได้: เซิร์ฟเวอร์เป็นตัวตัดสินกับ `POST /api/submit` (ไม่เชื่อถือ client อย่างเดียว)

## การทดสอบอัตโนมัติ

- **Backend:** `cd backend && go test ./...` — use case + mock repository (มีแล้ว)
- **Frontend:** Vitest สำหรับ store/views เมื่อเพิ่ม

บันทึกการรันและรายละเอียด: [testing.md](./testing.md)
