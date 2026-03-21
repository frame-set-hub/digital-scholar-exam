# API Reference (Backend)

แหล่งอ้างอิง **เดียว** สำหรับ endpoint และตัวอย่าง request/response — อัปเดตที่นี่เมื่อ API เปลี่ยน

Base URL (dev): `http://localhost:8080`

| Method | Path | คำอธิบาย |
|--------|------|-----------|
| `GET` | `/api/questions` | รายการข้อสอบ + ตัวเลือก **ไม่มี** `correctOptionId` |
| `POST` | `/api/submit` | รับคำตอบ → คำนวณคะแนนที่เซิร์ฟเวอร์ → บันทึก `exam_results` |

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

Flow การประมวลผลอยู่ใน [architech.md](./architech.md)

## Troubleshooting (Chrome DevTools)

ข้อความ **"Failed to load response data. No resource with given identifier found"** มักเกิดเมื่อมี **หลาย request ซ้ำ** หรือเปิดดู response ของ request เก่าหลังรีเฟรช — ฝั่ง frontend ลดการเรียก `GET /api/questions` ซ้ำเมื่อมีข้อสอบใน Pinia แล้ว

ถ้ายังเจอ: เปิด **Preserve log**, คลิก request ล่าสุดทันทีหลังโหลด — หรือใช้ `curl` ยืนยัน body แทน
