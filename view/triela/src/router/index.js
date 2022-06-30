import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'MainMenu',
    component: () => import('../views/Menu.vue')
  },
  {
    path: '/manage/Ex',
    name: 'ExerciseManage',
    component: () => import('../views/ExerciseManagement.vue')
  },
  {
    path: '/manage/fix/:id',
    name: 'ExerciseFix',
    component: () => import('../views/ExerciseFix')
  },
  {
    path: '/wizard/Ex',
    name: 'ExerciseWizard',
    component: () => import('../views/ExCreate.vue')
  },
  {
    path: '/:catchAll(.*)',
    name: 'NotFound',
    component: () => import('../views/NotFound.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
