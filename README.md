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

Run **both** — the app loads questions via `GET /api/questions`, submits with `POST /api/submit`, and the Leaderboard page fetches `GET /api/leaderboard` (the backend must be running).

### Mock data for testing the Leaderboard

The SQLite database lives at `backend/data/exam.db` (created when the API runs — listed in `.gitignore`).

**Stop the backend first** (to avoid DB file locking) then use one of the scripts from the repo root:

| Action | File | Command |
|--------|------|---------|
| Insert sample rows into `exam_results` | [`backend/scripts/mock_exam_results.sql`](./backend/scripts/mock_exam_results.sql) | `sqlite3 backend/data/exam.db < backend/scripts/mock_exam_results.sql` |
| **Delete all exam results** (clear leaderboard / reset before retrying) | [`backend/scripts/clear_exam_results.sql`](./backend/scripts/clear_exam_results.sql) | `sqlite3 backend/data/exam.db < backend/scripts/clear_exam_results.sql` |

Running mock repeatedly will keep INSERTing rows — to start fresh before inserting mock data, run **clear** first.

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
| [`backend/`](./backend/) | Go + Gin + GORM (`cmd/api`, `internal/`) · Test SQL scripts: [`mock_exam_results.sql`](./backend/scripts/mock_exam_results.sql), [`clear_exam_results.sql`](./backend/scripts/clear_exam_results.sql) |
| [`docs/`](./docs/) | Design docs and API reference |
