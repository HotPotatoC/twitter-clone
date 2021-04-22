import { RouteRecordRaw } from 'vue-router'
import BaseLayout from '../../modules/layouts/BaseLayout.vue'
import { makeRoutesWithLayout } from '../../utils/routes'

const routes: RouteRecordRaw[] = [
  {
    path: '',
    component: () => import('./index.vue'),
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: 'status/:tweetId',
    component: () => import('./status/index.vue'),
    meta: {
      requiresAuth: true,
    },
  },
]

export default makeRoutesWithLayout('/:name', BaseLayout, routes)
