<script setup>
import { ref, computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useExamStore } from '@/stores/examStore'

const exam = useExamStore()
const { candidateName, questions, answers } = storeToRefs(exam)

const formError = ref('')

const allAnswered = computed(() => {
  return questions.value.every((q) => answers.value[q.id] != null)
})

function isSelected(questionId, optionId) {
  return answers.value[questionId] === optionId
}

function selectOption(questionId, optionId) {
  exam.setAnswer(questionId, optionId)
}

function handleSubmit() {
  formError.value = ''
  const name = candidateName.value.trim()
  if (!name) {
    formError.value = 'กรุณากรอกชื่อผู้สอบ'
    return
  }
  if (!allAnswered.value) {
    formError.value = 'กรุณาตอบให้ครบทุกข้อ'
    return
  }
  exam.submitExam()
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
      <!-- Header + Candidate Name -->
      <div class="mb-12">
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
                class="w-full rounded-t-lg border-none border-b-2 border-transparent bg-surface-container-highest py-3 pl-4 pr-4 font-sans text-on-surface transition-all duration-300 placeholder:text-outline/50 focus:border-primary focus:ring-0"
              />
              <div
                class="pointer-events-none absolute inset-x-0 bottom-0 h-0.5 bg-outline-variant/15 transition-colors group-focus-within:bg-primary"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- Questions: v-for จาก Store — ทุกข้อใช้ choice card สไตล์เดียวกัน -->
      <div class="space-y-12">
        <section
          v-for="(q, index) in questions"
          :key="q.id"
          class="rounded-3xl border border-outline-variant/10 bg-surface-container-lowest p-8 shadow-[0_8px_32px_rgba(25,28,30,0.04)] md:p-12"
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
        </section>
      </div>

      <!-- Submit -->
      <div class="mt-16 flex flex-col items-center gap-6">
        <p v-if="formError" class="text-center text-sm font-medium text-red-600" role="alert">
          {{ formError }}
        </p>
        <button
          type="button"
          class="group relative w-full max-w-md overflow-hidden rounded-2xl bg-gradient-to-br from-primary to-primary-container p-4 text-xl font-bold text-on-primary shadow-xl shadow-indigo-500/20 transition-all active:scale-[0.98] hover:shadow-indigo-500/35"
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
