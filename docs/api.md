# API Reference (Backend)

แหล่งอ้างอิง **เดียว** สำหรับ endpoint และตัวอย่าง request/response — อัปเดตที่นี่เมื่อ API เปลี่ยน

Base URL (dev): `http://localhost:8080`

| Method | Path | คำอธิบาย |
|--------|------|-----------|
| `GET` | `/api/questions` | รายการข้อสอบ + ตัวเลือก **ไม่มี** `correctOptionId` |
| `POST` | `/api/submit` | รับคำตอบ → คำนวณคะแนนที่เซิร์ฟเวอร์ → บันทึก `exam_results` |
| `GET` | `/api/leaderboard` | อันดับผู้สอบจาก `exam_results` (คะแนนสูงก่อน; คะแนนเท่ากันคนส่งก่อนอยู่บน) — ไม่ส่ง `answers` |

## GET `/api/leaderboard`

- Query (optional): `limit` — จำนวนอันดับสูงสุด (ค่าเริ่มต้น 20, สูงสุด 20)
- Response ตัวอย่าง: `{ "entries": [ { "rank", "candidateName", "score", "total", "createdAt" } ] }`

## POST `/api/submit`

**Request body (JSON)**

```json
{
  "candidateName": "ชื่อผู้สอบ",
  "answers": {
    "1": "1c",
    "2": "2b",
    "3": "3b"
  }
}
```

- คีย์ใน `answers` เป็น **string** ของ question `id` ตรงกับฐานข้อมูล
- Response ตัวอย่าง: `{ "candidateName": "...", "score": 3, "total": 3 }`

**ข้อผิดพลาด (JSON body `{ "error": "..." }`)**

| HTTP | เมื่อไหร่ |
|------|-----------|
| `400` | body ไม่ถูกต้อง, `answers` ว่าง, หรือชื่อว่างหลัง trim |
| `409` | ชื่อผู้สอบซ้ำกับผลที่บันทึกแล้ว (`exam_results.candidate_name` เทียบตรงหลัง trim) |
| `500` | ฐานข้อมูล / use case อื่น |

ข้อความตัวอย่าง: `409` → `ชื่อนี้ถูกใช้ส่งข้อสอบแล้ว — กรุณาใช้ชื่ออื่น`

Flow การประมวลผล: [architech.md](./architech.md)

## Troubleshooting (Chrome DevTools)

ข้อความ **"Failed to load response data. No resource with given identifier found"** มักเกิดเมื่อมี **หลาย request ซ้ำ** หรือเปิดดู response ของ request เก่าหลังรีเฟรช — ฝั่ง frontend ลดการเรียก `GET /api/questions` ซ้ำเมื่อมีข้อสอบใน Pinia แล้ว

ถ้ายังเจอ: เปิด **Preserve log**, คลิก request ล่าสุดทันทีหลังโหลด — หรือใช้ `curl` ยืนยัน body แทน
