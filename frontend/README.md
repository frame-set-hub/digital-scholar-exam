# Frontend (Vue 3)

SPA สำหรับทำข้อสอบ — รายละเอียดสถาปัตยกรรมและ flow อยู่ที่ **[`../docs/`](../docs/)** (ไม่ซ้ำในไฟล์นี้)

## Prerequisites

- [Node.js](https://nodejs.org/) LTS

## Environment

1. คัดลอก [`.env.example`](./.env.example) เป็น **`.env`** ที่ root ของโฟลเดอร์ `frontend/`
2. ปรับ `VITE_API_BASE_URL` (ว่าง = ใช้ `/api` + proxy) และ `API_PROXY_TARGET` (ที่ dev proxy ส่งต่อไป)

ไฟล์ **`.env` ถูก gitignore** — ไม่ขึ้น commit; มีเฉพาะ `.env.example` เป็นต้นแบบ

## Commands

```bash
npm install
npm run dev          # http://localhost:5173 — proxy /api → http://localhost:8080 (ต้องรัน backend ด้วย)
npm run build
npm run preview      # ถ้าเทส build กับ API ให้ตั้ง VITE_API_BASE_URL ใน .env.production
```

ดู [`.env.example`](./.env.example) — ตอน `vite build` ถ้าไม่ใช้ reverse proxy ร่วม host ให้ตั้ง `VITE_API_BASE_URL` ชี้ backend

## Layout

`src/views/` · `src/stores/` · `src/router/` · `src/assets/`

---

[← กลับโปรเจกต์](../README.md) · [เอกสารทั้งระบบ](../docs/README.md) · [API](../docs/api.md)
