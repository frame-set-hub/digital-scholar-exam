# Backend (Go API)

REST API + SQLite — รายละเอียดสถาปัตยกรรม, ชั้น **handler → usecase → repository**, และ **ตาราง API แบบเต็ม** อยู่ที่ **[`../docs/`](../docs/)** — โดยเฉพาะ [`../docs/api.md`](../docs/api.md)

## Prerequisites

- Go 1.22+ (ดู `go.mod`)

## Commands

```bash
go run ./cmd/api      # default :8080, หรือตั้ง PORT
go test ./... -count=1 -v
```

## Layout

`cmd/api/main.go` · `internal/{handler,usecase,repository,models}` · `data/exam.db` (สร้างอัตโนมัติ)

---

[← กลับโปรเจกต์](../README.md) · [เอกสารทั้งระบบ](../docs/README.md) · [API reference](../docs/api.md)
