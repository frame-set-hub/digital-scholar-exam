# Planning & Future Design (FE + BE)

## Table of contents

- [Planning \& Future Design (FE + BE)](#planning--future-design-fe--be)
  - [Table of contents](#table-of-contents)
  - [Current state](#current-state)
  - [Short-term goals](#short-term-goals)
    - [Leaderboard: your rank when outside top 20](#leaderboard-your-rank-when-outside-top-20)
  - [Long-term goals](#long-term-goals)
  - [Supporting multiple exams](#supporting-multiple-exams)
  - [Existing backend (reference)](#existing-backend-reference)
  - [Frontend and backend integration (current)](#frontend-and-backend-integration-current)
  - [UX and security](#ux-and-security)
  - [Automated testing](#automated-testing)

## Current state

- **Frontend:** Full flow — Pinia loads questions from API (`GET /api/questions`) only
- **Backend:** APIs `GET /api/questions`, `POST /api/submit`, SQLite + automatic seed — see [`api.md`](./api.md)
- **Phase 7:** Added Leaderboard system using data from the `ExamResult` table, sorted by score (`Score DESC`) and time (`CreatedAt ASC`) to display candidate rankings — API `GET /api/leaderboard`, page `/leaderboard` in Vue

## Short-term goals

- Frontend needs a running backend (or proxy to an API) to load questions and submit answers
- Backend is the source of truth and scores submissions after integration

### Leaderboard: your rank when outside top 20

**Problem:** `GET /api/leaderboard` returns at most **20** entries (see [`api.md`](./api.md)). The UI lists only those rows — if a candidate’s rank is **21+**, they never see their own row on `/leaderboard` (e.g. via Vite dev `http://localhost:5173/api/leaderboard` proxied to the backend).

**Direction (to implement later):**

- **API:** e.g. optional `?name=` or `GET /api/leaderboard/me` after submit (needs stable identity — today only `candidateName` string) to return `{ rank, candidateName, score, total }` or merge a “your row” into the list response
- **FE:** `LeaderboardView` — show top 20 plus a **“Your position”** strip (rank, score) when the user is not in the top list (requires passing name from result flow or session)

Tracked in project progress: [`../execute.md`](../execute.md) (Backlog).

## Long-term goals

This document sets direction when the system grows beyond "single machine / single user" — not an immediate implementation spec, but helps decide stack and sequencing later.

### Scale and load

- **Horizontal:** Keep the API **stateless** (sessions not tied to one machine's memory) so multiple instances can sit behind a load balancer
- **Database:** SQLite fits dev/demo — for high concurrent writes or backup/HA, target **PostgreSQL** (or MySQL); consider read replicas if read-heavy
- **Bottleneck:** Usually DB and query/transaction design, not Go/Vue themselves

### Multiple users (and roles)

- **Many candidates:** Need **identity** — registration/login or SSO — not only a string in `candidateName`
- **Roles:** Separate **exam administrators** from **candidates**; may extend to orgs/classes (multi-tenant) later
- **Tradeoff:** Backend auth (JWT / session cookie + HTTPS) is clearer than trusting client tokens alone

### Cloud / production deploy

- **Frontend:** Static Vite build to **CDN / object storage** (e.g. S3 + CloudFront) or a **Pages** product; set `VITE_API_BASE_URL` to the real API
- **Backend:** Go binary in a **container** (Docker) on **ECS, Cloud Run, Fly.io, Railway**, etc. — choose by budget, team, and familiarity
- **Environments:** Separate `dev` / `staging` / `prod`, secrets outside the repo (platform env, vault)

### Tradeoffs and stack choices

| Topic | Main options | Notes |
|--------|----------------|-----------|
| DB | SQLite → PostgreSQL | Migrate when concurrent writes matter; versioned migrations (e.g. golang-migrate) |
| Cache / session | None → Redis | When sessions span instances or for rate limiting |
| Queue | None → SQS / Rabbit / NATS | Heavy post-submit work (email, reports) should not block the request |
| Monolith vs services | Keep Go monolith first | Split when boundaries are clear (e.g. grading service) — avoid premature split |
| Observability | Plain logs → structured logs + metrics + tracing | Helps debug production at scale |

Summary: **Go + Gin + GORM + Vue** can stay the base for a long — what changes first is usually **database, auth, and deploy**, not a full rewrite.

## Supporting multiple exams

- Define **Exam** as an entity: `id`, `title`, `slug`, `version`
- **Backend:** Add `exams` table, `questions.exam_id`; API e.g. `GET /api/exams/:id/questions`
- **Frontend:** route `/exam/:examId`, store separates catalog vs session or namespaces Pinia

## Existing backend (reference)

| Item | Detail |
|--------|-------------|
| Entry | `backend/cmd/api/main.go` |
| DB | SQLite `backend/data/exam.db`, GORM `AutoMigrate` + seed when empty |
| API | `GET /api/questions` — no answers in response; `POST /api/submit` — `{ candidateName, answers }` → `{ candidateName, score, total }` and persist `exam_results`; `GET /api/leaderboard` — ranked candidates (no raw answers) |
| Layers | `handler` → `usecase` → `repository` — see [architech.md](./architech.md) |

## Frontend and backend integration (current)

- Dev: Vite proxy `/api` → `http://localhost:8080` — `examStore.loadQuestions()` / `submitExam()` / `loadLeaderboard()` call the API
- Load: `GET /api/questions` — on failure show error; no local question set
- Submit: `POST /api/submit` — score from server only
- Leaderboard: `GET /api/leaderboard` — `LeaderboardView` calls `loadLeaderboard()` displaying entries from `exam_results` (no raw answers)
- Production: set `VITE_API_BASE_URL` or same-host reverse proxy — see `frontend/.env.example`

## UX and security

- Candidate authentication (if added) — usually on BE (JWT / session), FE sends token
- Time limits, autosave during exam
- Trusted scoring: server is authoritative with `POST /api/submit` (do not trust the client alone)

## Automated testing

- **Backend:** `cd backend && go test ./...` — use case + mock repository (already present)
- **Frontend:** Vitest for store/views when added

Run logs and details: [testing.md](./testing.md)
