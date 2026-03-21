# digital-scholar-exam

ระบบทำแบบทดสอบออนไลน์ — **Frontend:** Vue 3 · **Backend:** Go + Gin + SQLite

## แนวทางเอกสาร 

| ชั้น | หน้าที่ | ไม่ควรทำ |
|------|---------|-----------|
| **[README.md](./README.md)** (ไฟล์นี้) | จุดเข้าเดียวของ repo: ภาพรวม, Quick start, แผนที่โฟลเดอร์, ลิงก์ไป `docs/` | คัดลอกเนื้อหายาวจาก `docs/` หรือทวนซ้ำ API แบบเต็ม |
| **[docs/](./docs/)** | เอกสารเชิงลึก **แหล่งเดียว (canonical)** — สถาปัตยกรรม, flow, แผน, testing, **[API Reference](./docs/api.md)** | — |
| **`frontend/README.md`** / **`backend/README.md`** | คำสั่งรันแพ็กเกจ + ลิงก์กลับ `docs/` — สำหรับคน `cd` เข้าโฟลเดอร์นั้น | ซ้ำสถาปัตยกรรมหรือตาราง API แบบยาว |
| **[execute.md](./execute.md)** | Checklist Phase / progress งาน | — |
| **[RULE.md](./RULE.md)** | กฎอัปเดตเอกสาร `.md` เมื่อการแก้โค้ดกระทบโครงสร้างหรือเนื้อหาที่เอกสารอธิบาย | — |

แนวทางนี้สอดคล้องกับโปรเจกต์ monorepo ทั่วไป: **หน้าแรกของ repo = นำทาง**, รายละเอียดอยู่ใน `docs/`, แพ็กเกจย่อยมี README บางๆ

## Quick start

เปิด **สองเทอร์มินัล** (หรือรัน backend ก่อน แล้วค่อย frontend)

**Backend** (port 8080):

```bash
cd backend && go run ./cmd/api
```

**Frontend** (port 5173 — proxy `/api` ไป backend ตาม `API_PROXY_TARGET` ใน `frontend/.env`):

```bash
cd frontend
cp .env.example .env   # ครั้งแรก — .env ถูก gitignore
npm install && npm run dev
```

รัน **ทั้งสองพร้อมกัน** — หน้าเว็บโหลดข้อสอบจาก `GET /api/questions` ส่งคำตอบด้วย `POST /api/submit` และหน้า Leaderboard ดึง `GET /api/leaderboard` (ต้องรัน backend ให้พร้อม)

### ข้อมูล mock สำหรับทดสอบ Leaderboard

ฐานข้อมูล SQLite อยู่ที่ `backend/data/exam.db` (สร้างตอนรัน API — อยู่ใน `.gitignore`)

**หยุด backend ก่อน** (กันติดล็อกไฟล์ DB) แล้วใช้สคริปต์ใดสคริปต์หนึ่งจาก root ของ repo:

| การทำ | ไฟล์ | คำสั่ง |
|--------|------|--------|
| แทรกแถวตัวอย่างใน `exam_results` | [`backend/scripts/mock_exam_results.sql`](./backend/scripts/mock_exam_results.sql) | `sqlite3 backend/data/exam.db < backend/scripts/mock_exam_results.sql` |
| **ลบผลการสอบทั้งหมด** (ล้าง Leaderboard / รีเซ็ตก่อนลองใหม่) | [`backend/scripts/clear_exam_results.sql`](./backend/scripts/clear_exam_results.sql) | `sqlite3 backend/data/exam.db < backend/scripts/clear_exam_results.sql` |

รัน mock ซ้ำจะ INSERT เพิ่มเรื่อยๆ — ถ้าอยากเริ่มว่างก่อนแทรก mock ให้รัน **clear** ก่อน

## เอกสาร (อ่านต่อ)

| หัวข้อ | ลิงก์ |
|--------|------|
| ดัชนีเอกสารทั้งหมด | [docs/README.md](./docs/README.md) |
| API (endpoint + JSON) | [docs/api.md](./docs/api.md) |
| สถาปัตยกรรม & stack | [docs/architech.md](./docs/architech.md) |
| Flow + diagram | [docs/architech.md](./docs/architech.md) |
| อ่านโค้ดทีละไฟล์ (บรรทัด) | [docs/code_analyze.md](./docs/code_analyze.md) |
| แผนอนาคต & roadmap | [docs/planning.md](./docs/planning.md) |
| Testing | [docs/testing.md](./docs/testing.md) |
| Progress / Phase | [execute.md](./execute.md) |
| กฎซิงค์เอกสารกับโค้ด | [RULE.md](./RULE.md) |

## โครงสร้าง repo

| โฟลเดอร์ | คำอธิบาย |
|----------|----------|
| [`frontend/`](./frontend/) | Vue 3 + Vite + Pinia + Tailwind |
| [`backend/`](./backend/) | Go + Gin + GORM (`cmd/api`, `internal/`) · SQL ช่วยทดสอบ: [`mock_exam_results.sql`](./backend/scripts/mock_exam_results.sql), [`clear_exam_results.sql`](./backend/scripts/clear_exam_results.sql) |
| [`docs/`](./docs/) | เอกสารออกแบบและ API reference |
