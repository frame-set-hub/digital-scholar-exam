# digital-scholar-exam

ระบบสอบออนไลน์ — **Frontend:** Vue 3 · **Backend:** Go + Gin + SQLite

## แนวทางเอกสาร

| ชั้น | หน้าที่ | ไม่ควรทำ |
|------|---------|-----------|
| **[README.md](./README.md)** (ไฟล์นี้) | จุดเข้าเดียวของ repo: ภาพรวม, Quick start, แผนที่โฟลเดอร์, ลิงก์ไป `docs/` | คัดลอกเนื้อหายาวจาก `docs/` หรือทวนซ้ำตาราง API แบบเต็ม |
| **[docs/](./docs/)** | เอกสารเชิงลึก **แหล่งเดียว (canonical)** — สถาปัตยกรรม, flow, แผน, testing, **[API Reference](./docs/api.md)** | — |
| **`frontend/README.md`** / **`backend/README.md`** | คำสั่งรันแพ็กเกจนั้น + ลิงก์กลับ `docs/` — สำหรับคน `cd` เข้าโฟลเดอร์นั้น | ซ้ำสถาปัตยกรรมหรือตาราง API ยาวๆ |
| **[execute.md](./execute.md)** | Checklist Phase / ความคืบหน้า | — |
| **[RULE.md](./RULE.md)** | กฎการอัปเดต `.md` เมื่อการแก้โค้ดกระทบโครงสร้างหรือพฤติกรรมที่เอกสารอธิบาย | — |

สอดคล้องกับ monorepo ทั่วไป: **หน้าแรกของ repo = นำทาง**, รายละเอียดอยู่ใน `docs/`, แพ็กเกจย่อยมี README สั้นๆ

## Quick start

### ครั้งแรก: คัดลอกเทมเพลต `.env`

จาก **root ของ repo** (ค่าเริ่มใช้ dev ในเครื่องได้ — แก้เมื่อเปลี่ยนพอร์ตหรือ URL API):

```bash
cp frontend/.env.example frontend/.env
cp backend/.env.example backend/.env
```

จากนั้นเปิด **สองเทอร์มินัล** (หรือรัน backend ก่อน แล้วค่อย frontend)

**Backend** — โหลด `backend/.env` (optional) (`PORT`, `DATABASE_DIR`); ค่าเริ่มฟัง **:8080**:

```bash
cd backend && go run ./cmd/api
```

**Frontend** — อ่าน `frontend/.env` (`DEV_SERVER_PORT`, `API_PROXY_TARGET`, `VITE_API_BASE_URL`); URL dev เริ่มต้น **http://localhost:5173**, proxy `/api` → backend:

```bash
cd frontend && npm install && npm run dev
```

รัน **ทั้งสองพร้อมกัน** — แอปโหลดข้อสอบด้วย `GET /api/questions`, ส่งด้วย `POST /api/submit`, หน้า Leaderboard ดึง `GET /api/leaderboard` (ต้องรัน backend ให้พร้อม)

**ExamView:** ถ้ากด **Submit** ก่อนตอบครบทุกข้อ แอปแสดงคำเตือนใต้ปุ่ม ใส่กรอบแดงให้การ์ดทุกข้อที่ยังไม่ตอบ (รวมข้อความคำถาม) และเลื่อนแบบ smooth ไปข้อแรกที่ยังว่าง

<a id="mock-data-for-testing-the-leaderboard"></a>

### ข้อมูล mock สำหรับทดสอบ Leaderboard

ฐานข้อมูล SQLite อยู่ที่ `backend/data/exam.db` (สร้างเมื่อรัน API — อยู่ใน `.gitignore`)

**หยุด backend ก่อน** (กันติดล็อกไฟล์ DB) รันคำสั่งต่อไปนี้ **จาก root ของ repo** เพื่อให้ path ตรงกัน

**แทรกแถวตัวอย่าง** ใน `exam_results` — [`backend/scripts/mock_exam_results.sql`](./backend/scripts/mock_exam_results.sql)

```bash
sqlite3 backend/data/exam.db < backend/scripts/mock_exam_results.sql
```

**ลบผลการสอบทั้งหมด** (ล้าง Leaderboard / รีเซ็ตก่อนลองใหม่) — [`backend/scripts/clear_exam_results.sql`](./backend/scripts/clear_exam_results.sql)

```bash
sqlite3 backend/data/exam.db < backend/scripts/clear_exam_results.sql
```

รัน mock ซ้ำจะ INSERT เพิ่มเรื่อยๆ — ถ้าอยากเริ่มว่างก่อนแทรก mock ให้รัน **clear** ก่อน

**ทำไมใช้ `backend/scripts/` ไม่ใช่ `cmd/`?** ในโปรเจกต์ Go [`cmd/`](./backend/cmd/) ใส่ **จุดเริ่ม executable** (`package main`, เช่น `cmd/api`) SQL สำหรับ dev / ดูแล DB ไม่ใช่ binary — เก็บใน `backend/scripts/` ตรงกับ layout ทั่วไปและแยกจากการประกอบแอป (handler → use case → repository)

## เอกสาร (อ่านต่อ)

| หัวข้อ | ลิงก์ |
|--------|------|
| ดัชนีเอกสารทั้งหมด | [docs/README.md](./docs/README.md) |
| API (endpoint + JSON) | [docs/api.md](./docs/api.md) |
| สถาปัตยกรรม & stack | [docs/architech.md](./docs/architech.md) |
| Flow + diagram | [docs/architech.md](./docs/architech.md) |
| อ่านโค้ดทีละไฟล์ (ตามบรรทัด) | [docs/code_analyze.md](./docs/code_analyze.md) |
| แผนอนาคต & roadmap | [docs/planning.md](./docs/planning.md) |
| Testing | [docs/testing.md](./docs/testing.md) |
| Progress / Phase | [execute.md](./execute.md) |
| กฎซิงค์เอกสารกับโค้ด | [RULE.md](./RULE.md) |

## โครงสร้าง repo

| โฟลเดอร์ | คำอธิบาย |
|----------|----------|
| [`frontend/`](./frontend/) | Vue 3 + Vite + Pinia + Tailwind |
| [`backend/`](./backend/) | Go + Gin + GORM (`cmd/api`, `internal/`) · สคริปต์ SQL ทดสอบ: [`mock_exam_results.sql`](./backend/scripts/mock_exam_results.sql), [`clear_exam_results.sql`](./backend/scripts/clear_exam_results.sql) |
| [`docs/`](./docs/) | เอกสารออกแบบและ API reference |
