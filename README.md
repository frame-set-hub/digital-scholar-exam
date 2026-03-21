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

รัน **ทั้งสองพร้อมกัน** — หน้าเว็บโหลดข้อสอบจาก `GET /api/questions` และส่งคำตอบด้วย `POST /api/submit` เท่านั้น (ต้องรัน backend ให้พร้อม)

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
| [`backend/`](./backend/) | Go + Gin + GORM (`cmd/api`, `internal/`) |
| [`docs/`](./docs/) | เอกสารออกแบบและ API reference |
