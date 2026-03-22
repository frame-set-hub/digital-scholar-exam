# Frontend (Vue 3)

SPA for taking exams — architecture and flows are in **[`../docs/`](../docs/)** (not duplicated here)

## Prerequisites

- [Node.js](https://nodejs.org/) LTS

## Environment

1. Copy [`.env.example`](./.env.example) to **`.env`** at the `frontend/` folder root (or from repo root: `cp frontend/.env.example frontend/.env`)
2. Adjust `DEV_SERVER_PORT` (Vite dev server), `VITE_API_BASE_URL` (empty = use `/api` + proxy), and `API_PROXY_TARGET` (where the dev proxy forwards — match the backend `PORT`)

**`.env` is gitignored** — not committed; only `.env.example` is the template

## Commands

```bash
npm install
npm run dev          # port from DEV_SERVER_PORT in .env (default 5173); proxy /api → API_PROXY_TARGET
npm run build
npm run preview      # to test build against API, set VITE_API_BASE_URL in .env.production
```

See [`.env.example`](./.env.example) — for `vite build` without a same-host reverse proxy, set `VITE_API_BASE_URL` to the backend

## Layout

`src/views/` · `src/stores/` · `src/router/` · `src/assets/`

---

[← Project root](../README.md) · [Full docs](../docs/README.md) · [API](../docs/api.md)
