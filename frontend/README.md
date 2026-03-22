# Frontend (Vue 3)

SPA สำหรับทำข้อสอบ — สถาปัตยกรรมและ flow อยู่ที่ **[`../docs/`](../docs/)** (ไม่ซ้ำในไฟล์นี้)

## Prerequisites

- [Node.js](https://nodejs.org/) LTS

## Environment

1. คัดลอก [`.env.example`](./.env.example) เป็น **`.env`** ที่ root ของโฟลเดอร์ `frontend/` (หรือจาก root: `cp frontend/.env.example frontend/.env`)
2. ปรับ `DEV_SERVER_PORT` (เซิร์ฟเวอร์ dev ของ Vite), `VITE_API_BASE_URL` (ว่าง = ใช้ `/api` + proxy), และ `API_PROXY_TARGET` (ที่ dev proxy ส่งต่อ — ให้ตรงกับ `PORT` ของ backend)

**`.env` ถูก gitignore** — ไม่ขึ้น commit; มีเฉพาะ `.env.example` เป็นต้นแบบ

## Commands

```bash
npm install
npm run dev          # พอร์ตจาก DEV_SERVER_PORT ใน .env (ค่าเริ่ม 5173); proxy /api → API_PROXY_TARGET
npm run build
npm run preview      # ถ้าทดสอบ build กับ API ให้ตั้ง VITE_API_BASE_URL ใน .env.production
```

ดู [`.env.example`](./.env.example) — ตอน `vite build` ถ้าไม่ใช้ reverse proxy ร่วม host ให้ตั้ง `VITE_API_BASE_URL` ชี้ backend

## Layout

`src/views/` · `src/stores/` · `src/router/` · `src/assets/`

---

[← โปรเจกต์](../README.md) · [เอกสารทั้งหมด](../docs/README.md) · [API](../docs/api.md)
