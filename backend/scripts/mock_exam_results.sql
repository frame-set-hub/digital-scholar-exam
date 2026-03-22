-- Mock rows for exam_results (7 questions, total = 7).
-- Run from repo: sqlite3 backend/data/exam.db < backend/scripts/mock_exam_results.sql
-- Prefer stopping the API first so the DB is not locked for write.
-- To wipe all results first: clear_exam_results.sql
--
-- อย่ารันสคริปต์นี้ซ้ำโดยไม่ clear ก่อน — แต่ละครั้งจะ INSERT เพิ่ม ทำให้ชื่อซ้ำหลายแถวใน Leaderboard (ข้อมูลไม่ตรงกับกฎห้ามชื่อซ้ำตอน submit)

INSERT INTO exam_results (candidate_name, score, total, answers_json, created_at) VALUES
(
  'สมชาย',
  7,
  7,
  '{"1":"1c","2":"2b","3":"3b","4":"4c","5":"5b","6":"6c","7":"7b"}',
  '2025-03-21 10:00:00'
),
(
  'วิชัย',
  6,
  7,
  '{"1":"1c","2":"2b","3":"3b","4":"4c","5":"5b","6":"6c","7":"7a"}',
  '2025-03-21 11:30:00'
),
(
  'Jamie',
  5,
  7,
  '{"1":"1a","2":"2b","3":"3b","4":"4c","5":"5b","6":"6c","7":"7b"}',
  '2025-03-21 12:00:00'
),
(
  'mock_mid',
  4,
  7,
  '{"1":"1c","2":"2a","3":"3b","4":"4c","5":"5a","6":"6c","7":"7b"}',
  '2025-03-22 08:15:00'
),
(
  'mock_low',
  2,
  7,
  '{"1":"1a","2":"2a","3":"3a","4":"4a","5":"5a","6":"6a","7":"7a"}',
  '2025-03-22 09:00:00'
),
(
  'mock_zero',
  0,
  7,
  '{"1":"1a","2":"2a","3":"3a","4":"4a","5":"5a","6":"6a","7":"7a"}',
  '2025-03-22 09:30:00'
);
