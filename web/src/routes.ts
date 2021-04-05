import { createRouter, createWebHistory } from 'vue-router'
import BaseLayout from './layouts/base-layout.vue'
import Home from './pages/home/index.vue'
import Login from './pages/login.vue'

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
        },
        {
          path: '/:name/status/:tweetId',
          component: () => import('./pages/profile/status/index.vue'),
        },
      ],
    },
    {
      path: '/login',
      component: Login,
    },
  ],
})
