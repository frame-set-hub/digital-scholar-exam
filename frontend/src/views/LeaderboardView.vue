<script setup>
import { computed, onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useExamStore } from '@/stores/examStore'

const exam = useExamStore()
const { leaderboard, leaderboardState, leaderboardError, leaderboardYourEntry } = storeToRefs(exam)

onMounted(() => {
  exam.loadLeaderboard()
})

const firstPlace = computed(() => leaderboard.value[0])
const secondPlace = computed(() => leaderboard.value[1])
const thirdPlace = computed(() => leaderboard.value[2])
const restRows = computed(() => leaderboard.value.slice(3))

function formatScore(entry) {
  if (!entry) return '—'
  return `${entry.score}/${entry.total}`
}

function formatDate(iso) {
  if (!iso) return '—'
  const d = new Date(iso)
  if (Number.isNaN(d.getTime())) return '—'
  return d.toLocaleString(undefined, { dateStyle: 'medium', timeStyle: 'short' })
}

function backToExam() {
  exam.resetExam()
}


</script>

<template>
  <div class="relative min-h-screen overflow-x-hidden bg-background font-sans text-on-surface">
    <header
      class="fixed top-0 z-50 w-full border-b border-outline-variant/20 bg-background/95 shadow-sm backdrop-blur-xl"
    >
      <div
        class="mx-auto flex min-h-[3.25rem] max-w-5xl items-center justify-between gap-3 px-3 py-3 sm:min-h-0 sm:px-4 sm:py-4"
      >
        <div class="min-w-0 flex-1">
          <span class="block truncate font-display text-base font-extrabold text-primary sm:text-lg"
            >Digital Scholar</span
          >
        </div>
        <button
          type="button"
          class="shrink-0 rounded-xl bg-surface-container-high px-3 py-2 text-sm font-medium text-on-surface transition-all hover:bg-surface-container-highest active:scale-95 sm:px-4"
          @click="backToExam"
        >
          Back to Exam
        </button>
      </div>
    </header>

    <!-- pt มากพอให้ไม่ทับ fixed header + หัวข้อ h1 (รวม safe area ด้านบนมือถือ) -->
    <main
      class="relative z-0 mx-auto w-full max-w-4xl px-3 pb-12 pt-[calc(5.5rem+env(safe-area-inset-top))] sm:max-w-5xl sm:px-5 sm:pb-16 sm:pt-28 md:px-6"
    >
      <div class="mb-8 text-center sm:mb-10">
        <h1 class="mb-2 font-display text-3xl font-extrabold tracking-tight text-on-surface sm:mb-3 sm:text-4xl md:text-5xl">
          Leaderboard
        </h1>
        <p class="text-base font-medium text-secondary sm:text-lg">Top scorers for Exam IT 10-1</p>
      </div>

      <div v-if="leaderboardState === 'loading'" class="py-20 text-center font-medium text-secondary">
        Loading…
      </div>
      <div
        v-else-if="leaderboardState === 'error'"
        class="rounded-2xl bg-red-50 px-4 py-6 text-center text-red-900"
      >
        {{ leaderboardError }}
      </div>
      <template v-else>
        <!-- Me: ตำแหน่งของคุณ — ด้านบน podium (ต้องโหลดด้วยชื่อใน store → API ส่ง ?forCandidate=...) -->
        <div
          v-if="leaderboardYourEntry"
          class="mx-auto mb-8 max-w-3xl overflow-hidden rounded-2xl border-2 border-primary/35 bg-gradient-to-br from-primary/5 to-surface-container-low px-4 py-4 shadow-md sm:px-6 sm:py-5"
          role="region"
          aria-label="Your position"
        >
          <div
            class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between sm:gap-6"
          >
            <div class="flex items-center justify-center gap-2 sm:justify-start">
              <span
                class="inline-flex shrink-0 items-center rounded-full bg-primary px-3 py-1.5 text-[11px] font-black uppercase tracking-[0.2em] text-on-primary"
                >Me</span
              >
              <span class="material-symbols-outlined text-2xl text-primary" aria-hidden="true">person</span>
            </div>
            <div class="min-w-0 flex-1 text-center sm:text-right">
              <p class="truncate font-display text-lg font-bold text-on-surface">
                {{ leaderboardYourEntry.candidateName }}
              </p>
              <p class="mt-1 text-base font-semibold text-on-surface">
                Rank <span class="font-black text-primary">#{{ leaderboardYourEntry.rank }}</span>
                · {{ formatScore(leaderboardYourEntry) }} points
              </p>
            </div>
          </div>
          <p v-if="leaderboardYourEntry.inTopList" class="mt-3 text-center text-sm text-secondary sm:text-left">
            You’re in the top 20 — find your row in the list below.
          </p>
          <p v-else class="mt-3 text-center text-sm font-medium text-amber-900 sm:text-left">
            Not in the top 20 below — this is your global rank (e.g. 0 points places after everyone with a higher score).
          </p>
        </div>

        <div v-if="leaderboard.length === 0" class="py-16 text-center font-medium text-secondary">
          No exam results yet.
        </div>

        <!-- Single leader — Champion: ป้ายทองหรูเท่านั้น (ไม่แสดงเลข 1) -->
        <div v-else-if="leaderboard.length === 1 && firstPlace" class="mb-12 flex justify-center">
          <div
            class="flex w-full max-w-sm flex-col items-center rounded-2xl border-2 border-indigo-100 bg-white px-6 pb-10 pt-8 text-center shadow-[0_20px_40px_rgba(53,37,205,0.06)] sm:px-8 sm:pt-10"
          >
            <div
              class="champion-badge mb-5 flex items-center gap-2 rounded-full px-5 py-2 text-[10px] font-bold uppercase tracking-[0.2em] sm:mb-6 sm:px-6"
            >
              <span class="material-symbols-outlined champion-badge-icon">workspace_premium</span>
              <span>Champion</span>
            </div>
            <h3 class="mb-1 font-display text-2xl font-bold">{{ firstPlace.candidateName }}</h3>
            <p class="font-display text-3xl font-black text-primary">{{ formatScore(firstPlace) }} Points</p>
            <p class="mt-6 text-xs font-medium text-on-surface-variant">
              {{ formatDate(firstPlace.createdAt) }}
            </p>
          </div>
        </div>

        <!-- Top 2: DOM 1→2 มือถือบนลงล่าง · md: ซ้าย=2 ขวา=1 -->
        <div
          v-else-if="leaderboard.length === 2 && firstPlace && secondPlace"
          class="mx-auto mb-10 grid w-full max-w-3xl grid-cols-1 items-end gap-6 md:mb-12 md:grid-cols-2"
        >
          <div
            class="flex flex-col items-center rounded-2xl border-2 border-indigo-100 bg-white px-6 pb-8 pt-8 text-center shadow-[0_20px_40px_rgba(53,37,205,0.06)] sm:px-8 sm:pb-10 sm:pt-10 md:order-2 md:scale-105"
          >
            <div
              class="champion-badge mb-5 flex items-center gap-2 rounded-full px-5 py-2 text-[10px] font-bold uppercase tracking-[0.2em] sm:mb-6 sm:px-6"
            >
              <span class="material-symbols-outlined champion-badge-icon">workspace_premium</span>
              <span>Champion</span>
            </div>
            <h3 class="mb-1 font-display text-2xl font-bold">{{ firstPlace.candidateName }}</h3>
            <p class="font-display text-3xl font-black text-primary">{{ formatScore(firstPlace) }} Points</p>
            <p class="mt-6 text-xs font-medium text-on-surface-variant">{{ formatDate(firstPlace.createdAt) }}</p>
          </div>
          <div
            class="flex flex-col items-center rounded-2xl border border-outline-variant/10 bg-surface-container-lowest p-6 text-center shadow-sm md:order-1"
          >
            <div
              class="rank-circle rank-circle-silver mx-auto mb-4 sm:mb-5"
            >
              {{ secondPlace.rank }}
            </div>
            <h3 class="mb-1 font-display text-xl font-bold">{{ secondPlace.candidateName }}</h3>
            <p class="text-lg font-bold text-primary">{{ formatScore(secondPlace) }} Points</p>
            <p class="mt-4 text-xs font-medium text-on-surface-variant opacity-70">
              {{ formatDate(secondPlace.createdAt) }}
            </p>
          </div>
        </div>

        <!-- Top 3 podium — DOM: 1→2→3 มือถืออ่านบนลงล่าง · min-[480px]: order ให้เป็น 2|1|3 -->
        <div
          v-else-if="leaderboard.length >= 3 && firstPlace && secondPlace && thirdPlace"
          class="mb-8 grid grid-cols-1 items-stretch justify-items-stretch gap-3 sm:mb-10 sm:items-end sm:gap-4 lg:gap-6 min-[480px]:grid-cols-3"
        >
          <!-- Rank 1 Champion — w-full ให้เต็มคอลัมน์ (ไม่หดตามเนื้อหาเหมือนการ์ด 2/3) -->
          <div
            class="relative z-10 flex min-h-[min(22rem,78svh)] w-full min-w-0 flex-col items-center rounded-2xl border-2 border-indigo-100 bg-white px-6 pb-8 pt-8 text-center shadow-[0_20px_40px_rgba(53,37,205,0.06)] sm:min-h-[24rem] sm:px-8 sm:pb-10 sm:pt-10 min-[480px]:order-2 min-[480px]:scale-[1.04] lg:scale-110"
          >
            <div
              class="champion-badge relative z-30 mb-5 flex items-center gap-2 whitespace-nowrap rounded-full px-4 py-2 text-[10px] font-bold uppercase tracking-[0.2em] sm:mb-6 sm:px-6"
            >
              <span class="material-symbols-outlined champion-badge-icon">workspace_premium</span>
              <span>Champion</span>
            </div>
            <h3 class="mb-1 font-display text-2xl font-bold">{{ firstPlace.candidateName }}</h3>
            <p class="font-display text-3xl font-black text-primary">{{ formatScore(firstPlace) }} Points</p>
            <p class="mt-4 text-xs font-medium text-on-surface-variant sm:mt-6">{{ formatDate(firstPlace.createdAt) }}</p>
          </div>
          <!-- Rank 2 (ซ้ายตอน 3 คอลัมน์) — กระชับกว่าเดิม (สลับกับอันดับ 3) -->
          <div
            class="flex min-h-0 w-full min-w-0 flex-col items-center rounded-2xl border border-outline-variant/10 bg-surface-container-lowest p-5 text-center shadow-sm transition-all hover:shadow-md min-[480px]:order-1 sm:p-6"
          >
            <div
              class="rank-circle rank-circle-silver mx-auto mb-4 sm:mb-5"
            >
              {{ secondPlace.rank }}
            </div>
            <h3 class="mb-1 font-display text-xl font-bold">{{ secondPlace.candidateName }}</h3>
            <p class="text-lg font-bold text-primary">{{ formatScore(secondPlace) }} Points</p>
            <p class="mt-2 text-xs font-medium text-on-surface-variant opacity-70 sm:mt-3">
              {{ formatDate(secondPlace.createdAt) }}
            </p>
          </div>
          <!-- Rank 3 (ขวา) — โปร่งกว่าเดิม (สลับกับอันดับ 2) -->
          <div
            class="flex min-h-0 w-full min-w-0 flex-col items-center rounded-2xl border border-outline-variant/10 bg-surface-container-lowest p-7 text-center shadow-sm transition-all hover:shadow-md min-[480px]:order-3 sm:p-8"
          >
            <div
              class="rank-circle rank-circle-bronze mx-auto mb-4 sm:mb-5"
            >
              {{ thirdPlace.rank }}
            </div>
            <h3 class="mb-1 font-display text-xl font-bold">{{ thirdPlace.candidateName }}</h3>
            <p class="text-lg font-bold text-primary">{{ formatScore(thirdPlace) }} Points</p>
            <p class="mt-4 text-xs font-medium text-on-surface-variant opacity-70 sm:mt-5">
              {{ formatDate(thirdPlace.createdAt) }}
            </p>
          </div>
        </div>

        <!-- Rank 4+ — วงเลข = div แรกในแถว (h-10 w-10 …) -->
        <div v-if="restRows.length > 0" class="mx-auto max-w-3xl space-y-3">
          <div
            v-for="row in restRows"
            :key="row.rank + row.candidateName + row.createdAt"
            class="group flex items-center gap-6 rounded-2xl bg-surface-container-low px-5 py-5 transition-all hover:bg-surface-container sm:gap-8 sm:px-6"
          >
            <div
              class="flex h-10 w-10 shrink-0 items-center justify-center self-center rounded-full bg-surface-container-highest font-display text-lg font-bold text-secondary ring-1 ring-outline-variant/25 mr-3 sm:mr-4"
            >
              {{ row.rank }}
            </div>
            <div class="min-w-0 flex-1 pl-0.5 sm:pl-1">
              <p class="truncate text-lg font-bold text-on-surface">{{ row.candidateName }}</p>
              <p class="text-xs font-medium text-on-surface-variant">{{ formatDate(row.createdAt) }}</p>
            </div>
            <div class="shrink-0 text-right">
              <p class="font-display text-xl font-black text-primary">{{ formatScore(row) }}</p>
              <p class="text-[10px] font-bold uppercase tracking-tighter text-secondary">Points Scored</p>
            </div>
          </div>
        </div>

        <div class="mt-10 flex justify-center px-1 sm:mt-14">
          <button
            type="button"
            class="inline-flex min-h-[3rem] w-full max-w-md items-center justify-center gap-3 rounded-2xl bg-primary py-3.5 pl-5 pr-6 text-base font-bold text-on-primary shadow-lg transition-all hover:shadow-indigo-200 active:scale-95 sm:w-auto sm:min-w-[min(100%,18rem)] sm:px-10 sm:py-4"
            @click="backToExam"
          >
            <span class="material-symbols-outlined shrink-0 text-[1.35rem] leading-none">arrow_back</span>
            <span class="text-center leading-snug">Back to Exam</span>
          </button>
        </div>
      </template>
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
/* ป้าย Champion — ทอง luxury: ไฮไลต์ + ขอบขาว/ทอง (ตำแหน่งใช้ class ใน template) */
.champion-badge {
  color: #1a1410;
  letter-spacing: 0.18em;
  background: linear-gradient(
    145deg,
    #fffef9 0%,
    #f5e6c8 18%,
    #e4c76a 42%,
    #c9a227 68%,
    #8a6520 100%
  );
  border: 1px solid rgba(255, 252, 245, 0.95);
  box-shadow:
    inset 0 2px 0 rgba(255, 255, 255, 0.65),
    inset 0 -2px 0 rgba(0, 0, 0, 0.08),
    0 8px 28px rgba(138, 101, 32, 0.38),
    0 0 0 2px rgba(255, 255, 255, 0.92),
    0 0 0 3px rgba(201, 162, 39, 0.28);
}
.champion-badge-icon {
  font-size: 1.05rem;
  color: #4a3720;
  filter: drop-shadow(0 1px 0 rgba(255, 255, 255, 0.45));
  font-variation-settings:
    'FILL' 1,
    'wght' 500,
    'GRAD' 0,
    'opsz' 24;
}
/* วงกลมอันดับ 2–3 — pure CSS ไม่พึ่ง Tailwind utility */
.rank-circle {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 4rem;
  height: 4rem;
  border-radius: 9999px;
  border: 3px solid #fff;
  font-size: 1.5rem;
  font-weight: 700;
  line-height: 1;
  color: #fff;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
  box-shadow:
    0 4px 14px rgba(0, 0, 0, 0.15),
    0 1px 3px rgba(0, 0, 0, 0.1);
  flex-shrink: 0;
}
.rank-circle-silver {
  background: linear-gradient(135deg, #b8b8b8 0%, #8a8a8a 100%);
}
.rank-circle-bronze {
  background: linear-gradient(135deg, #cd7f32 0%, #a0522d 100%);
}
</style>
