<script setup>
import { computed, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useExamStore } from '@/stores/examStore'

const router = useRouter()
const exam = useExamStore()
const { candidateName, score, totalQuestions } = storeToRefs(exam)

onBeforeMount(() => {
  if (score.value === null) {
    router.replace({ name: 'exam' })
  }
})

const circumference = 2 * Math.PI * 88
const scoreProgress = computed(() => {
  if (totalQuestions.value === 0) return 0
  return (score.value ?? 0) / totalQuestions.value
})

const strokeDashoffset = computed(() => {
  return circumference * (1 - scoreProgress.value)
})

function retake() {
  exam.resetExam()
}

function goLeaderboard() {
  router.push({ name: 'leaderboard' })
}
</script>

<template>
  <div class="relative flex min-h-screen flex-col overflow-x-hidden bg-background font-sans text-on-surface">
    <!-- ไม่มี Navbar / Bottom nav — ไม่มี Time Taken / Percentile -->

    <main class="flex flex-grow flex-col items-center justify-center px-4 py-12 sm:px-6">
      <div class="w-full max-w-3xl">
        <div class="grid grid-cols-1 gap-6 md:grid-cols-12">
          <!-- Hero -->
          <div
            class="relative overflow-hidden rounded-3xl bg-surface-container-lowest p-8 text-center md:col-span-12 md:p-12"
          >
            <div class="absolute left-0 top-0 h-1 w-full bg-gradient-to-r from-primary to-tertiary" />
            <div
              class="mx-auto mb-6 flex h-20 w-20 items-center justify-center rounded-full bg-primary-fixed text-primary"
            >
              <span class="material-symbols-outlined text-4xl">assignment_turned_in</span>
            </div>
            <h1 class="mb-2 font-display text-3xl font-extrabold text-on-surface md:text-4xl">
              Assessment Complete
            </h1>
            <p class="font-sans font-medium text-secondary">Exam IT 10-2</p>
            <div class="mb-8 mt-10 inline-block rounded-2xl bg-surface-container-low px-8 py-6">
              <span class="mb-1 block font-sans text-sm font-medium uppercase tracking-widest text-on-surface-variant">
                NAME
              </span>
              <span class="block font-display text-2xl font-bold text-on-surface">
                {{ candidateName.trim() || '—' }}
              </span>
            </div>
          </div>

          <!-- Score -->
          <div
            class="flex min-h-[280px] flex-col items-center justify-center rounded-3xl bg-surface-container-lowest p-8 md:col-span-12"
          >
            <span class="mb-6 font-sans font-semibold text-on-surface-variant">Final Score</span>
            <div class="relative flex h-48 w-48 items-center justify-center">
              <svg
                class="absolute inset-0 h-full w-full -rotate-90"
                viewBox="0 0 192 192"
                aria-hidden="true"
              >
                <circle
                  cx="96"
                  cy="96"
                  r="88"
                  fill="transparent"
                  stroke="currentColor"
                  stroke-width="12"
                  class="text-surface-container"
                />
                <circle
                  cx="96"
                  cy="96"
                  r="88"
                  fill="transparent"
                  stroke="currentColor"
                  stroke-width="12"
                  stroke-linecap="round"
                  class="text-primary transition-[stroke-dashoffset] duration-500 ease-out"
                  :stroke-dasharray="String(circumference)"
                  :stroke-dashoffset="String(strokeDashoffset)"
                />
              </svg>
              <div class="relative z-10 text-center">
                <span class="font-display text-5xl font-extrabold text-primary sm:text-6xl">
                  {{ score }} / {{ totalQuestions }}
                </span>
                <span class="mt-1 block font-sans font-medium text-secondary">Points</span>
              </div>
            </div>

            <div class="mt-10 w-full max-w-sm space-y-3">
              <button
                type="button"
                class="flex w-full items-center justify-center gap-3 rounded-2xl border-2 border-primary/30 bg-surface-container-lowest py-4 px-8 font-display text-lg font-bold text-primary shadow-sm transition-transform active:scale-95"
                @click="goLeaderboard"
              >
                <span class="material-symbols-outlined">leaderboard</span>
                View Leaderboard
              </button>
              <button
                type="button"
                class="flex w-full items-center justify-center gap-3 rounded-2xl bg-gradient-to-br from-primary to-primary-container py-5 px-8 font-display text-lg font-bold text-on-primary shadow-lg shadow-indigo-500/20 transition-transform active:scale-95"
                @click="retake"
              >
                <span class="material-symbols-outlined">refresh</span>
                Retake Exam
              </button>
            </div>
          </div>
        </div>
      </div>
    </main>

    <div
      class="pointer-events-none fixed top-0 right-0 -z-10 h-1/2 w-1/2 bg-gradient-to-l from-primary/10 to-transparent blur-3xl"
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
