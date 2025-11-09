import { createRouter, createWebHistory } from 'vue-router'
import AlertsView from '../views/AlertsView.vue'

const routes = [{ path: '/', name: 'alerts', component: AlertsView }]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
})

export default router
