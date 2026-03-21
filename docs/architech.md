# Architecture & Tech Stack (Full Stack)

## สารบัญ

- [Architecture \& Tech Stack (Full Stack)](#architecture--tech-stack-full-stack)
  - [สารบัญ](#สารบัญ)
  - [ภาพรวม](#ภาพรวม)
  - [Tech Stack ฝั่ง Frontend](#tech-stack-ฝั่ง-frontend)
  - [Tech Stack ฝั่ง Backend](#tech-stack-ฝั่ง-backend)
  - [เหตุผลที่เลือก Vue 3 + Pinia](#เหตุผลที่เลือก-vue-3--pinia)
  - [เหตุผลที่เลือก Go + Gin + SQLite](#เหตุผลที่เลือก-go--gin--sqlite)
  - [โครงสร้างโฟลเดอร์ Frontend](#โครงสร้างโฟลเดอร์-frontend)
  - [โครงสร้างโฟลเดอร์ Backend (Pragmatic Clean Architecture)](#โครงสร้างโฟลเดอร์-backend-pragmatic-clean-architecture)
  - [การสื่อสารระหว่าง FE / BE](#การสื่อสารระหว่าง-fe--be)

## ภาพรวม

ระบบประกอบด้วย **SPA ฝั่ง Frontend** (Vue 3) สำหรับผู้สอบกรอกชื่อ ทำข้อสอบแบบเลือกคำตอบเดียว และดูผลคะแนน และ **API ฝั่ง Backend** (Go + Gin) ที่จัดเก็บข้อสอบ/เฉลยใน SQLite รับการส่งข้อสอบ คำนวณคะแนนที่ฝั่งเซิร์ฟเวอร์ และบันทึกผลการสอบ

ชั้น Frontend แยก UI (Vue), การนำทาง (Vue Router) และ state ชั่วคราว (Pinia) ชั้น Backend แยก HTTP (Handler), กฎธุรกิจ (Usecase), และการเข้าถึงข้อมูล (Repository + GORM)

## Tech Stack ฝั่ง Frontend

| เทคโนโลยี | บทบาท |
|-----------|--------|
| **Vue 3** | UI framework — Composition API + `<script setup>` |
| **Vite** | build และ dev server |
| **Tailwind CSS** | สไตล์ utility-first, responsive, mobile-first |
| **Vue Router** | เส้นทางหน้าทำข้อสอบ (IT 10-1) กับหน้าผล (IT 10-2) |
| **Pinia** | state ชื่อผู้สอบ, คำถาม, คำตอบ, คะแนน (ปัจจุบันยัง mock / คำนวณ client ได้จนกว่าจะผสาน API) |

## Tech Stack ฝั่ง Backend

| เทคโนโลยี | บทบาท |
|-----------|--------|
| **Go** | ภาษาและ runtime |
| **Gin** | HTTP router / middleware |
| **GORM** | ORM สำหรับ SQLite |
| **SQLite** | ฐานข้อมูลไฟล์เดียว (`backend/data/exam.db`) — zero extra install |
| **testify** | `assert` + `mock` สำหรับ unit test usecase |

## เหตุผลที่เลือก Vue 3 + Pinia

- **Vue 3** มี Composition API ที่จัดกลุ่ม logic ตาม feature ได้ชัด
- **Pinia** แยก state ของ **exam** ออกจากคอมโพเนนต์ ทำให้ `ExamView` / `ResultView` โฟกัสการแสดงผลและ event

## เหตุผลที่เลือก Go + Gin + SQLite

- **Go** deploy ง่าย binary เดียว, concurrency ชัดเจน
- **Gin** เป็นที่นิยมใน community, middleware ครบสำหรับ REST
- **SQLite** เหมาะกับโปรเจกต์เรียนรู้/สาธิต — ไม่ต้องติดตั้งเซิร์ฟเวอร์ DB แยก; ย้ายไป PostgreSQL ได้เมื่อต้องการ scale
- โครงสร้าง **Pragmatic Clean Architecture**: handler → usecase → repository — ทดสอบ usecase ด้วย mock repository ได้โดยไม่ต้องแตะ SQLite

## โครงสร้างโฟลเดอร์ Frontend

- `frontend/src/views/` — หน้าจอหลักตาม route
- `frontend/src/components/` — คอมโพเนนต์ย่อยที่ใช้ซ้ำ
- `frontend/src/stores/` — Pinia (`examStore`)
- `frontend/src/router/` — เส้นทางและ meta (title)
- `frontend/src/assets/` — CSS global และ Tailwind theme

## โครงสร้างโฟลเดอร์ Backend (Pragmatic Clean Architecture)

```
backend/
├── cmd/api/main.go          # entry, SQLite path, AutoMigrate, Seed, DI, Gin
├── internal/
│   ├── models/              # Question, Option, ExamResult
│   ├── repository/          # GORM: GetQuestions, SaveExamResult, migrate, seed
│   ├── usecase/             # Exam, ports (interfaces), ScoreAnswers
│   └── handler/             # Gin: GET /api/questions, POST /api/submit
├── go.mod
└── data/exam.db             # สร้างเมื่อรัน (อยู่ใน .gitignore)
```

- **Handler** รับ/ส่ง JSON ไม่มี business logic หนัก
- **Usecase** รวม `GetQuestions` (map เป็น DTO ไม่ส่งเฉลย), `SubmitExam` (ดึงเฉลยจาก DB → คำนวณคะแนน → `SaveExamResult`)
- **Repository** คุยกับ GORM/SQLite เท่านั้น

รายละเอียด endpoint และตัวอย่าง JSON: [api.md](./api.md)

## การสื่อสารระหว่าง FE / BE

สรุปสั้น: API ฐาน `http://localhost:8080` — ดูตารางและ payload ฉบับเต็มใน [api.md](./api.md)

รายละเอียด flow อยู่ใน [code_analyze.md](./code_analyze.md) และแผนผสานใน [planning.md](./planning.md)
