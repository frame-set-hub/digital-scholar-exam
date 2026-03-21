# digital-scholar-exam

Online exam system — **Frontend:** Vue 3 · **Backend:** Go + Gin + SQLite

## Documentation layout

| Layer | Purpose | Avoid |
|------|---------|--------|
| **[README.md](./README.md)** (this file) | Single entry to the repo: overview, quick start, folder map, links to `docs/` | Pasting long content from `docs/` or duplicating full API tables |
| **[docs/](./docs/)** | In-depth **canonical** docs — architecture, flows, planning, testing, **[API Reference](./docs/api.md)** | — |
| **`frontend/README.md`** / **`backend/README.md`** | Run commands for that package + links back to `docs/` — for people who `cd` into that folder | Duplicating architecture or long API tables |
| **[execute.md](./execute.md)** | Phase / progress checklist | — |
| **[RULE.md](./RULE.md)** | Rules for updating `.md` when code changes affect structure or behavior described in docs | — |

This matches common monorepo practice: **repo home = navigation**, details live in `docs/`, subpackages have short READMEs.

## Quick start

Open **two terminals** (or run the backend first, then the frontend)

**Backend** (port 8080):

```bash
cd backend && go run ./cmd/api
```

**Frontend** (port 5173 — proxies `/api` to the backend per `API_PROXY_TARGET` in `frontend/.env`):

```bash
cd frontend
cp .env.example .env   # first time — .env is gitignored
npm install && npm run dev
```

<<<<<<< HEAD
รัน **ทั้งสองพร้อมกัน** — หน้าเว็บโหลดข้อสอบจาก `GET /api/questions` ส่งคำตอบด้วย `POST /api/submit` และหน้า Leaderboard ดึง `GET /api/leaderboard` (ต้องรัน backend ให้พร้อม)

### ข้อมูล mock สำหรับทดสอบ Leaderboard

ฐานข้อมูล SQLite อยู่ที่ `backend/data/exam.db` (สร้างตอนรัน API — อยู่ใน `.gitignore`)

**หยุด backend ก่อน** (กันติดล็อกไฟล์ DB) แล้วใช้สคริปต์ใดสคริปต์หนึ่งจาก root ของ repo:

| การทำ | ไฟล์ | คำสั่ง |
|--------|------|--------|
| แทรกแถวตัวอย่างใน `exam_results` | [`backend/scripts/mock_exam_results.sql`](./backend/scripts/mock_exam_results.sql) | `sqlite3 backend/data/exam.db < backend/scripts/mock_exam_results.sql` |
| **ลบผลการสอบทั้งหมด** (ล้าง Leaderboard / รีเซ็ตก่อนลองใหม่) | [`backend/scripts/clear_exam_results.sql`](./backend/scripts/clear_exam_results.sql) | `sqlite3 backend/data/exam.db < backend/scripts/clear_exam_results.sql` |

รัน mock ซ้ำจะ INSERT เพิ่มเรื่อยๆ — ถ้าอยากเริ่มว่างก่อนแทรก mock ให้รัน **clear** ก่อน
=======
Run **both** — the app loads questions via `GET /api/questions` and submits with `POST /api/submit` only (the backend must be running).
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

## Documentation (read more)

| Topic | Link |
|--------|------|
| Full documentation index | [docs/README.md](./docs/README.md) |
| API (endpoints + JSON) | [docs/api.md](./docs/api.md) |
| Architecture & stack | [docs/architech.md](./docs/architech.md) |
| Flow + diagrams | [docs/architech.md](./docs/architech.md) |
| Code walkthrough (by line) | [docs/code_analyze.md](./docs/code_analyze.md) |
| Future plans & roadmap | [docs/planning.md](./docs/planning.md) |
| Testing | [docs/testing.md](./docs/testing.md) |
| Progress / phases | [execute.md](./execute.md) |
| Doc–code sync rules | [RULE.md](./RULE.md) |

## Repository layout

| Folder | Description |
|----------|----------|
| [`frontend/`](./frontend/) | Vue 3 + Vite + Pinia + Tailwind |
<<<<<<< HEAD
| [`backend/`](./backend/) | Go + Gin + GORM (`cmd/api`, `internal/`) · SQL ช่วยทดสอบ: [`mock_exam_results.sql`](./backend/scripts/mock_exam_results.sql), [`clear_exam_results.sql`](./backend/scripts/clear_exam_results.sql) |
| [`docs/`](./docs/) | เอกสารออกแบบและ API reference |
=======
| [`backend/`](./backend/) | Go + Gin + GORM (`cmd/api`, `internal/`) |
| [`docs/`](./docs/) | Design docs and API reference |
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)
