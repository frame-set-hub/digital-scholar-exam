# Backend (Go API)

REST API + SQLite — architecture, **handler → use case → repository** layers, and the **full API table** live under **[`../docs/`](../docs/)**, especially [`../docs/api.md`](../docs/api.md)

## Prerequisites

- Go 1.22+ (see `go.mod`)

## Commands

```bash
go run ./cmd/api      # default :8080, or set PORT
go test ./... -count=1 -v
```

## Layout

`cmd/api/main.go` · `internal/{handler,usecase,repository,models}` · `scripts/` (optional dev SQL — run commands in [root README](../README.md#mock-data-for-testing-the-leaderboard)) · `data/exam.db` (created automatically)

---

[← Project root](../README.md) · [Full docs](../docs/README.md) · [API reference](../docs/api.md)
