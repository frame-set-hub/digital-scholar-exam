# Planning & Future Design (FE + BE)

## สารบัญ

- [Planning \& Future Design (FE + BE)](#planning--future-design-fe--be)
  - [สารบัญ](#สารบัญ)
  - [สถานะปัจจุบัน](#สถานะปัจจุบัน)
  - [เป้าหมายระยะสั้น](#เป้าหมายระยะสั้น)
  - [เป้าหมายระยะยาว](#เป้าหมายระยะยาว)
  - [การขยายรองรับข้อสอบหลายชุด](#การขยายรองรับข้อสอบหลายชุด)
  - [Backend ที่มีอยู่แล้ว (อ้างอิง)](#backend-ที่มีอยู่แล้ว-อ้างอิง)
  - [การผสาน Frontend กับ Backend (สถานะปัจจุบัน)](#การผสาน-frontend-กับ-backend-สถานะปัจจุบัน)
  - [ประสบการณ์ผู้ใช้และความปลอดภัย](#ประสบการณ์ผู้ใช้และความปลอดภัย)
  - [การทดสอบอัตโนมัติ](#การทดสอบอัตโนมัติ)

## สถานะปัจจุบัน

- **Frontend:** flow ครบ — Pinia โหลดข้อสอบจาก API (`GET /api/questions`) เท่านั้น
- **Backend:** มี API `GET /api/questions`, `POST /api/submit`, SQLite + seed อัตโนมัติ — ดู [`api.md`](./api.md)

## เป้าหมายระยะสั้น

- FE ต้องมี backend พร้อม (หรือ proxy ชี้ API) เพื่อโหลดข้อสอบและส่งคำตอบ
- BE พร้อมเป็นข้อมูลจริงและเป็นผู้ตัดสินคะแนนเมื่อผสานแล้ว

## เป้าหมายระยะยาว

เอกสารนี้ใช้วางทิศทางเมื่อระบบต้องโตเกิน “เครื่องเดียว / user เดียว” — ยังไม่ใช่สเปกลงมือทำทันที แต่ช่วยตัดสินใจ stack และลำดับงานภายหลัง

### Scale และโหลด

- **แนวนอน:** ทำ API ให้ **stateless** (session ไม่ผูก memory เครื่อง) จึง scale instance ได้หลายตัวหลัง load balancer
- **ฐานข้อมูล:** SQLite เหมาะ dev/สาธิต — เมื่อ concurrent เขียนสูงหรือต้อง backup/HA ให้วาง **PostgreSQL** (หรือ MySQL) เป็นเป้าหมายถัดไป; อาจแยก read replica ถ้าอ่านหนัก
- **ขีดจำกัด:** จุดคอขวดมักอยู่ที่ DB และการออกแบบ query/transaction ไม่ใช่ภาษา Go/Vue เอง

### ผู้ใช้มากกว่าหนึ่งคน (และบทบาท)

- **ผู้สอบหลายคน:** ต้องมี **ตัวตน (identity)** — ลงทะเบียน/ล็อกอิน หรือเชื่อม SSO — ไม่พึ่งแค่ชื่อสตริงใน `candidateName`
- **บทบาท:** แยก **ผู้ดูแลข้อสอบ** กับ **ผู้ทำข้อสอบ**; อาจขยายเป็นองค์กร/คลาส (multi-tenant) ภายหลัง
- **Tradeoff:** auth ฝั่ง backend (JWT / session cookie + HTTPS) ชัดกว่าเชื่อ token จาก client อย่างเดียว

### Deploy บน cloud / production

- **Frontend:** static build (Vite) ไป **CDN / object storage** (เช่น S3 + CloudFront) หรือ **Pages** ของผู้ให้บริการ; ตั้ง `VITE_API_BASE_URL` ชี้ API จริง
- **Backend:** binary Go ใน **container** (Docker) แล้วรันบน **ECS, Cloud Run, Fly.io, Railway** ฯลฯ — เลือกตามงบ ทีม และความคุ้นเคย
- **สิ่งแวดล้อม:** แยก `dev` / `staging` / `prod`, เก็บ secret นอก repo (ตัวแปรแพลตฟอร์ม, vault)

### Tradeoff และ tech stack ที่ควรคิดล่วงหน้า

| หัวข้อ | ทางเลือกหลัก | ข้อควรรู้ |
|--------|----------------|-----------|
| DB | SQLite → PostgreSQL | ย้ายเมื่อต้องการ concurrent write, migration แบบมีเวอร์ชัน (เช่น golang-migrate) |
| Cache / session | ไม่มี → Redis | ใช้เมื่อ session ร่วมหลาย instance หรือ rate limit |
| Queue | ไม่มี → SQS / Rabbit / NATS | เมื่อมีงานหนักหลัง submit (อีเมล, รายงาน) ไม่ควรบล็อก request |
| Monolith vs แยกบริการ | คง Go monolith ก่อน | แยกเมื่อมีขอบเขตชัด (เช่น grading service) — อย่าแยกเร็วเกินเหตุ |
| Observability | log ธรรมดา → structured log + metrics + tracing | ช่วย debug บน prod เมื่อ user เยอะ |

สรุป: **คง Go + Gin + GORM + Vue เป็นฐาน** ได้นาน — ส่วนที่เปลี่ยนก่อนมักเป็น **ฐานข้อมูล, auth, และวิธี deploy** ไม่ใช่รีเขียนทั้งสแตก

## การขยายรองรับข้อสอบหลายชุด

- นิยาม **Exam** เป็น entity: `id`, `title`, `slug`, `version`
- **Backend:** เพิ่มตาราง `exams`, ผูก `questions.exam_id`; API เป็น `GET /api/exams/:id/questions`
- **Frontend:** route `/exam/:examId`, store แยก catalog vs session หรือ namespace ใน Pinia

## Backend ที่มีอยู่แล้ว (อ้างอิง)

| รายการ | รายละเอียด |
|--------|-------------|
| Entry | `backend/cmd/api/main.go` |
| DB | SQLite `backend/data/exam.db`, GORM `AutoMigrate` + seed เมื่อยังไม่มีข้อ |
| API | `GET /api/questions` — ข้อสอบไม่รวมเฉลย; `POST /api/submit` — `{ candidateName, answers }` → `{ candidateName, score, total }` และบันทึก `exam_results` |
| ชั้นโค้ด | `handler` → `usecase` → `repository` — ดู [architech.md](./architech.md) |

## การผสาน Frontend กับ Backend (สถานะปัจจุบัน)

- Dev: Vite proxy `/api` → `http://localhost:8080` — `examStore.loadQuestions()` / `submitExam()` เรียก API
- โหลดข้อสอบ: `GET /api/questions` — ล้มเหลวแสดงข้อความ error ไม่มีชุดข้อสอบในเครื่อง
- ส่งข้อสอบ: `POST /api/submit` — คะแนนจากเซิร์ฟเวอร์เท่านั้น
- Production: ตั้ง `VITE_API_BASE_URL` หรือ reverse proxy ร่วม host — ดู `frontend/.env.example`

## ประสบการณ์ผู้ใช้และความปลอดภัย

- Authentication ผู้สอบ (ถ้ามี) — มักเพิ่มที่ BE (JWT / session) แล้ว FE ส่ง token
- จำกัดเวลา, autosave ระหว่างทำข้อสอบ
- คะแนนที่เชื่อถือได้: ฝั่ง server เป็นตัวตัดสินเมื่อใช้ `POST /api/submit` แบบไม่เชื่อถือ client

## การทดสอบอัตโนมัติ

- **Backend:** `cd backend && go test ./...` — usecase + mock repository (มีแล้ว)
- **Frontend:** Vitest สำหรับ store/views เมื่อเพิ่ม

รายละเอียดบันทึกผลรันเทสดูที่ [testing.md](./testing.md)
