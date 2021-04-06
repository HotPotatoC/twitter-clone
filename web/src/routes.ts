import { createRouter, createWebHistory } from 'vue-router'
import BaseLayout from './layouts/base-layout.vue'
import Home from './pages/home/index.vue'
import Login from './pages/login.vue'
import { store, useStore } from './store'
import { ActionTypes } from './store/auth/actions'

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: BaseLayout,
      children: [
        {
          path: '/home',
          component: Home,
          meta: {
            requiresAuth: true,
          },
        },
        {
          path: '/:name/status/:tweetId',
          component: () => import('./pages/profile/status/index.vue'),
          meta: {
            requiresAuth: true,
          },
        },
      ],
    },
    {
      path: '/login',
      component: Login,
    },
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
