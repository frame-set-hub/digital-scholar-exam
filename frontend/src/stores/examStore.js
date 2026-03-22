import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import router from '@/router'
import { apiUrl, fetchJSON } from '@/api/client'

export const useExamStore = defineStore('exam', () => {
  const candidateName = ref('')
  const questions = ref([])
  const answers = ref({})
  const score = ref(null)

  const leaderboard = ref([])
  /** เมื่อโหลดพร้อม `forCandidate` — อันดับรวม + inTopList (อยู่ในช่วง top N ของ response หรือไม่) */
  const leaderboardYourEntry = ref(null)
  const leaderboardState = ref('idle') // idle | loading | error
  const leaderboardError = ref(null)

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

  /** ชื่อสำหรับ yourEntry: `?forCandidate=` บน URL ก่อน (ตั้งใจจากลิงก์/แชร์) แล้วค่อย Pinia */
  function resolveLeaderboardCandidateName() {
    const raw = router.currentRoute.value.query.forCandidate
    const q = Array.isArray(raw) ? raw[0] : raw
    const fromQuery = typeof q === 'string' ? q.trim() : ''
    if (fromQuery) return fromQuery
    return (candidateName.value && String(candidateName.value).trim()) || ''
  }

  /**
   * GET /api/leaderboard — ส่ง forCandidate เพื่อขอ yourEntry
   * @param {string} [explicitForCandidate] — ส่งจาก LeaderboardView จาก route.query (ตัดปัญหา store ชนคิวกับ URL / router ใน store)
   */
  async function loadLeaderboard(explicitForCandidate) {
    leaderboardState.value = 'loading'
    leaderboardError.value = null
    leaderboardYourEntry.value = null
    try {
      let name = ''
      if (typeof explicitForCandidate === 'string' && explicitForCandidate.trim()) {
        name = explicitForCandidate.trim()
      } else {
        name = resolveLeaderboardCandidateName()
      }
      const q =
        name.length > 0
          ? `?forCandidate=${encodeURIComponent(name)}`
          : ''
      const data = await fetchJSON(apiUrl(`/api/leaderboard${q}`))
      leaderboard.value = Array.isArray(data.entries) ? data.entries : []
      leaderboardYourEntry.value = data.yourEntry ?? null
      leaderboardState.value = 'idle'
    } catch (e) {
      leaderboard.value = []
      leaderboardYourEntry.value = null
      leaderboardError.value =
        e?.message ||
        'ไม่สามารถโหลดกระดานจัดอันดับ — ตรวจสอบว่า backend รันที่ :8080'
      leaderboardState.value = 'error'
    }
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
    leaderboard.value = []
    leaderboardYourEntry.value = null
    leaderboardState.value = 'idle'
    leaderboardError.value = null
    router.push({ name: 'exam' })
  }

  return {
    candidateName,
    questions,
    answers,
    score,
    leaderboard,
    leaderboardYourEntry,
    leaderboardState,
    leaderboardError,
    totalQuestions,
    loadState,
    loadError,
    setAnswer,
    loadQuestions,
    loadLeaderboard,
    submitExam,
    resetExam,
  }
})
