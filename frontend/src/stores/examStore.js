import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import router from '@/router'

/** Mock ข้อสอบ — แทนที่ด้วยข้อมูลจาก API ภายหลัง */
export const MOCK_QUESTIONS = [
  {
    id: 1,
    prompt: 'ข้อใดต่างจากข้ออื่น',
    subtitle: null,
    options: [
      { id: '1a', letter: 'A', text: '3' },
      { id: '1b', letter: 'B', text: '5' },
      { id: '1c', letter: 'C', text: '9' },
      { id: '1d', letter: 'D', text: '11' },
    ],
    correctOptionId: '1c',
  },
  {
    id: 2,
    prompt: 'X + 2 = 4',
    subtitle: 'จงหาค่า X',
    options: [
      { id: '2a', letter: 'A', text: '1' },
      { id: '2b', letter: 'B', text: '2' },
      { id: '2c', letter: 'C', text: '3' },
      { id: '2d', letter: 'D', text: '4' },
    ],
    correctOptionId: '2b',
  },
  {
    id: 3,
    prompt: '2 + 2 = ?',
    subtitle: null,
    options: [
      { id: '3a', letter: 'A', text: '3' },
      { id: '3b', letter: 'B', text: '4' },
      { id: '3c', letter: 'C', text: '5' },
      { id: '3d', letter: 'D', text: '6' },
    ],
    correctOptionId: '3b',
  },
]

export const useExamStore = defineStore('exam', () => {
  const candidateName = ref('')
  const questions = ref([...MOCK_QUESTIONS])
  /** คีย์ = id ข้อ, ค่า = id ตัวเลือกที่เลือก */
  const answers = ref({})
  /** null = ยังไม่ส่ง, ตัวเลข = คะแนนหลังส่ง */
  const score = ref(null)

  const totalQuestions = computed(() => questions.value.length)

  function setAnswer(questionId, optionId) {
    answers.value = { ...answers.value, [questionId]: optionId }
  }

  function computeScore() {
    let correct = 0
    for (const q of questions.value) {
      if (answers.value[q.id] === q.correctOptionId) correct += 1
    }
    return correct
  }

  function submitExam() {
    score.value = computeScore()
    router.push({ name: 'result' })
  }

  function resetExam() {
    candidateName.value = ''
    answers.value = {}
    score.value = null
    router.push({ name: 'exam' })
  }

  return {
    candidateName,
    questions,
    answers,
    score,
    totalQuestions,
    setAnswer,
    submitExam,
    resetExam,
  }
})
