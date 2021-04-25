import { createRouter, createWebHistory } from 'vue-router'
import { store } from './store'
import routes from './modules/routes'
import requiresAuth from './guards/requiresAuth'
import isAuthenticated from './guards/isAuthenticated'

export const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(async (to, from, next) => {
  if (to.matched.some((route) => route.meta.requiresAuth)) {
    await requiresAuth({ to, from, next, store })
  } else {
    await isAuthenticated({ to, from, next, store })
  }
})
