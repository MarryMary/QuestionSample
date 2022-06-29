import { createRouter, createWebHistory } from 'vue-router'


const routes = [
  {
    path: '/qwizard/create',
    name: 'create',
    component: () => import('../views/WhichMode.vue')
  },
  {
    path: '/qwizard/detail',
    name: 'DetailVue',
    component: () => import('../views/Detail.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
