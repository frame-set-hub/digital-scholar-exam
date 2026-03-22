# Backend (Go API)

REST API + SQLite — architecture, **handler → use case → repository** layers, and the **full API table** live under **[`../docs/`](../docs/)**, especially [`../docs/api.md`](../docs/api.md)

## Prerequisites

- Go 1.22+ (see `go.mod`)

## Environment

1. Copy [`.env.example`](./.env.example) to **`.env`** in `backend/` (or from repo root: `cp backend/.env.example backend/.env`)
2. Optional keys: `PORT` (default `8080`), `DATABASE_DIR` (default `data/` under cwd). Values load from `.env` on startup via [`godotenv`](https://github.com/joho/godotenv); shell exports still win.

**`.env` is gitignored** — only `.env.example` is committed.

## Commands

```bash
go run ./cmd/api      # :8080 or PORT from .env / environment
go test ./... -count=1 -v
```

## Layout

`cmd/api/main.go` · `internal/{handler,usecase,repository,models}` · `scripts/` (optional dev SQL — run commands in [root README](../README.md#mock-data-for-testing-the-leaderboard)) · `data/exam.db` (created automatically)

---

[← Project root](../README.md) · [Full docs](../docs/README.md) · [API reference](../docs/api.md)
