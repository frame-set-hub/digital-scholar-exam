# API Reference (Backend)

Single **canonical** reference for endpoints and example request/response — update here when the API changes.

Base URL (dev): `http://localhost:8080`

| Method | Path | Description |
|--------|------|-----------|
<<<<<<< HEAD
| `GET` | `/api/questions` | รายการข้อสอบ + ตัวเลือก **ไม่มี** `correctOptionId` |
| `POST` | `/api/submit` | รับคำตอบ → คำนวณคะแนนที่เซิร์ฟเวอร์ → บันทึก `exam_results` |
| `GET` | `/api/leaderboard` | อันดับผู้สอบจาก `exam_results` (คะแนนสูงก่อน, คะแนนเท่ากันคนสอบก่อนอยู่บน) — ไม่ส่ง `answers` |

## GET `/api/leaderboard`

- Query (optional): `limit` — จำนวนอันดับสูงสุด (ค่าเริ่มต้น 20, สูงสุด 20)
- Response ตัวอย่าง: `{ "entries": [ { "rank", "candidateName", "score", "total", "createdAt" } ] }`
=======
| `GET` | `/api/questions` | Question list + options **without** `correctOptionId` |
| `POST` | `/api/submit` | Accept answers → score on server → persist `exam_results` |
>>>>>>> 59f10ee (Refactor documentation for clarity and consistency; update execute.md, README.md, RULE.md, and various API references to enhance user understanding and maintainability.)

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

Processing flow: [architech.md](./architech.md)

## Troubleshooting (Chrome DevTools)

The message **"Failed to load response data. No resource with given identifier found"** often appears when there are **duplicate requests** or you inspect an old response after refresh — the frontend avoids repeating `GET /api/questions` when questions are already in Pinia.

If it persists: enable **Preserve log**, click the latest request right after load — or use `curl` to verify the body.
