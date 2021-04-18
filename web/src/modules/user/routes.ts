import { RouteRecordRaw } from 'vue-router'
import BaseLayout from '../../layouts/BaseLayout.vue'
import { makeRoutesWithLayout } from '../../services/routes'

const routes: RouteRecordRaw[] = [
  {
    path: 'status/:tweetId',
    component: () => import('./status/index.vue'),
    meta: {
      requiresAuth: true,
    },
  },
]

export default makeRoutesWithLayout('/:name', BaseLayout, routes)
