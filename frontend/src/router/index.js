import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'exam',
      component: () => import('@/views/ExamView.vue'),
      meta: { title: 'IT 10-1 — Digital Scholar' },
    },
    {
      path: '/result',
      name: 'result',
      component: () => import('@/views/ResultView.vue'),
      meta: { title: 'IT 10-2 — Digital Scholar' },
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/',
    },
  ],
})

router.afterEach((to) => {
  const title = to.meta.title
  if (title) document.title = String(title)
})

export default router
