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
- Example response: `{ "entries": [ { "rank", "candidateName", "score", "total", "createdAt" } ] }`

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
| `400` | Body invalid, `answers` empty, or ชื่อว่างหลัง trim |
| `409` | ชื่อผู้สอบซ้ำกับผลที่บันทึกแล้ว (`exam_results.candidate_name` เทียบตรงหลัง trim) |
| `500` | ฐานข้อมูล / use case อื่น |

ข้อความตัวอย่าง: `409` → `ชื่อนี้ถูกใช้ส่งข้อสอบแล้ว — กรุณาใช้ชื่ออื่น`

Processing flow: [architech.md](./architech.md)

## Troubleshooting (Chrome DevTools)

The message **"Failed to load response data. No resource with given identifier found"** often appears when there are **duplicate requests** or you inspect an old response after refresh — the frontend avoids repeating `GET /api/questions` when questions are already in Pinia.

If it persists: enable **Preserve log**, click the latest request right after load — or use `curl` to verify the body.
