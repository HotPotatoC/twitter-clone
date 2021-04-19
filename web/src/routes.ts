import { createRouter, createWebHistory } from 'vue-router'
import routes from './modules/routes'
import { store } from './store'
import { ActionTypes } from './modules/auth/store/actions'

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/home',
    },
    ...routes,
  ],
})

router.beforeEach(async (to, from, next) => {
  if (to.matched.some((route) => route.meta.requiresAuth)) {
    await store.dispatch(ActionTypes.REFRESH_AUTH_TOKEN)
    if (!store.getters['isLoggedIn']) {
      next({
        path: '/login',
        query: { redirect: to.fullPath },
      })
    } else {
      next()
    }

    await store.dispatch(ActionTypes.GET_USER_DATA)
    if (!store.getters['isLoggedIn']) {
      next({
        path: '/login',
        query: { redirect: to.fullPath },
      })
    } else {
      next()
    }
  } else {
    next()
  }
})
