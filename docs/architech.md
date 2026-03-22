# Architecture & Tech Stack (Full Stack)

## Table of contents

- [Architecture \& Tech Stack (Full Stack)](#architecture--tech-stack-full-stack)
  - [Table of contents](#table-of-contents)
  - [Overview](#overview)
  - [Application flow (user → system)](#application-flow-user--system)
  - [Backend use case flow](#backend-use-case-flow)
  - [Data flow](#data-flow)
  - [API contract summary (`api.md`)](#api-contract-summary-apimd)
  - [Diagram — Frontend relationships](#diagram--frontend-relationships)
  - [Diagram — Backend request sequence](#diagram--backend-request-sequence)
  - [Frontend tech stack](#frontend-tech-stack)
  - [Backend tech stack](#backend-tech-stack)
  - [Why Vue 3 + Pinia](#why-vue-3--pinia)
  - [Why Go + Gin + SQLite](#why-go--gin--sqlite)
  - [Frontend folder layout](#frontend-folder-layout)
  - [Backend folder layout (Pragmatic Clean Architecture)](#backend-folder-layout-pragmatic-clean-architecture)
  - [Frontend and backend communication](#frontend-and-backend-communication)

## Overview

The system consists of a **Frontend SPA** (Vue 3) where candidates enter a name, take a single-choice exam, and view their score, and a **Backend API** (Go + Gin) that stores questions/answers in SQLite, accepts submissions, computes scores on the server, and persists exam results.

The frontend separates UI (Vue), routing (Vue Router), and transient state (Pinia). The backend separates HTTP (Handler), business rules (Use case), and data access (Repository + GORM).

**File-by-file reading with line numbers:** [code_analyze.md](./code_analyze.md)  
**Endpoints and JSON examples:** [api.md](./api.md)

## Application flow (user → system)

1. User opens the site → Vite loads the bundle from `main.js` → `App.vue` → `RouterView` for the path
2. Path `/` loads `ExamView` → `onMounted` calls **`GET /api/questions`** (via `examStore.loadQuestions()`)
3. **Success:** questions stored in Pinia  
   **Failure:** clear `questions`, set `loadError`, show a message — no offline question set
4. User enters name and selects answers → `setAnswer` updates `answers`
5. Submit → validate name and all questions answered → **`POST /api/submit`** with `{ candidateName, answers }` → receive `score` from server → navigate to `/result`
6. `ResultView` shows name and score → **View Leaderboard** → `/leaderboard` → `LeaderboardView` calls **`GET /api/leaderboard`** (`loadLeaderboard`) or **Retake** → `resetExam()`
7. `LeaderboardView` button **Back to Exam** → `resetExam()` (clears name/answers/score/leaderboard, returns to `/` — does not clear questions to avoid redundant `GET /api/questions`)

**DevTools / duplicate API calls:** see [api.md](./api.md)

## Backend use case flow

| Step | Owner | What happens |
|------|----------------|------------------|
| HTTP | `handler.ExamHTTP` | Accept request, bind JSON, HTTP status |
| Business rules | `usecase.Exam` | `GetQuestions`: load from store → map to DTO **without answers** |
| | | `SubmitExam`: load questions with answers from DB → `ScoreAnswers` → build `ExamResult` (including answer JSON) → `SaveExamResult` |
| | | `GetLeaderboard`: load `ExamResult` from store → map to `LeaderboardEntryDTO` (does not include `answers`) |
| Data | `repository.QuestionGorm` / `ExamResultGorm` | GORM read/write SQLite — `GetLeaderboard` sorts by `score DESC`, `created_at ASC` |

## Data flow

**Load questions (GET)**

- **DB:** `questions` + `options` tables (`correct_option_id` per question — not exposed via API)
- **Repository** → **Use case** strips answers → **Handler** → JSON `{ "questions": [...] }`
- **Frontend** stores in `examStore.questions` for display and `answers`

**Submit (POST)**

- **Frontend** sends `candidateName` and `answers` (keys are string question ids)
- **Use case** loads questions + answers from DB as before → compares to `answers` → `score`, `total`
- **DB:** `INSERT` into `exam_results` (name, score, total, `answers_json`)

**Leaderboard (GET)**

- **DB:** reads `exam_results` — sorted by highest score first; ties broken by earliest `created_at`
- **Repository** → **Use case** assigns `rank` and strips deep data → **Handler** → JSON `{ "entries": [...] }`
- **Frontend** stores in `examStore.leaderboard` for `LeaderboardView`

## API contract summary (`api.md`)

| Topic | Status |
|--------|--------|
| `GET /api/questions` does not send `correctOptionId` | OK — use case DTO has no answer fields |
| `POST /api/submit` body `candidateName`, `answers` (string keys) | OK |
| Response `{ candidateName, score, total }` | OK — result view uses `score` and `totalQuestions` (= loaded count), which should align with `total` |
| `GET /api/leaderboard` → `{ entries: LeaderboardEntryDTO[] }` | OK — ranked from `exam_results` sorted by score descending then `created_at` ascending |

Details and examples: [api.md](./api.md)

## Diagram — Frontend relationships

```mermaid
flowchart LR
  subgraph views [Views]
    E[ExamView]
    R[ResultView]
    L[LeaderboardView]
  end
  subgraph store [Pinia examStore]
    CN[candidateName]
    Q[questions]
    A[answers]
    S[score]
    LB[leaderboard]
  end
  E -->|loadQuestions GET| Q
  E -->|setAnswer| A
  E -->|submitExam| S
  S --> R
  R -->|goLeaderboard| L
  L -->|loadLeaderboard GET| LB
  R -->|resetExam| CN
  R -->|resetExam| A
  R -->|resetExam| S
  R -->|resetExam| LB
  L -->|resetExam| CN
  L -->|resetExam| A
  L -->|resetExam| S
  L -->|resetExam| LB
```

## Diagram — Backend request sequence

```mermaid
sequenceDiagram
  participant C as Client
  participant H as Gin Handler
  participant U as Usecase
  participant R as Repository
  participant DB as SQLite

  C->>H: GET /api/questions
  H->>U: GetQuestions()
  U->>R: GetQuestions()
  R->>DB: SELECT questions + options
  DB-->>R: rows
  R-->>U: []Question
  U-->>H: []QuestionDTO (no correct id)
  H-->>C: JSON

  C->>H: POST /api/submit
  H->>U: SubmitExam(name, answers)
  U->>R: GetQuestions()
  R->>DB: SELECT
  U->>U: ScoreAnswers()
  U->>R: SaveExamResult()
  R->>DB: INSERT exam_results
  U-->>H: SubmitResponse
  H-->>C: JSON

  C->>H: GET /api/leaderboard
  H->>U: GetLeaderboard(limit)
  U->>R: GetLeaderboard(limit)
  R->>DB: SELECT exam_results ORDER BY score DESC, created_at ASC LIMIT n
  DB-->>R: rows
  R-->>U: []ExamResult
  U-->>H: []LeaderboardEntryDTO
  H-->>C: JSON entries
```

## Frontend tech stack

| Technology | Role |
|-----------|--------|
| **Vue 3** | UI framework — Composition API + `<script setup>` |
| **Vite** | Build and dev server |
| **Tailwind CSS** | Utility-first styling, responsive, mobile-first |
| **Vue Router** | Routes: `/` (exam), `/result` (score), `/leaderboard` (rankings) |
| **Pinia** | State: candidate name, questions, answers, score, leaderboard — loads questions and rankings from API only |

## Backend tech stack

| Technology | Role |
|-----------|--------|
| **Go** | Language and runtime |
| **Gin** | HTTP router / middleware |
| **GORM** | ORM for SQLite |
| **SQLite** | Single-file DB (`backend/data/exam.db`) — zero extra install |
| **testify** | `assert` + `mock` for use case unit tests |

## Why Vue 3 + Pinia

- **Vue 3** Composition API groups logic by feature clearly
- **Pinia** keeps **exam** state out of components so `ExamView` / `ResultView` focus on UI and events

## Why Go + Gin + SQLite

- **Go** — simple single-binary deploy, clear concurrency
- **Gin** — widely used, middleware fits REST
- **SQLite** — good for learning/demos — no separate DB server; can move to PostgreSQL when scaling
- **Pragmatic Clean Architecture**: handler → use case → repository — test use cases with mock repositories without touching SQLite

## Frontend folder layout

- `frontend/src/views/` — main screens per route
- `frontend/src/components/` — reusable subcomponents
- `frontend/src/stores/` — Pinia (`examStore`)
- `frontend/src/router/` — routes and meta (title)
- `frontend/src/api/` — HTTP (`client.js`)
- `frontend/src/assets/` — global CSS and Tailwind theme

## Backend folder layout (Pragmatic Clean Architecture)

```
backend/
├── cmd/api/main.go          # entry, SQLite path, AutoMigrate, seed, DI, Gin
├── internal/
│   ├── models/              # Question, Option, ExamResult
│   ├── repository/          # GORM: GetQuestions, SaveExamResult, GetLeaderboard, migrate, seed
│   ├── usecase/             # Exam, ports (interfaces), ScoreAnswers, GetLeaderboard
│   └── handler/             # Gin: GET /api/questions, POST /api/submit, GET /api/leaderboard
├── scripts/                 # optional dev SQL (mock/clear exam_results) — not under cmd/ (cmd = mains only)
├── go.mod
└── data/exam.db             # created at run time (in .gitignore)
```

- **Handler** — JSON in/out, minimal business logic
- **Use case** — `GetQuestions` (map to DTO without answers), `SubmitExam` (load answers from DB → score → `SaveExamResult`), `GetLeaderboard` (ranked DTO without raw answers)
- **Repository** — GORM/SQLite only — includes `GetLeaderboard` for reading `exam_results`

Endpoint details and JSON examples: [api.md](./api.md)

## Frontend and backend communication

In short: API base `http://localhost:8080` — full tables and payloads in [api.md](./api.md)

More planning and roadmap: [planning.md](./planning.md)
