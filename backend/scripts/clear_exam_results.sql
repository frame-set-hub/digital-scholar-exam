-- ลบผลการสอบทั้งหมด (ตาราง exam_results เท่านั้น — ไม่แตะ questions / options)
-- Run from repo: sqlite3 backend/data/exam.db < backend/scripts/clear_exam_results.sql
-- หยุด backend ก่อน (กันติดล็อกไฟล์ DB)

DELETE FROM exam_results;
