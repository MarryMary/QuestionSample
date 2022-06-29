import { createRouter, createWebHistory } from 'vue-router'
import NotFound from '../views/NotFound'


const routes = [
  {
    path: '/exwizard/start',
    name: 'create',
    component: () => import('../views/WhichMode.vue')
  },
  {
    path: '/exwizard/detail',
    name: 'DetailVue',
    component: () => import('../views/Detail.vue')
  },
  {
    path: '/exwizard/check',
    name: 'CheckDetail',
    component: () => import('../views/Check.vue')
  },
  {
    path: '/exwizard/SuccessFinished',
    name: 'Finished',
    component: () => import('../views/SuccessFinished.vue')
  },
  {
    path: '/:pathMatch(.*)*',
    component: NotFound
  }

]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
