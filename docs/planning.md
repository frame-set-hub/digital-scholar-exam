# Planning & Future Design (FE + BE)

## สารบัญ

- [Planning \& Future Design (FE + BE)](#planning--future-design-fe--be)
  - [สารบัญ](#สารบัญ)
  - [สถานะปัจจุบัน](#สถานะปัจจุบัน)
  - [เป้าหมายระยะสั้น](#เป้าหมายระยะสั้น)
  - [การขยายรองรับข้อสอบหลายชุด](#การขยายรองรับข้อสอบหลายชุด)
  - [Backend ที่มีอยู่แล้ว (อ้างอิง)](#backend-ที่มีอยู่แล้ว-อ้างอิง)
  - [Phase 6 — ผสาน Frontend กับ Backend](#phase-6--ผสาน-frontend-กับ-backend)
  - [ประสบการณ์ผู้ใช้และความปลอดภัย](#ประสบการณ์ผู้ใช้และความปลอดภัย)
  - [การทดสอบอัตโนมัติ](#การทดสอบอัตโนมัติ)

## สถานะปัจจุบัน

- **Frontend:** flow ครบด้วย mock ใน Pinia (หรือข้อมูลเดิมจาก store)
- **Backend:** มี API `GET /api/questions`, `POST /api/submit`, SQLite + seed อัตโนมัติ — ดู [`api.md`](./api.md)

## เป้าหมายระยะสั้น

- คง FE ให้ทำข้อสอบและแสดงผลได้แม้ยังไม่ต่อ API
- BE พร้อมเป็นข้อมูลจริงและเป็นผู้ตัดสินคะแนนเมื่อผสานแล้ว

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

## Phase 6 — ผสาน Frontend กับ Backend

- ตั้งค่า base URL (เช่น `VITE_API_URL=http://localhost:8080`) และ `fetch` / `axios` ใน store หรือ service
- โหลดข้อสอบ: `GET /api/questions` → แทนที่ mock ใน `examStore.questions`
- ส่งข้อสอบ: `POST /api/submit` → ใช้ `score` / `total` จาก response แทน `computeScore()` ใน client
- จัดการ error / loading state และ CORS ใน production (จำกัด origin แทน `*`)

## ประสบการณ์ผู้ใช้และความปลอดภัย

- Authentication ผู้สอบ (ถ้ามี) — มักเพิ่มที่ BE (JWT / session) แล้ว FE ส่ง token
- จำกัดเวลา, autosave ระหว่างทำข้อสอบ
- คะแนนที่เชื่อถือได้: ฝั่ง server เป็นตัวตัดสินเมื่อใช้ `POST /api/submit` แบบไม่เชื่อถือ client

## การทดสอบอัตโนมัติ

- **Backend:** `cd backend && go test ./...` — usecase + mock repository (มีแล้ว)
- **Frontend:** Vitest สำหรับ store/views เมื่อเพิ่ม

รายละเอียดบันทึกผลรันเทสดูที่ [testing.md](./testing.md)
