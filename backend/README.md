# Backend (Go API)

REST API + SQLite — สถาปัตยกรรม, ชั้น **handler → use case → repository** และ **ตาราง API แบบเต็ม** อยู่ที่ **[`../docs/`](../docs/)** โดยเฉพาะ [`../docs/api.md`](../docs/api.md)

## Prerequisites

- Go 1.22+ (ดู `go.mod`)

## Environment

1. คัดลอก [`.env.example`](./.env.example) เป็น **`.env`** ใน `backend/` (หรือจาก root: `cp backend/.env.example backend/.env`)
2. คีย์ optional: `PORT` (ค่าเริ่ม `8080`), `DATABASE_DIR` (ค่าเริ่ม `data/` ภายใต้ cwd) โหลดจาก `.env` ตอนเริ่มด้วย [`godotenv`](https://github.com/joho/godotenv); ค่าที่ export ใน shell ยังชนะ

**`.env` ถูก gitignore** — commit เฉพาะ `.env.example`

## Commands

```bash
go run ./cmd/api      # :8080 หรือ PORT จาก .env / environment
go test ./... -count=1 -v
```

## Layout

`cmd/api/main.go` · `internal/{handler,usecase,repository,models}` · `scripts/` (SQL dev optional — คำสั่งรันใน [README ที่ root](../README.md#mock-data-for-testing-the-leaderboard)) · `data/exam.db` (สร้างอัตโนมัติ)

---

[← โปรเจกต์](../README.md) · [เอกสารทั้งหมด](../docs/README.md) · [API reference](../docs/api.md)
