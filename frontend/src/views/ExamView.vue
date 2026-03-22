<script setup>
import { ref, computed, onMounted, nextTick } from 'vue'
import { storeToRefs } from 'pinia'
import { useExamStore } from '@/stores/examStore'

const exam = useExamStore()
const { candidateName, questions, answers, loadState, loadError } = storeToRefs(exam)

const nameError = ref('')
const submitError = ref('')
/** After submit while incomplete — unanswered questions get a red border (scroll to first gap). */
const showUnansweredHighlight = ref(false)
/** DOM refs per question section (for scrollIntoView). */
const sectionRefs = {}

function setSectionRef(questionId, el) {
  if (el) {
    sectionRefs[questionId] = el
  } else {
    delete sectionRefs[questionId]
  }
}

const nameBlockRef = ref(null)
const submitSectionRef = ref(null)
const submitBtnRef = ref(null)

onMounted(() => {
  exam.loadQuestions()
})

const allAnswered = computed(() => {
  if (questions.value.length === 0) return false
  return questions.value.every((q) => answers.value[q.id] != null)
})

function isSelected(questionId, optionId) {
  return answers.value[questionId] === optionId
}

function selectOption(questionId, optionId) {
  exam.setAnswer(questionId, optionId)
}

function isQuestionUnanswered(questionId) {
  return answers.value[questionId] == null
}

function questionSectionClasses(questionId) {
  const invalid = showUnansweredHighlight.value && isQuestionUnanswered(questionId)
  return [
    'rounded-3xl border p-8 shadow-[0_8px_32px_rgba(25,28,30,0.04)] md:p-12',
    invalid
      ? 'border-red-500 bg-red-50/80 shadow-[0_0_0_4px_rgba(239,68,68,0.12)]'
      : 'border-outline-variant/10 bg-surface-container-lowest',
  ]
}

function scrollToName() {
  nameBlockRef.value?.scrollIntoView({ behavior: 'smooth', block: 'center' })
}

function scrollToSubmit() {
  submitSectionRef.value?.scrollIntoView({ behavior: 'smooth', block: 'center' })
}

function focusSubmitButton() {
  submitBtnRef.value?.focus()
}

async function handleSubmit() {
  nameError.value = ''
  submitError.value = ''
  showUnansweredHighlight.value = false
  const name = candidateName.value.trim()
  if (!name) {
    nameError.value = 'กรุณากรอกชื่อผู้สอบ'
    await nextTick()
    scrollToName()
    return
  }
  if (!allAnswered.value) {
    submitError.value = 'กรุณาตอบให้ครบทุกข้อ'
    showUnansweredHighlight.value = true
    const firstUnanswered = questions.value.find((q) => answers.value[q.id] == null)
    if (firstUnanswered) {
      await nextTick()
      const el = sectionRefs[firstUnanswered.id]
      el?.scrollIntoView({ behavior: 'smooth', block: 'center' })
    }
    return
  }
  try {
    await exam.submitExam()
  } catch (e) {
    const status = e?.status
    const msg = e?.message || 'ส่งข้อสอบไม่สำเร็จ'
    const isNameConflict = status === 409
    const isNameRequiredFromAPI =
      status === 400 && typeof msg === 'string' && msg.includes('กรุณากรอกชื่อ')
    if (isNameConflict || isNameRequiredFromAPI) {
      nameError.value = msg
      await nextTick()
      scrollToName()
      return
    }
    submitError.value =
      e?.message?.includes('fetch') || e?.name === 'TypeError'
        ? 'ส่งข้อสอบไม่สำเร็จ — ตรวจสอบว่า backend รันที่ :8080 และลองใหม่'
        : msg
    await nextTick()
    scrollToSubmit()
  }
}

/** คลาสเดียวกันทุกข้อ — choice card (min-h, padding, โครงสร้าง indicator + ข้อความ) */
function optionCardClasses(questionId, optionId) {
  const selected = isSelected(questionId, optionId)
  return [
    'group flex min-h-[88px] w-full items-center gap-4 rounded-xl border-l-4 p-6 text-left transition-all duration-300',
    'focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary',
    selected
      ? 'border-primary bg-secondary-container shadow-[0_4px_16px_rgba(53,37,205,0.18)] ring-1 ring-primary/30'
      : 'border-transparent bg-surface-container-low hover:bg-surface-container-high',
  ]
}

function indicatorClasses(questionId, optionId) {
  const selected = isSelected(questionId, optionId)
  return [
    'option-indicator flex h-8 w-8 shrink-0 items-center justify-center rounded-full border-2 text-sm font-bold transition-all',
    selected
      ? 'border-primary bg-primary text-on-primary'
      : 'border-outline-variant text-on-surface-variant group-hover:border-primary group-hover:text-primary',
  ]
}

function optionTextClasses(questionId, optionId) {
  const selected = isSelected(questionId, optionId)
  return [
    'option-text text-lg font-medium transition-colors',
    selected ? 'font-bold text-primary' : 'text-on-surface',
  ]
}
</script>

<template>
  <div class="relative flex min-h-screen flex-col overflow-x-hidden bg-background font-sans text-on-surface">
    <!-- ไม่มี Navbar / เมนูปลอม — เริ่มที่เนื้อหาหลักตาม ui-example -->
    <main class="mx-auto w-full max-w-3xl flex-1 px-4 py-10 pb-28 sm:px-6 sm:py-12">
      <p
        v-if="loadError"
        class="mb-6 rounded-xl border border-amber-200 bg-amber-50 px-4 py-3 text-sm text-amber-900"
        role="status"
      >
        {{ loadError }}
      </p>
      <!-- Header + Candidate Name -->
      <div ref="nameBlockRef" class="mb-12 scroll-mt-8">
        <div class="flex flex-col justify-between gap-6 md:flex-row md:items-end">
          <div class="space-y-2">
            <span
              class="inline-flex items-center gap-2 rounded-full bg-tertiary-fixed px-3 py-1 text-xs font-bold uppercase tracking-wider text-on-tertiary-fixed"
            >
              <span class="material-symbols-outlined text-[14px]">verified</span>
              Live Session
            </span>
            <h1 class="font-display text-3xl font-extrabold tracking-tight text-on-surface sm:text-4xl">
              IT 10-1 Exam
            </h1>
            <p class="font-sans text-sm text-secondary sm:text-base">Mathematics and Logical Reasoning Module</p>
          </div>
          <div class="w-full md:w-72">
            <label
              class="mb-2 block text-sm font-semibold text-on-surface-variant"
              for="candidate-name"
              >NAME</label
            >
            <div class="group relative">
              <input
                id="candidate-name"
                v-model="candidateName"
                type="text"
                autocomplete="name"
                placeholder="Enter your full name"
                :aria-invalid="nameError ? 'true' : 'false'"
                :aria-describedby="nameError ? 'candidate-name-error' : undefined"
                class="w-full rounded-t-lg border-none border-b-2 bg-surface-container-highest py-3 pl-4 pr-4 font-sans text-on-surface transition-all duration-300 placeholder:text-outline/50 focus:ring-0"
                :class="
                  nameError
                    ? 'border-red-500 focus:border-red-600'
                    : 'border-transparent focus:border-primary'
                "
                @keydown.enter.prevent="focusSubmitButton"
              />
              <div
                class="pointer-events-none absolute inset-x-0 bottom-0 h-0.5 transition-colors"
                :class="
                  nameError
                    ? 'bg-red-500'
                    : 'bg-outline-variant/15 group-focus-within:bg-primary'
                "
              />
            </div>
            <p
              v-if="nameError"
              id="candidate-name-error"
              class="mt-2 text-sm font-medium text-red-600"
              role="alert"
            >
              {{ nameError }}
            </p>
          </div>
        </div>
      </div>

      <div
        v-if="loadState === 'loading'"
        class="flex flex-col items-center justify-center gap-4 py-24 text-secondary"
      >
        <span
          class="inline-block h-10 w-10 animate-spin rounded-full border-2 border-primary border-t-transparent"
          aria-hidden="true"
        />
        <p class="text-sm font-medium">กำลังโหลดข้อสอบจากเซิร์ฟเวอร์…</p>
      </div>

      <!-- Questions: v-for จาก Store — ทุกข้อใช้ choice card สไตล์เดียวกัน -->
      <div v-else class="space-y-12">
        <section
          v-for="(q, index) in questions"
          :id="'question-' + q.id"
          :key="q.id"
          :ref="(el) => setSectionRef(q.id, el)"
          :class="questionSectionClasses(q.id)"
        >
          <div class="mb-8 flex items-start gap-4">
            <span
              class="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary font-display text-sm font-bold text-on-primary sm:text-base"
              >{{ index + 1 }}</span
            >
            <div>
              <h2 class="font-display text-xl font-bold leading-tight text-on-surface sm:text-2xl">
                {{ q.prompt }}
              </h2>
              <p v-if="q.subtitle" class="mt-1 font-sans text-sm text-secondary">{{ q.subtitle }}</p>
            </div>
          </div>

          <div class="flex flex-col gap-4">
            <button
              v-for="opt in q.options"
              :key="opt.id"
              type="button"
              :class="optionCardClasses(q.id, opt.id)"
              @click="selectOption(q.id, opt.id)"
            >
              <span :class="indicatorClasses(q.id, opt.id)">{{ opt.letter }}</span>
              <span :class="optionTextClasses(q.id, opt.id)">{{ opt.text }}</span>
            </button>
          </div>
          <div
            v-if="showUnansweredHighlight && answers[q.id] == null"
            class="mt-6 flex items-center gap-2 font-sans text-sm font-bold text-red-600"
            role="alert"
          >
            <span class="material-symbols-outlined text-[18px]" aria-hidden="true">error</span>
            โปรดเลือกคำตอบก่อนดำเนินการต่อ
          </div>
        </section>
      </div>

      <!-- Submit -->
      <div
        v-if="loadState !== 'loading'"
        ref="submitSectionRef"
        class="mt-16 flex scroll-mt-8 flex-col items-center gap-6"
      >
        <p
          v-if="submitError"
          :class="
            showUnansweredHighlight
              ? 'animate-pulse text-center text-sm font-bold text-red-600'
              : 'text-center text-sm font-medium text-red-600'
          "
          role="alert"
        >
          {{ submitError }}
        </p>
        <button
          ref="submitBtnRef"
          type="button"
          :disabled="questions.length === 0"
          class="group relative w-full max-w-md overflow-hidden rounded-2xl bg-gradient-to-br from-primary to-primary-container p-4 text-xl font-bold text-on-primary shadow-xl shadow-indigo-500/20 transition-all enabled:active:scale-[0.98] enabled:hover:shadow-indigo-500/35 disabled:cursor-not-allowed disabled:opacity-50"
          @click="handleSubmit"
        >
          <span class="relative z-10 flex items-center justify-center gap-3">
            Submit Exam
            <span class="material-symbols-outlined transition-transform group-hover:translate-x-1">arrow_forward</span>
          </span>
          <div class="absolute inset-0 bg-white/10 opacity-0 transition-opacity group-hover:opacity-100" />
        </button>
      </div>
    </main>

    <div
      class="pointer-events-none fixed top-0 right-0 -z-10 h-full w-1/2 bg-gradient-to-l from-primary/5 to-transparent blur-3xl"
    />
    <div
      class="pointer-events-none fixed bottom-0 left-0 -z-10 h-2/3 w-1/3 bg-gradient-to-tr from-tertiary/5 to-transparent blur-3xl"
    />
  </div>
</template>

<style scoped>
.material-symbols-outlined {
  font-variation-settings:
    'FILL' 0,
    'wght' 400,
    'GRAD' 0,
    'opsz' 24;
}
</style>
