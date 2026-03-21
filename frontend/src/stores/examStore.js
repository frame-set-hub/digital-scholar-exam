import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import router from '@/router'
import { apiUrl, fetchJSON } from '@/api/client'

export const useExamStore = defineStore('exam', () => {
  const candidateName = ref('')
  const questions = ref([])
  const answers = ref({})
  const score = ref(null)

  const loadState = ref('idle') // idle | loading | error
  const loadError = ref(null)

  /** กันโหลดซ้ำพร้อมกัน / ลดจำนวน request เพื่อให้ DevTools ไม่ทิ้ง response ของ request เก่าเยอะเกินไป */
  let loadQuestionsInflight = null

  const totalQuestions = computed(() => questions.value.length)

  function setAnswer(questionId, optionId) {
    answers.value = { ...answers.value, [questionId]: optionId }
  }

  /**
   * โหลดข้อสอบจาก API เท่านั้น: GET /api/questions
   * @param {{ force?: boolean }} [options] — force=true บังคับดึงใหม่
   */
  async function loadQuestions(options = {}) {
    const force = options.force === true

    if (!force && questions.value.length > 0) {
      return
    }

    if (loadQuestionsInflight) {
      return loadQuestionsInflight
    }

    loadState.value = 'loading'
    loadError.value = null

    loadQuestionsInflight = (async () => {
      try {
        const data = await fetchJSON(apiUrl('/api/questions'))
        questions.value = data.questions ?? []
        loadState.value = 'idle'
      } catch (e) {
        questions.value = []
        loadError.value =
          e?.message ||
          'ไม่สามารถโหลดข้อสอบจากเซิร์ฟเวอร์ — ตรวจสอบว่า backend รันที่ :8080 และ proxy ใน Vite ถูกต้อง'
        loadState.value = 'error'
      } finally {
        loadQuestionsInflight = null
      }
    })()

    return loadQuestionsInflight
  }

  function answersForSubmit() {
    const out = {}
    for (const [k, v] of Object.entries(answers.value)) {
      out[String(k)] = v
    }
    return out
  }

  async function submitExam() {
    const data = await fetchJSON(apiUrl('/api/submit'), {
      method: 'POST',
      body: JSON.stringify({
        candidateName: candidateName.value.trim(),
        answers: answersForSubmit(),
      }),
    })
    score.value = data.score
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
    loadState,
    loadError,
    setAnswer,
    loadQuestions,
    submitExam,
    resetExam,
  }
})
