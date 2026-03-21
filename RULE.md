# กฎการอัปเดตเอกสาร (Documentation sync)

เมื่อมีการ**แก้โค้ด** (หรือโครงสร้างโปรเจกต์) แล้วการเปลี่ยนแปลงนั้น**กระทบ**เนื้อหาในเอกสาร Markdown — ให้**อัปเดตไฟล์ `.md` ที่เกี่ยวข้องในเชิงเดียวกัน** ไม่ปล่อยให้เอกสารอ้างพฤติกรรมเก่า

## สิ่งที่ถือว่า “กระทบ”

- เปลี่ยน flow, API, state, หรือชื่อไฟล์/โฟลเดอร์ที่เอกสารอ้างถึง
- ย้าย/ลบ/รวมบรรทัดจนช่วงบรรทัดใน [docs/code_analyze.md](./docs/code_analyze.md) ไม่ตรงกับโค้ดจริง
- เปลี่ยนพฤติกรรมที่ [README.md](./README.md), [execute.md](./execute.md), [docs/planning.md](./docs/planning.md), [docs/architech.md](./docs/architech.md), [docs/api.md](./docs/api.md) อธิบายไว้

## แนวทางปฏิบัติ

1. ระบุว่าแก้ไขกระทบส่วนใด (เช่น ไม่มี fallback mock บน frontend แล้ว → แก้ทุกที่ที่พูดถึง MOCK / fallback)
2. อัปเดต **ทุกไฟล์ `.md` ที่เกี่ยวข้อง** ในรอบเดียวกับการ merge การแก้โค้ด (หรือทันทีหลัง merge)
3. สำหรับ **code map** ใน `docs/code_analyze.md`: ตรวจช่วงบรรทัดกับไฟล์จริงใน repo (เปิดไฟล์หรือใช้เครื่องมือนับบรรทัด) แล้วแก้ตารางให้สอดคล้อง

## แหล่งเอกสารหลัก

| บทบาท | ไฟล์ |
|--------|------|
| จุดเข้า repo | [README.md](./README.md) |
| ดัชนีเอกสารเชิงลึก | [docs/README.md](./docs/README.md) |
| อ่านโค้ดทีละไฟล์ (บรรทัด) | [docs/code_analyze.md](./docs/code_analyze.md) |
| สถาปัตยกรรม / flow | [docs/architech.md](./docs/architech.md) |
| API | [docs/api.md](./docs/api.md) |
| แผน / roadmap | [docs/planning.md](./docs/planning.md) |
| Progress | [execute.md](./execute.md) |

ไฟล์นี้เป็นกฎโปรเจกต์ — ไม่แทนที่คำอธิบายในแต่ละเอกสาร แต่บังคับให้ซิงค์เมื่อมีการเปลี่ยนที่กระทบเอกสาร
