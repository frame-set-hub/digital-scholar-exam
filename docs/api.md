# API Reference (Backend)

Single **canonical** reference for endpoints and example request/response — update here when the API changes.

Base URL (dev): `http://localhost:8080`

| Method | Path | Description |
|--------|------|-----------|
| `GET` | `/api/questions` | Question list + options **without** `correctOptionId` |
| `POST` | `/api/submit` | Accept answers → score on server → persist `exam_results` |
| `GET` | `/api/leaderboard` | Ranked candidates from `exam_results` (highest score first; ties broken by earliest submission) — does not include `answers` |

## GET `/api/leaderboard`

- Query (optional): `limit` — max number of entries (default 20, max 20)
- Query (optional): `forCandidate` — candidate name (URL-encoded) to include **`yourEntry`**: their **global rank** when a row exists in `exam_results`. If omitted, `yourEntry` is `null` (use e.g. `?forCandidate=mock_zero` in `curl` to get an object).
- Response shape: `yourEntry` is listed **first** in JSON (before `entries`) for easier inspection with `curl | head`.
- Example: `{ "yourEntry": null, "entries": [ { "rank", "candidateName", "score", "total", "createdAt" } ] }`
- When `forCandidate` is present **and** a matching row exists: `yourEntry` is `{ "rank", "candidateName", "score", "total", "createdAt", "inTopList" }` — `inTopList` is `true` when global rank is within the first `limit` rows (same window as `entries`); `false` when rank is outside that window (e.g. rank 21 with `limit=20`).

## POST `/api/submit`

**Request body (JSON)**

```json
{
  "candidateName": "Jane Doe",
  "answers": {
    "1": "1c",
    "2": "2b",
    "3": "3b"
  }
}
```

- Keys in `answers` are **strings** of question `id` values as in the database
- Example response: `{ "candidateName": "...", "score": 3, "total": 3 }`

**Errors (JSON body `{ "error": "..." }`)**

| HTTP | When |
|------|------|
| `400` | Body invalid, `answers` empty, or candidate name empty after trim |
| `409` | Duplicate candidate name already stored in `exam_results` (exact match after trim) |
| `500` | Database or other server error |

Example `409` message: `ชื่อนี้ถูกใช้ส่งข้อสอบแล้ว — กรุณาใช้ชื่ออื่น` (Thai UI string from the API)

Processing flow: [architech.md](./architech.md)

## Troubleshooting (Chrome DevTools)

The message **"Failed to load response data. No resource with given identifier found"** often appears when there are **duplicate requests** or you inspect an old response after refresh — the frontend avoids repeating `GET /api/questions` when questions are already in Pinia.

If it persists: enable **Preserve log**, click the latest request right after load — or use `curl` to verify the body.
