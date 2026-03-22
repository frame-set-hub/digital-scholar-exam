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

### One-time: copy env templates

From the **repository root** (defaults work for local dev — edit only if you change ports or API URLs):

```bash
cp frontend/.env.example frontend/.env
cp backend/.env.example backend/.env
```

Then open **two terminals** (or run the backend first, then the frontend).

**Backend** — loads optional `backend/.env` (`PORT`, `DATABASE_DIR`); default listen **:8080**:

```bash
cd backend && go run ./cmd/api
```

**Frontend** — reads `frontend/.env` (`DEV_SERVER_PORT`, `API_PROXY_TARGET`, `VITE_API_BASE_URL`); default dev URL **http://localhost:5173**, proxies `/api` → backend:

```bash
cd frontend && npm install && npm run dev
```

Run **both** — the app loads questions via `GET /api/questions`, submits with `POST /api/submit`, and the Leaderboard page fetches `GET /api/leaderboard` (the backend must be running).

**ExamView:** If you click **Submit** before answering every question, the app shows a warning under the button, adds a red border to each unanswered card (including the prompt text), and smooth-scrolls to the first unanswered question.

### Mock data for testing the Leaderboard

The SQLite database lives at `backend/data/exam.db` (created when the API runs — listed in `.gitignore`).

**Stop the backend first** (to avoid DB file locking). Run the following **from the repository root** so paths match.

**Insert sample rows** into `exam_results` — [`backend/scripts/mock_exam_results.sql`](./backend/scripts/mock_exam_results.sql)

```bash
sqlite3 backend/data/exam.db < backend/scripts/mock_exam_results.sql
```

**Delete all exam results** (clear leaderboard / reset before retrying) — [`backend/scripts/clear_exam_results.sql`](./backend/scripts/clear_exam_results.sql)

```bash
sqlite3 backend/data/exam.db < backend/scripts/clear_exam_results.sql
```

Running mock repeatedly will keep INSERTing rows — to start fresh before inserting mock data, run **clear** first.

**Why `backend/scripts/` and not `cmd/`?** In Go projects, [`cmd/`](./backend/cmd/) is for **executable entrypoints** (`package main`, e.g. `cmd/api`). SQL used for local dev / DB maintenance is not a binary — keeping it in `backend/scripts/` matches common layout and stays separate from application composition (handler → use case → repository).

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
